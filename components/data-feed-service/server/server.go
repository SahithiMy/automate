package server

import (
	"context"
	"strconv"

	//"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	"bytes"
	"compress/gzip"
	"encoding/json"
	"net/http"

	datafeed "github.com/chef/automate/api/external/data_feed"
	"github.com/chef/automate/api/external/lib/errorutils"
	"github.com/chef/automate/api/external/secrets"
	"github.com/chef/automate/components/data-feed-service/config"
	"github.com/chef/automate/components/data-feed-service/dao"
	"github.com/chef/automate/components/data-feed-service/service"
	"github.com/chef/automate/lib/grpc/health"
	"github.com/chef/automate/lib/grpc/secureconn"
	"github.com/pkg/errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DatafeedServer is the interface to this component.
type DatafeedServer struct {
	db                  *dao.DB
	health              *health.Service
	secrets             secrets.SecretsServiceClient
	acceptedStatusCodes []int32
}

// New creates a new DatafeedServer instance.
func NewDatafeedServer(db *dao.DB, config *config.DataFeedConfig, connFactory *secureconn.Factory) (*DatafeedServer, error) {
	log.Info("NewDatafeedServer")
	secretsConn, err := connFactory.Dial("secrets-service", config.SecretsConfig.Target)
	if err != nil {
		return nil, errors.Wrap(err, "could not connect to secrets-service")
	}
	return &DatafeedServer{
		db:                  db,
		health:              health.NewService(),
		secrets:             secrets.NewSecretsServiceClient(secretsConn),
		acceptedStatusCodes: config.ServiceConfig.AcceptedStatusCodes,
	}, nil
}

// Add a new destination
func (datafeedServer *DatafeedServer) AddDestination(ctx context.Context, destination *datafeed.AddDestinationRequest) (*datafeed.AddDestinationResponse, error) {
	log.Infof("AddDestination %s", destination)
	response := &datafeed.AddDestinationResponse{Name: destination.Name, Url: destination.Url, Secret: destination.Secret, Services: destination.Services, IntegrationTypes: destination.IntegrationTypes, MetaData: destination.MetaData}
	id, err := datafeedServer.db.AddDestination(destination)
	response.Id = id
	if err != nil {
		return response, errorutils.FormatErrorMsg(err, "")
	}

	return response, nil
}

func (datafeedServer *DatafeedServer) TestDestination(ctx context.Context, request *datafeed.URLValidationRequest) (*datafeed.TestDestinationResponse, error) {
	response := &datafeed.TestDestinationResponse{Success: false}
	// http client to endpoint {text: "TEST: Successful validation completed by Automate"}
	// if it's secret - get the credentials
	// otherwise use passwd
	username := ""
	password := ""
	headers := ""
	accesskey := ""
	secretAccessKey := ""
	region := ""
	bucket := ""

	var err error
	var credentials service.Credentials
	url := request.Url
	serviceType := request.Services
	integrationType := request.IntegrationTypes

	
	zaMap := make(map[string]string, 0)
	for _, kv := range request.MetaData {
		zaMap[kv.Key] = kv.Value
	}
	jsonMap, err := json.Marshal(zaMap)
	if err != nil {
		logrus.Println(errors.Wrap(err, "keyValueToRawMap unable to marshal map"))
	}
	metaData := string(jsonMap)


	switch request.Credentials.(type) {
	case *datafeed.URLValidationRequest_UsernamePassword:
		username = request.GetUsernamePassword().GetUsername()
		password = request.GetUsernamePassword().GetPassword()
		credentials = service.NewBasicAuthCredentials(username, password)

	case *datafeed.URLValidationRequest_Header:
		headers = request.GetHeader().GetValue()
		credentials = service.NewCustomHeaderCredentials(headers)
	case *datafeed.URLValidationRequest_Aws:
		accesskey = request.GetAws().GetAccessKey()
		secretAccessKey = request.GetAws().GetSecretAccessKey()
		region = request.GetAws().GetRegion()
		bucket = request.GetAws().GetBucket()
		credentials = service.NewS3Credentials(accesskey, secretAccessKey, region, bucket)
	case *datafeed.URLValidationRequest_SecretId:
		secretId := request.GetSecretId().GetId()
		// call secrets api

		credentials, err = service.GetCredentials(context.Background(), datafeedServer.secrets, secretId, serviceType, integrationType, metaData)
		if err != nil {
			return response, err
		}
	}

	messageBytes, err := json.Marshal(map[string]string{
		"text": "TEST: Successful validation completed by Automate",
	})
	if err != nil {
		log.Errorf("Error creating json bytes %v", err)
		return response, err
	}

	var contentBuffer bytes.Buffer
	zip := gzip.NewWriter(&contentBuffer)
	_, err = zip.Write(messageBytes)
	if err != nil {
		return response, err
	}
	err = zip.Close()
	if err != nil {
		return response, err
	}
	if serviceType == "S3" || serviceType == "Minio" {
		cred := credentials.GetValues().AwsCreds
		sess := service.ConnectAWS(cred, url, serviceType)
		_, err := service.FileUploadInAws(sess, cred, messageBytes, "TestConnection")
		if err != nil {
			log.Error("Error creating Sending data to", serviceType)
			return response, err
		} else {
			response.Success = true
		}
	} else {
		httpRequest, err := http.NewRequest("POST", url, &contentBuffer)
		if err != nil {
			log.Error("Error creating request")
			return response, err
		}
		httpRequest.Header.Add("Authorization", credentials.GetValues().AuthorizationHeader)
		httpRequest.Header.Add("Content-Type", "application/json")
		httpRequest.Header.Add("Content-Encoding", "gzip")
		httpRequest.Header.Add("Accept", "application/json")

		service.AddCustomHeader(credentials, httpRequest.Header)

		client := http.Client{}
		httpResponse, err := client.Do(httpRequest)
		if err != nil {
			log.Errorf("Error sending test message %v", err)
			return response, err
		} else {
			log.Infof("Test data posted to %v, Status %v", url, httpResponse.Status)
		}
		if config.IsAcceptedStatusCode(int32(httpResponse.StatusCode), datafeedServer.acceptedStatusCodes) {
			response.Success = true
		} else {
			return response, status.Newf(codes.Internal, "%s posting test message to: %s", httpResponse.Status, url).Err()
		}
		err = httpResponse.Body.Close()
		if err != nil {
			log.Warnf("Error closing response body %v", err)
			return response, errorutils.FormatErrorMsg(err, "")
		}
	}

	return response, nil
}

func (datafeedServer *DatafeedServer) DeleteDestination(ctx context.Context, destination *datafeed.DeleteDestinationRequest) (*datafeed.DeleteDestinationResponse, error) {
	log.Infof("DeleteDestination %s", destination)
	fullDestination, err := datafeedServer.GetDestination(ctx, &datafeed.GetDestinationRequest{Id: destination.Id})
	var response *datafeed.DeleteDestinationResponse
	if err != nil {
		log.Warnf("Could not get destination details for delete response id: %d,  err: %s", destination.Id, err)
	} else {
		response = &datafeed.DeleteDestinationResponse{Id: fullDestination.Id, Name: fullDestination.Name, Url: fullDestination.Url, Secret: fullDestination.Secret, Services: fullDestination.Services, IntegrationTypes: fullDestination.IntegrationTypes, MetaData: fullDestination.MetaData}
	}

	err = datafeedServer.db.DeleteDestination(destination)
	if err != nil {
		return response, errorutils.FormatErrorMsg(err, strconv.FormatInt(destination.Id, 10))
	}

	return response, nil
}

func (datafeedServer *DatafeedServer) GetDestination(ctx context.Context, destination *datafeed.GetDestinationRequest) (*datafeed.GetDestinationResponse, error) {
	log.Infof("GetDestination %s", destination)
	response, err := datafeedServer.db.GetDestination(destination)
	if err != nil {
		return nil, errorutils.FormatErrorMsg(err, strconv.FormatInt(destination.Id, 10))
	}
	return response, nil
}

func (datafeedServer *DatafeedServer) ListDestinations(ctx context.Context, destination *datafeed.ListDestinationRequest) (*datafeed.ListDestinationResponse, error) {
	log.Infof("ListDestinations %s", destination)
	response, err := datafeedServer.db.ListDestinations()
	if err != nil {
		return nil, errorutils.FormatErrorMsg(err, "")
	}

	return response, nil
}

func (datafeedServer *DatafeedServer) UpdateDestination(ctx context.Context, destination *datafeed.UpdateDestinationRequest) (*datafeed.UpdateDestinationResponse, error) {
	log.Infof("UpdateDestination %s", destination)
	response := &datafeed.UpdateDestinationResponse{Name: destination.Name, Url: destination.Url, Secret: destination.Secret, Services: destination.Services, IntegrationTypes: destination.IntegrationTypes, MetaData: destination.MetaData}
	err := datafeedServer.db.UpdateDestination(destination)

	response.Id, _ = strconv.ParseInt(destination.Id, 10, 64)
	if err != nil {
		return nil, errorutils.FormatErrorMsg(err, "")
	}
	return response, nil
}

// Health returns the servers embedded health check service
func (datafeedServer *DatafeedServer) Health() *health.Service {
	return datafeedServer.health
}
