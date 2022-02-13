// Code generated by MockGen. DO NOT EDIT.
// Source: .\account\account.go

// Package mock is a generated GoMock package.
package mock

import (
	account "my_account/account"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIAccountService is a mock of IAccountService interface.
type MockIAccountService struct {
	ctrl     *gomock.Controller
	recorder *MockIAccountServiceMockRecorder
}

// MockIAccountServiceMockRecorder is the mock recorder for MockIAccountService.
type MockIAccountServiceMockRecorder struct {
	mock *MockIAccountService
}

// NewMockIAccountService creates a new mock instance.
func NewMockIAccountService(ctrl *gomock.Controller) *MockIAccountService {
	mock := &MockIAccountService{ctrl: ctrl}
	mock.recorder = &MockIAccountServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAccountService) EXPECT() *MockIAccountServiceMockRecorder {
	return m.recorder
}

// AccountOf mocks base method.
func (m *MockIAccountService) AccountOf(username string) *account.Account {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccountOf", username)
	ret0, _ := ret[0].(*account.Account)
	return ret0
}

// AccountOf indicates an expected call of AccountOf.
func (mr *MockIAccountServiceMockRecorder) AccountOf(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountOf", reflect.TypeOf((*MockIAccountService)(nil).AccountOf), username)
}

// Accounts mocks base method.
func (m *MockIAccountService) Accounts() []*account.Account {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accounts")
	ret0, _ := ret[0].([]*account.Account)
	return ret0
}

// Accounts indicates an expected call of Accounts.
func (mr *MockIAccountServiceMockRecorder) Accounts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accounts", reflect.TypeOf((*MockIAccountService)(nil).Accounts))
}

// Put mocks base method.
func (m *MockIAccountService) Put(username string) (float32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", username)
	ret0, _ := ret[0].(float32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockIAccountServiceMockRecorder) Put(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockIAccountService)(nil).Put), username)
}

// UpdateBalance mocks base method.
func (m *MockIAccountService) UpdateBalance(username string, amount float32) (float32, float32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBalance", username, amount)
	ret0, _ := ret[0].(float32)
	ret1, _ := ret[1].(float32)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateBalance indicates an expected call of UpdateBalance.
func (mr *MockIAccountServiceMockRecorder) UpdateBalance(username, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBalance", reflect.TypeOf((*MockIAccountService)(nil).UpdateBalance), username, amount)
}

// MockIAccountRepository is a mock of IAccountRepository interface.
type MockIAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIAccountRepositoryMockRecorder
}

// MockIAccountRepositoryMockRecorder is the mock recorder for MockIAccountRepository.
type MockIAccountRepositoryMockRecorder struct {
	mock *MockIAccountRepository
}

// NewMockIAccountRepository creates a new mock instance.
func NewMockIAccountRepository(ctrl *gomock.Controller) *MockIAccountRepository {
	mock := &MockIAccountRepository{ctrl: ctrl}
	mock.recorder = &MockIAccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAccountRepository) EXPECT() *MockIAccountRepositoryMockRecorder {
	return m.recorder
}

// AccountOf mocks base method.
func (m *MockIAccountRepository) AccountOf(username string) *account.Account {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccountOf", username)
	ret0, _ := ret[0].(*account.Account)
	return ret0
}

// AccountOf indicates an expected call of AccountOf.
func (mr *MockIAccountRepositoryMockRecorder) AccountOf(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountOf", reflect.TypeOf((*MockIAccountRepository)(nil).AccountOf), username)
}

// Accounts mocks base method.
func (m *MockIAccountRepository) Accounts() []*account.Account {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accounts")
	ret0, _ := ret[0].([]*account.Account)
	return ret0
}

// Accounts indicates an expected call of Accounts.
func (mr *MockIAccountRepositoryMockRecorder) Accounts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accounts", reflect.TypeOf((*MockIAccountRepository)(nil).Accounts))
}

// Count mocks base method.
func (m *MockIAccountRepository) Count() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count")
	ret0, _ := ret[0].(int)
	return ret0
}

// Count indicates an expected call of Count.
func (mr *MockIAccountRepositoryMockRecorder) Count() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockIAccountRepository)(nil).Count))
}

// Put mocks base method.
func (m *MockIAccountRepository) Put(username string, initBalance float32) (float32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", username, initBalance)
	ret0, _ := ret[0].(float32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockIAccountRepositoryMockRecorder) Put(username, initBalance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockIAccountRepository)(nil).Put), username, initBalance)
}

// UpdateBalance mocks base method.
func (m *MockIAccountRepository) UpdateBalance(username string, amount float32) (float32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBalance", username, amount)
	ret0, _ := ret[0].(float32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBalance indicates an expected call of UpdateBalance.
func (mr *MockIAccountRepositoryMockRecorder) UpdateBalance(username, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBalance", reflect.TypeOf((*MockIAccountRepository)(nil).UpdateBalance), username, amount)
}
