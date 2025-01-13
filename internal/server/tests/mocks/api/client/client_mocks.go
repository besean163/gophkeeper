// Code generated by MockGen. DO NOT EDIT.
// Source: /opt/gophkeeper/internal/server/api/client/http_client/http_client.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	httpclient "github.com/besean163/gophkeeper/internal/server/api/client/http_client"
	gomock "github.com/golang/mock/gomock"
)

// MockResponse is a mock of Response interface.
type MockResponse struct {
	ctrl     *gomock.Controller
	recorder *MockResponseMockRecorder
}

// MockResponseMockRecorder is the mock recorder for MockResponse.
type MockResponseMockRecorder struct {
	mock *MockResponse
}

// NewMockResponse creates a new mock instance.
func NewMockResponse(ctrl *gomock.Controller) *MockResponse {
	mock := &MockResponse{ctrl: ctrl}
	mock.recorder = &MockResponseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResponse) EXPECT() *MockResponseMockRecorder {
	return m.recorder
}

// Body mocks base method.
func (m *MockResponse) Body() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Body")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Body indicates an expected call of Body.
func (mr *MockResponseMockRecorder) Body() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Body", reflect.TypeOf((*MockResponse)(nil).Body))
}

// StatusCode mocks base method.
func (m *MockResponse) StatusCode() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatusCode")
	ret0, _ := ret[0].(int)
	return ret0
}

// StatusCode indicates an expected call of StatusCode.
func (mr *MockResponseMockRecorder) StatusCode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusCode", reflect.TypeOf((*MockResponse)(nil).StatusCode))
}

// MockHTTPClient is a mock of HTTPClient interface.
type MockHTTPClient struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPClientMockRecorder
}

// MockHTTPClientMockRecorder is the mock recorder for MockHTTPClient.
type MockHTTPClientMockRecorder struct {
	mock *MockHTTPClient
}

// NewMockHTTPClient creates a new mock instance.
func NewMockHTTPClient(ctrl *gomock.Controller) *MockHTTPClient {
	mock := &MockHTTPClient{ctrl: ctrl}
	mock.recorder = &MockHTTPClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPClient) EXPECT() *MockHTTPClientMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockHTTPClient) Delete(uri string, body interface{}, headers map[string]string) (httpclient.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", uri, body, headers)
	ret0, _ := ret[0].(httpclient.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockHTTPClientMockRecorder) Delete(uri, body, headers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockHTTPClient)(nil).Delete), uri, body, headers)
}

// Get mocks base method.
func (m *MockHTTPClient) Get(uri string, headers map[string]string) (httpclient.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", uri, headers)
	ret0, _ := ret[0].(httpclient.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockHTTPClientMockRecorder) Get(uri, headers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHTTPClient)(nil).Get), uri, headers)
}

// Post mocks base method.
func (m *MockHTTPClient) Post(uri string, body interface{}, headers map[string]string) (httpclient.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", uri, body, headers)
	ret0, _ := ret[0].(httpclient.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post.
func (mr *MockHTTPClientMockRecorder) Post(uri, body, headers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockHTTPClient)(nil).Post), uri, body, headers)
}

// Put mocks base method.
func (m *MockHTTPClient) Put(uri string, body interface{}, headers map[string]string) (httpclient.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", uri, body, headers)
	ret0, _ := ret[0].(httpclient.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockHTTPClientMockRecorder) Put(uri, body, headers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockHTTPClient)(nil).Put), uri, body, headers)
}
