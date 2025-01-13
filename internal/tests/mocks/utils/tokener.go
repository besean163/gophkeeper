// Code generated by MockGen. DO NOT EDIT.
// Source: /opt/gophkeeper/internal/utils/api_token/tokener.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	server "github.com/besean163/gophkeeper/internal/models/server"
	gomock "github.com/golang/mock/gomock"
)

// MockTokener is a mock of Tokener interface.
type MockTokener struct {
	ctrl     *gomock.Controller
	recorder *MockTokenerMockRecorder
}

// MockTokenerMockRecorder is the mock recorder for MockTokener.
type MockTokenerMockRecorder struct {
	mock *MockTokener
}

// NewMockTokener creates a new mock instance.
func NewMockTokener(ctrl *gomock.Controller) *MockTokener {
	mock := &MockTokener{ctrl: ctrl}
	mock.recorder = &MockTokenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokener) EXPECT() *MockTokenerMockRecorder {
	return m.recorder
}

// GetToken mocks base method.
func (m *MockTokener) GetToken(user *server.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToken indicates an expected call of GetToken.
func (mr *MockTokenerMockRecorder) GetToken(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken", reflect.TypeOf((*MockTokener)(nil).GetToken), user)
}

// GetUserId mocks base method.
func (m *MockTokener) GetUserId(token string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserId", token)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserId indicates an expected call of GetUserId.
func (mr *MockTokenerMockRecorder) GetUserId(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserId", reflect.TypeOf((*MockTokener)(nil).GetUserId), token)
}
