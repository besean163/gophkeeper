// Code generated by MockGen. DO NOT EDIT.
// Source: internal/server/interfaces/bucket_service.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	models "github.com/besean163/gophkeeper/internal/server/models"
	gomock "github.com/golang/mock/gomock"
)

// MockBucketService is a mock of BucketService interface.
type MockBucketService struct {
	ctrl     *gomock.Controller
	recorder *MockBucketServiceMockRecorder
}

// MockBucketServiceMockRecorder is the mock recorder for MockBucketService.
type MockBucketServiceMockRecorder struct {
	mock *MockBucketService
}

// NewMockBucketService creates a new mock instance.
func NewMockBucketService(ctrl *gomock.Controller) *MockBucketService {
	mock := &MockBucketService{ctrl: ctrl}
	mock.recorder = &MockBucketServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBucketService) EXPECT() *MockBucketServiceMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockBucketService) CreateAccount(account *models.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", account)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockBucketServiceMockRecorder) CreateAccount(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockBucketService)(nil).CreateAccount), account)
}

// DeleteAccount mocks base method.
func (m *MockBucketService) DeleteAccount(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockBucketServiceMockRecorder) DeleteAccount(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockBucketService)(nil).DeleteAccount), id)
}

// GetAccounts mocks base method.
func (m *MockBucketService) GetAccounts() []*models.Account {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccounts")
	ret0, _ := ret[0].([]*models.Account)
	return ret0
}

// GetAccounts indicates an expected call of GetAccounts.
func (mr *MockBucketServiceMockRecorder) GetAccounts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*MockBucketService)(nil).GetAccounts))
}

// UpdateAccount mocks base method.
func (m *MockBucketService) UpdateAccount(account *models.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", account)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockBucketServiceMockRecorder) UpdateAccount(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockBucketService)(nil).UpdateAccount), account)
}
