// Code generated by MockGen. DO NOT EDIT.
// Source: /opt/gophkeeper/internal/client/core/interfaces/data_service.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	client "github.com/besean163/gophkeeper/internal/models/client"
	gomock "github.com/golang/mock/gomock"
)

// MockDataService is a mock of DataService interface.
type MockDataService struct {
	ctrl     *gomock.Controller
	recorder *MockDataServiceMockRecorder
}

// MockDataServiceMockRecorder is the mock recorder for MockDataService.
type MockDataServiceMockRecorder struct {
	mock *MockDataService
}

// NewMockDataService creates a new mock instance.
func NewMockDataService(ctrl *gomock.Controller) *MockDataService {
	mock := &MockDataService{ctrl: ctrl}
	mock.recorder = &MockDataServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataService) EXPECT() *MockDataServiceMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockDataService) CreateAccount(user client.User, item client.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", user, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockDataServiceMockRecorder) CreateAccount(user, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockDataService)(nil).CreateAccount), user, item)
}

// CreateCard mocks base method.
func (m *MockDataService) CreateCard(user client.User, item client.Card) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCard", user, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCard indicates an expected call of CreateCard.
func (mr *MockDataServiceMockRecorder) CreateCard(user, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCard", reflect.TypeOf((*MockDataService)(nil).CreateCard), user, item)
}

// CreateNote mocks base method.
func (m *MockDataService) CreateNote(user client.User, item client.Note) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNote", user, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNote indicates an expected call of CreateNote.
func (mr *MockDataServiceMockRecorder) CreateNote(user, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNote", reflect.TypeOf((*MockDataService)(nil).CreateNote), user, item)
}

// DeleteAccount mocks base method.
func (m *MockDataService) DeleteAccount(user client.User, item client.Account, soft bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", user, item, soft)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockDataServiceMockRecorder) DeleteAccount(user, item, soft interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockDataService)(nil).DeleteAccount), user, item, soft)
}

// DeleteCard mocks base method.
func (m *MockDataService) DeleteCard(user client.User, item client.Card, soft bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCard", user, item, soft)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCard indicates an expected call of DeleteCard.
func (mr *MockDataServiceMockRecorder) DeleteCard(user, item, soft interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCard", reflect.TypeOf((*MockDataService)(nil).DeleteCard), user, item, soft)
}

// DeleteNote mocks base method.
func (m *MockDataService) DeleteNote(user client.User, item client.Note, soft bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNote", user, item, soft)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNote indicates an expected call of DeleteNote.
func (mr *MockDataServiceMockRecorder) DeleteNote(user, item, soft interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNote", reflect.TypeOf((*MockDataService)(nil).DeleteNote), user, item, soft)
}

// GetAccounts mocks base method.
func (m *MockDataService) GetAccounts(user client.User) ([]client.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccounts", user)
	ret0, _ := ret[0].([]client.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccounts indicates an expected call of GetAccounts.
func (mr *MockDataServiceMockRecorder) GetAccounts(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*MockDataService)(nil).GetAccounts), user)
}

// GetCards mocks base method.
func (m *MockDataService) GetCards(user client.User) ([]client.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCards", user)
	ret0, _ := ret[0].([]client.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCards indicates an expected call of GetCards.
func (mr *MockDataServiceMockRecorder) GetCards(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCards", reflect.TypeOf((*MockDataService)(nil).GetCards), user)
}

// GetNotes mocks base method.
func (m *MockDataService) GetNotes(user client.User) ([]client.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotes", user)
	ret0, _ := ret[0].([]client.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotes indicates an expected call of GetNotes.
func (mr *MockDataServiceMockRecorder) GetNotes(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotes", reflect.TypeOf((*MockDataService)(nil).GetNotes), user)
}

// GetUserByLogin mocks base method.
func (m *MockDataService) GetUserByLogin(login string) *client.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByLogin", login)
	ret0, _ := ret[0].(*client.User)
	return ret0
}

// GetUserByLogin indicates an expected call of GetUserByLogin.
func (mr *MockDataServiceMockRecorder) GetUserByLogin(login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByLogin", reflect.TypeOf((*MockDataService)(nil).GetUserByLogin), login)
}

// LoginUser mocks base method.
func (m *MockDataService) LoginUser(login, password string) (*client.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUser", login, password)
	ret0, _ := ret[0].(*client.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginUser indicates an expected call of LoginUser.
func (mr *MockDataServiceMockRecorder) LoginUser(login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUser", reflect.TypeOf((*MockDataService)(nil).LoginUser), login, password)
}

// RegisterUser mocks base method.
func (m *MockDataService) RegisterUser(login, password string) (*client.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", login, password)
	ret0, _ := ret[0].(*client.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockDataServiceMockRecorder) RegisterUser(login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockDataService)(nil).RegisterUser), login, password)
}

// SaveUser mocks base method.
func (m *MockDataService) SaveUser(user client.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveUser indicates an expected call of SaveUser.
func (mr *MockDataServiceMockRecorder) SaveUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUser", reflect.TypeOf((*MockDataService)(nil).SaveUser), user)
}

// UpdateAccount mocks base method.
func (m *MockDataService) UpdateAccount(user client.User, item client.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", user, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockDataServiceMockRecorder) UpdateAccount(user, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockDataService)(nil).UpdateAccount), user, item)
}

// UpdateCard mocks base method.
func (m *MockDataService) UpdateCard(user client.User, item client.Card) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCard", user, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCard indicates an expected call of UpdateCard.
func (mr *MockDataServiceMockRecorder) UpdateCard(user, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCard", reflect.TypeOf((*MockDataService)(nil).UpdateCard), user, item)
}

// UpdateNote mocks base method.
func (m *MockDataService) UpdateNote(user client.User, item client.Note) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNote", user, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNote indicates an expected call of UpdateNote.
func (mr *MockDataServiceMockRecorder) UpdateNote(user, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNote", reflect.TypeOf((*MockDataService)(nil).UpdateNote), user, item)
}