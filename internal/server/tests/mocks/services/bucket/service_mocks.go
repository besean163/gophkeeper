// Code generated by MockGen. DO NOT EDIT.
// Source: /opt/gophkeeper/internal/server/services/bucket/service.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	server "github.com/besean163/gophkeeper/internal/models/server"
	changes "github.com/besean163/gophkeeper/internal/server/services/bucket/changes"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// DeleteAccount mocks base method.
func (m *MockRepository) DeleteAccount(uuid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", uuid)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockRepositoryMockRecorder) DeleteAccount(uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockRepository)(nil).DeleteAccount), uuid)
}

// DeleteCard mocks base method.
func (m *MockRepository) DeleteCard(uuid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCard", uuid)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCard indicates an expected call of DeleteCard.
func (mr *MockRepositoryMockRecorder) DeleteCard(uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCard", reflect.TypeOf((*MockRepository)(nil).DeleteCard), uuid)
}

// DeleteNote mocks base method.
func (m *MockRepository) DeleteNote(uuid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNote", uuid)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNote indicates an expected call of DeleteNote.
func (mr *MockRepositoryMockRecorder) DeleteNote(uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNote", reflect.TypeOf((*MockRepository)(nil).DeleteNote), uuid)
}

// GetAccount mocks base method.
func (m *MockRepository) GetAccount(uuid string) (*server.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", uuid)
	ret0, _ := ret[0].(*server.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockRepositoryMockRecorder) GetAccount(uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockRepository)(nil).GetAccount), uuid)
}

// GetAccounts mocks base method.
func (m *MockRepository) GetAccounts(user server.User) ([]*server.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccounts", user)
	ret0, _ := ret[0].([]*server.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccounts indicates an expected call of GetAccounts.
func (mr *MockRepositoryMockRecorder) GetAccounts(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*MockRepository)(nil).GetAccounts), user)
}

// GetCard mocks base method.
func (m *MockRepository) GetCard(uuid string) (*server.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCard", uuid)
	ret0, _ := ret[0].(*server.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCard indicates an expected call of GetCard.
func (mr *MockRepositoryMockRecorder) GetCard(uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCard", reflect.TypeOf((*MockRepository)(nil).GetCard), uuid)
}

// GetCards mocks base method.
func (m *MockRepository) GetCards(user server.User) ([]*server.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCards", user)
	ret0, _ := ret[0].([]*server.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCards indicates an expected call of GetCards.
func (mr *MockRepositoryMockRecorder) GetCards(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCards", reflect.TypeOf((*MockRepository)(nil).GetCards), user)
}

// GetNote mocks base method.
func (m *MockRepository) GetNote(uuid string) (*server.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNote", uuid)
	ret0, _ := ret[0].(*server.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNote indicates an expected call of GetNote.
func (mr *MockRepositoryMockRecorder) GetNote(uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNote", reflect.TypeOf((*MockRepository)(nil).GetNote), uuid)
}

// GetNotes mocks base method.
func (m *MockRepository) GetNotes(user server.User) ([]*server.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotes", user)
	ret0, _ := ret[0].([]*server.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotes indicates an expected call of GetNotes.
func (mr *MockRepositoryMockRecorder) GetNotes(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotes", reflect.TypeOf((*MockRepository)(nil).GetNotes), user)
}

// SaveAccount mocks base method.
func (m *MockRepository) SaveAccount(account *server.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveAccount", account)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveAccount indicates an expected call of SaveAccount.
func (mr *MockRepositoryMockRecorder) SaveAccount(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveAccount", reflect.TypeOf((*MockRepository)(nil).SaveAccount), account)
}

// SaveCard mocks base method.
func (m *MockRepository) SaveCard(card *server.Card) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCard", card)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveCard indicates an expected call of SaveCard.
func (mr *MockRepositoryMockRecorder) SaveCard(card interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCard", reflect.TypeOf((*MockRepository)(nil).SaveCard), card)
}

// SaveNote mocks base method.
func (m *MockRepository) SaveNote(note *server.Note) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveNote", note)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveNote indicates an expected call of SaveNote.
func (mr *MockRepositoryMockRecorder) SaveNote(note interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveNote", reflect.TypeOf((*MockRepository)(nil).SaveNote), note)
}

// MockChangeDetector is a mock of ChangeDetector interface.
type MockChangeDetector struct {
	ctrl     *gomock.Controller
	recorder *MockChangeDetectorMockRecorder
}

// MockChangeDetectorMockRecorder is the mock recorder for MockChangeDetector.
type MockChangeDetectorMockRecorder struct {
	mock *MockChangeDetector
}

// NewMockChangeDetector creates a new mock instance.
func NewMockChangeDetector(ctrl *gomock.Controller) *MockChangeDetector {
	mock := &MockChangeDetector{ctrl: ctrl}
	mock.recorder = &MockChangeDetectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChangeDetector) EXPECT() *MockChangeDetectorMockRecorder {
	return m.recorder
}

// GetAccountChanges mocks base method.
func (m *MockChangeDetector) GetAccountChanges(user server.User, compare changes.AccountCompare) changes.AccountChanges {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountChanges", user, compare)
	ret0, _ := ret[0].(changes.AccountChanges)
	return ret0
}

// GetAccountChanges indicates an expected call of GetAccountChanges.
func (mr *MockChangeDetectorMockRecorder) GetAccountChanges(user, compare interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountChanges", reflect.TypeOf((*MockChangeDetector)(nil).GetAccountChanges), user, compare)
}

// GetCardsChanges mocks base method.
func (m *MockChangeDetector) GetCardsChanges(user server.User, compare changes.CardCompare) changes.CardChanges {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCardsChanges", user, compare)
	ret0, _ := ret[0].(changes.CardChanges)
	return ret0
}

// GetCardsChanges indicates an expected call of GetCardsChanges.
func (mr *MockChangeDetectorMockRecorder) GetCardsChanges(user, compare interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCardsChanges", reflect.TypeOf((*MockChangeDetector)(nil).GetCardsChanges), user, compare)
}

// GetNotesChanges mocks base method.
func (m *MockChangeDetector) GetNotesChanges(user server.User, compare changes.NoteCompare) changes.NoteChanges {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotesChanges", user, compare)
	ret0, _ := ret[0].(changes.NoteChanges)
	return ret0
}

// GetNotesChanges indicates an expected call of GetNotesChanges.
func (mr *MockChangeDetectorMockRecorder) GetNotesChanges(user, compare interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotesChanges", reflect.TypeOf((*MockChangeDetector)(nil).GetNotesChanges), user, compare)
}
