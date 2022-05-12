// -*- mode: protobuf; indent-tabs-mode: t; c-basic-offset: 8; tab-width: 8 -*-

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: config/report_manager_minio/config_request.proto

package report_manager_minio

import (
	shared "github.com/chef/automate/api/config/shared"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-a2-config/api/a2conf"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V1 *ConfigRequest_V1 `protobuf:"bytes,1,opt,name=v1,proto3" json:"v1,omitempty" toml:"v1,omitempty" mapstructure:"v1,omitempty"`
}

func (x *ConfigRequest) Reset() {
	*x = ConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_report_manager_minio_config_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest) ProtoMessage() {}

func (x *ConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_report_manager_minio_config_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest.ProtoReflect.Descriptor instead.
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return file_config_report_manager_minio_config_request_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigRequest) GetV1() *ConfigRequest_V1 {
	if x != nil {
		return x.V1
	}
	return nil
}

type ConfigRequest_V1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sys *ConfigRequest_V1_System `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty" toml:"sys,omitempty" mapstructure:"sys,omitempty"`
	Svc *ConfigRequest_Service   `protobuf:"bytes,2,opt,name=svc,proto3" json:"svc,omitempty" toml:"svc,omitempty" mapstructure:"svc,omitempty"`
}

func (x *ConfigRequest_V1) Reset() {
	*x = ConfigRequest_V1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_report_manager_minio_config_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1) ProtoMessage() {}

func (x *ConfigRequest_V1) ProtoReflect() protoreflect.Message {
	mi := &file_config_report_manager_minio_config_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1) Descriptor() ([]byte, []int) {
	return file_config_report_manager_minio_config_request_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ConfigRequest_V1) GetSys() *ConfigRequest_V1_System {
	if x != nil {
		return x.Sys
	}
	return nil
}

func (x *ConfigRequest_V1) GetSvc() *ConfigRequest_Service {
	if x != nil {
		return x.Svc
	}
	return nil
}

type ConfigRequest_Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConfigRequest_Service) Reset() {
	*x = ConfigRequest_Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_report_manager_minio_config_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_Service) ProtoMessage() {}

func (x *ConfigRequest_Service) ProtoReflect() protoreflect.Message {
	mi := &file_config_report_manager_minio_config_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_Service.ProtoReflect.Descriptor instead.
func (*ConfigRequest_Service) Descriptor() ([]byte, []int) {
	return file_config_report_manager_minio_config_request_proto_rawDescGZIP(), []int{0, 1}
}

type ConfigRequest_V1_System struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mlsa    *shared.Mlsa                     `protobuf:"bytes,1,opt,name=mlsa,proto3" json:"mlsa,omitempty" toml:"mlsa,omitempty" mapstructure:"mlsa,omitempty"`
	Tls     *shared.TLSCredentials           `protobuf:"bytes,2,opt,name=tls,proto3" json:"tls,omitempty" toml:"tls,omitempty" mapstructure:"tls,omitempty"`
	Service *ConfigRequest_V1_System_Service `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty" toml:"service,omitempty" mapstructure:"service,omitempty"`
}

func (x *ConfigRequest_V1_System) Reset() {
	*x = ConfigRequest_V1_System{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_report_manager_minio_config_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_System) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_System) ProtoMessage() {}

func (x *ConfigRequest_V1_System) ProtoReflect() protoreflect.Message {
	mi := &file_config_report_manager_minio_config_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_System.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_System) Descriptor() ([]byte, []int) {
	return file_config_report_manager_minio_config_request_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ConfigRequest_V1_System) GetMlsa() *shared.Mlsa {
	if x != nil {
		return x.Mlsa
	}
	return nil
}

func (x *ConfigRequest_V1_System) GetTls() *shared.TLSCredentials {
	if x != nil {
		return x.Tls
	}
	return nil
}

func (x *ConfigRequest_V1_System) GetService() *ConfigRequest_V1_System_Service {
	if x != nil {
		return x.Service
	}
	return nil
}

type ConfigRequest_V1_System_Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port *wrapperspb.Int32Value `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty" toml:"port,omitempty" mapstructure:"port,omitempty"`
}

func (x *ConfigRequest_V1_System_Service) Reset() {
	*x = ConfigRequest_V1_System_Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_report_manager_minio_config_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_System_Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_System_Service) ProtoMessage() {}

func (x *ConfigRequest_V1_System_Service) ProtoReflect() protoreflect.Message {
	mi := &file_config_report_manager_minio_config_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_System_Service.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_System_Service) Descriptor() ([]byte, []int) {
	return file_config_report_manager_minio_config_request_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

func (x *ConfigRequest_V1_System_Service) GetPort() *wrapperspb.Int32Value {
	if x != nil {
		return x.Port
	}
	return nil
}

var File_config_report_manager_minio_config_request_proto protoreflect.FileDescriptor

var file_config_report_manager_minio_config_request_proto_rawDesc = []byte{
	0x0a, 0x30, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x28, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x1a, 0x1a, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x74, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x3f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x61, 0x32, 0x2d, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x32, 0x63, 0x6f, 0x6e, 0x66,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf2, 0x04, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x02, 0x76, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3a, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0x52, 0x02, 0x76, 0x31,
	0x1a, 0xe5, 0x03, 0x0a, 0x02, 0x56, 0x31, 0x12, 0x53, 0x0a, 0x03, 0x73, 0x79, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x31,
	0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x03, 0x73, 0x79, 0x73, 0x12, 0x51, 0x0a, 0x03,
	0x73, 0x76, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x63, 0x68, 0x65, 0x66,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x6d,
	0x69, 0x6e, 0x69, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x03, 0x73, 0x76, 0x63, 0x1a,
	0xb6, 0x02, 0x0a, 0x06, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x34, 0x0a, 0x04, 0x6d, 0x6c,
	0x73, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d, 0x6c, 0x73, 0x61, 0x52, 0x04, 0x6d, 0x6c, 0x73, 0x61,
	0x12, 0x3c, 0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x4c, 0x53, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x03, 0x74, 0x6c, 0x73, 0x12, 0x63,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x49, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x5f, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x1a, 0x53, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x17, 0xc2, 0xf3, 0x18, 0x13, 0x0a,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x10, 0xd5, 0x4f, 0x1a, 0x05, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x09, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x3a, 0x22, 0xc2, 0xf3, 0x18, 0x1e, 0x0a, 0x1c, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2d, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x2d,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x6d, 0x69,
	0x6e, 0x69, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_report_manager_minio_config_request_proto_rawDescOnce sync.Once
	file_config_report_manager_minio_config_request_proto_rawDescData = file_config_report_manager_minio_config_request_proto_rawDesc
)

func file_config_report_manager_minio_config_request_proto_rawDescGZIP() []byte {
	file_config_report_manager_minio_config_request_proto_rawDescOnce.Do(func() {
		file_config_report_manager_minio_config_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_report_manager_minio_config_request_proto_rawDescData)
	})
	return file_config_report_manager_minio_config_request_proto_rawDescData
}

var file_config_report_manager_minio_config_request_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_config_report_manager_minio_config_request_proto_goTypes = []interface{}{
	(*ConfigRequest)(nil),                   // 0: chef.automate.infra.report_manager_minio.ConfigRequest
	(*ConfigRequest_V1)(nil),                // 1: chef.automate.infra.report_manager_minio.ConfigRequest.V1
	(*ConfigRequest_Service)(nil),           // 2: chef.automate.infra.report_manager_minio.ConfigRequest.Service
	(*ConfigRequest_V1_System)(nil),         // 3: chef.automate.infra.report_manager_minio.ConfigRequest.V1.System
	(*ConfigRequest_V1_System_Service)(nil), // 4: chef.automate.infra.report_manager_minio.ConfigRequest.V1.System.Service
	(*shared.Mlsa)(nil),                     // 5: chef.automate.infra.config.Mlsa
	(*shared.TLSCredentials)(nil),           // 6: chef.automate.infra.config.TLSCredentials
	(*wrapperspb.Int32Value)(nil),           // 7: google.protobuf.Int32Value
}
var file_config_report_manager_minio_config_request_proto_depIdxs = []int32{
	1, // 0: chef.automate.infra.report_manager_minio.ConfigRequest.v1:type_name -> chef.automate.infra.report_manager_minio.ConfigRequest.V1
	3, // 1: chef.automate.infra.report_manager_minio.ConfigRequest.V1.sys:type_name -> chef.automate.infra.report_manager_minio.ConfigRequest.V1.System
	2, // 2: chef.automate.infra.report_manager_minio.ConfigRequest.V1.svc:type_name -> chef.automate.infra.report_manager_minio.ConfigRequest.Service
	5, // 3: chef.automate.infra.report_manager_minio.ConfigRequest.V1.System.mlsa:type_name -> chef.automate.infra.config.Mlsa
	6, // 4: chef.automate.infra.report_manager_minio.ConfigRequest.V1.System.tls:type_name -> chef.automate.infra.config.TLSCredentials
	4, // 5: chef.automate.infra.report_manager_minio.ConfigRequest.V1.System.service:type_name -> chef.automate.infra.report_manager_minio.ConfigRequest.V1.System.Service
	7, // 6: chef.automate.infra.report_manager_minio.ConfigRequest.V1.System.Service.port:type_name -> google.protobuf.Int32Value
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_config_report_manager_minio_config_request_proto_init() }
func file_config_report_manager_minio_config_request_proto_init() {
	if File_config_report_manager_minio_config_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_report_manager_minio_config_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_report_manager_minio_config_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_report_manager_minio_config_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_Service); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_report_manager_minio_config_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_System); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_report_manager_minio_config_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_System_Service); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_report_manager_minio_config_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_report_manager_minio_config_request_proto_goTypes,
		DependencyIndexes: file_config_report_manager_minio_config_request_proto_depIdxs,
		MessageInfos:      file_config_report_manager_minio_config_request_proto_msgTypes,
	}.Build()
	File_config_report_manager_minio_config_request_proto = out.File
	file_config_report_manager_minio_config_request_proto_rawDesc = nil
	file_config_report_manager_minio_config_request_proto_goTypes = nil
	file_config_report_manager_minio_config_request_proto_depIdxs = nil
}
