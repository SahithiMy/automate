// Code generated by MockGen. DO NOT EDIT.
// Source: authn/authenticate.pb.go

// Package authn is a generated GoMock package.
package authn

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockAuthenticationServiceClient is a mock of AuthenticationServiceClient interface.
type MockAuthenticationServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationServiceClientMockRecorder
}

// MockAuthenticationServiceClientMockRecorder is the mock recorder for MockAuthenticationServiceClient.
type MockAuthenticationServiceClientMockRecorder struct {
	mock *MockAuthenticationServiceClient
}

// NewMockAuthenticationServiceClient creates a new mock instance.
func NewMockAuthenticationServiceClient(ctrl *gomock.Controller) *MockAuthenticationServiceClient {
	mock := &MockAuthenticationServiceClient{ctrl: ctrl}
	mock.recorder = &MockAuthenticationServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthenticationServiceClient) EXPECT() *MockAuthenticationServiceClientMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockAuthenticationServiceClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Authenticate", varargs...)
	ret0, _ := ret[0].(*AuthenticateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockAuthenticationServiceClientMockRecorder) Authenticate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockAuthenticationServiceClient)(nil).Authenticate), varargs...)
}

// MockAuthenticationServiceServer is a mock of AuthenticationServiceServer interface.
type MockAuthenticationServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationServiceServerMockRecorder
}

// MockAuthenticationServiceServerMockRecorder is the mock recorder for MockAuthenticationServiceServer.
type MockAuthenticationServiceServerMockRecorder struct {
	mock *MockAuthenticationServiceServer
}

// NewMockAuthenticationServiceServer creates a new mock instance.
func NewMockAuthenticationServiceServer(ctrl *gomock.Controller) *MockAuthenticationServiceServer {
	mock := &MockAuthenticationServiceServer{ctrl: ctrl}
	mock.recorder = &MockAuthenticationServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthenticationServiceServer) EXPECT() *MockAuthenticationServiceServerMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockAuthenticationServiceServer) Authenticate(arg0 context.Context, arg1 *AuthenticateRequest) (*AuthenticateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", arg0, arg1)
	ret0, _ := ret[0].(*AuthenticateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockAuthenticationServiceServerMockRecorder) Authenticate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockAuthenticationServiceServer)(nil).Authenticate), arg0, arg1)
}
