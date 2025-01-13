// Code generated by MockGen. DO NOT EDIT.
// Source: /opt/gophkeeper/internal/client/interfaces/core_interface.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	client "github.com/besean163/gophkeeper/internal/models/client"
	gomock "github.com/golang/mock/gomock"
)

// MockCore is a mock of Core interface.
type MockCore struct {
	ctrl     *gomock.Controller
	recorder *MockCoreMockRecorder
}

// MockCoreMockRecorder is the mock recorder for MockCore.
type MockCoreMockRecorder struct {
	mock *MockCore
}

// NewMockCore creates a new mock instance.
func NewMockCore(ctrl *gomock.Controller) *MockCore {
	mock := &MockCore{ctrl: ctrl}
	mock.recorder = &MockCoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCore) EXPECT() *MockCoreMockRecorder {
	return m.recorder
}

// DeleteAccount mocks base method.
func (m *MockCore) DeleteAccount(item client.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", item)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockCoreMockRecorder) DeleteAccount(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockCore)(nil).DeleteAccount), item)
}

// DeleteCard mocks base method.
func (m *MockCore) DeleteCard(item client.Card) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCard", item)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCard indicates an expected call of DeleteCard.
func (mr *MockCoreMockRecorder) DeleteCard(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCard", reflect.TypeOf((*MockCore)(nil).DeleteCard), item)
}

// DeleteNote mocks base method.
func (m *MockCore) DeleteNote(item client.Note) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNote", item)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNote indicates an expected call of DeleteNote.
func (mr *MockCoreMockRecorder) DeleteNote(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNote", reflect.TypeOf((*MockCore)(nil).DeleteNote), item)
}

// GetAccounts mocks base method.
func (m *MockCore) GetAccounts() ([]client.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccounts")
	ret0, _ := ret[0].([]client.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccounts indicates an expected call of GetAccounts.
func (mr *MockCoreMockRecorder) GetAccounts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*MockCore)(nil).GetAccounts))
}

// GetCards mocks base method.
func (m *MockCore) GetCards() ([]client.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCards")
	ret0, _ := ret[0].([]client.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCards indicates an expected call of GetCards.
func (mr *MockCoreMockRecorder) GetCards() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCards", reflect.TypeOf((*MockCore)(nil).GetCards))
}

// GetNotes mocks base method.
func (m *MockCore) GetNotes() ([]client.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotes")
	ret0, _ := ret[0].([]client.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotes indicates an expected call of GetNotes.
func (mr *MockCoreMockRecorder) GetNotes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotes", reflect.TypeOf((*MockCore)(nil).GetNotes))
}

// Login mocks base method.
func (m *MockCore) Login(login, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", login, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// Login indicates an expected call of Login.
func (mr *MockCoreMockRecorder) Login(login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockCore)(nil).Login), login, password)
}

// Register mocks base method.
func (m *MockCore) Register(login, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", login, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockCoreMockRecorder) Register(login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockCore)(nil).Register), login, password)
}

// SaveAccount mocks base method.
func (m *MockCore) SaveAccount(item client.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveAccount", item)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveAccount indicates an expected call of SaveAccount.
func (mr *MockCoreMockRecorder) SaveAccount(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveAccount", reflect.TypeOf((*MockCore)(nil).SaveAccount), item)
}

// SaveCard mocks base method.
func (m *MockCore) SaveCard(item client.Card) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCard", item)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveCard indicates an expected call of SaveCard.
func (mr *MockCoreMockRecorder) SaveCard(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCard", reflect.TypeOf((*MockCore)(nil).SaveCard), item)
}

// SaveNote mocks base method.
func (m *MockCore) SaveNote(item client.Note) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveNote", item)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveNote indicates an expected call of SaveNote.
func (mr *MockCoreMockRecorder) SaveNote(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveNote", reflect.TypeOf((*MockCore)(nil).SaveNote), item)
}
