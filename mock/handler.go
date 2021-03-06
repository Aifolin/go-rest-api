// Code generated by MockGen. DO NOT EDIT.
// Source: handler.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	handler "github.com/Toshik1978/go-rest-api/handler"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAccountBuilder is a mock of AccountBuilder interface
type MockAccountBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockAccountBuilderMockRecorder
}

// MockAccountBuilderMockRecorder is the mock recorder for MockAccountBuilder
type MockAccountBuilderMockRecorder struct {
	mock *MockAccountBuilder
}

// NewMockAccountBuilder creates a new mock instance
func NewMockAccountBuilder(ctrl *gomock.Controller) *MockAccountBuilder {
	mock := &MockAccountBuilder{ctrl: ctrl}
	mock.recorder = &MockAccountBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountBuilder) EXPECT() *MockAccountBuilderMockRecorder {
	return m.recorder
}

// SetUID mocks base method
func (m *MockAccountBuilder) SetUID(uid string) handler.AccountBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUID", uid)
	ret0, _ := ret[0].(handler.AccountBuilder)
	return ret0
}

// SetUID indicates an expected call of SetUID
func (mr *MockAccountBuilderMockRecorder) SetUID(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUID", reflect.TypeOf((*MockAccountBuilder)(nil).SetUID), uid)
}

// SetBalance mocks base method
func (m *MockAccountBuilder) SetBalance(balance float64) handler.AccountBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBalance", balance)
	ret0, _ := ret[0].(handler.AccountBuilder)
	return ret0
}

// SetBalance indicates an expected call of SetBalance
func (mr *MockAccountBuilderMockRecorder) SetBalance(balance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBalance", reflect.TypeOf((*MockAccountBuilder)(nil).SetBalance), balance)
}

// SetCurrency mocks base method
func (m *MockAccountBuilder) SetCurrency(currency string) handler.AccountBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCurrency", currency)
	ret0, _ := ret[0].(handler.AccountBuilder)
	return ret0
}

// SetCurrency indicates an expected call of SetCurrency
func (mr *MockAccountBuilderMockRecorder) SetCurrency(currency interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCurrency", reflect.TypeOf((*MockAccountBuilder)(nil).SetCurrency), currency)
}

// Build mocks base method
func (m *MockAccountBuilder) Build(ctx context.Context) (*handler.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", ctx)
	ret0, _ := ret[0].(*handler.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Build indicates an expected call of Build
func (mr *MockAccountBuilderMockRecorder) Build(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockAccountBuilder)(nil).Build), ctx)
}

// MockPaymentBuilder is a mock of PaymentBuilder interface
type MockPaymentBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentBuilderMockRecorder
}

// MockPaymentBuilderMockRecorder is the mock recorder for MockPaymentBuilder
type MockPaymentBuilderMockRecorder struct {
	mock *MockPaymentBuilder
}

// NewMockPaymentBuilder creates a new mock instance
func NewMockPaymentBuilder(ctrl *gomock.Controller) *MockPaymentBuilder {
	mock := &MockPaymentBuilder{ctrl: ctrl}
	mock.recorder = &MockPaymentBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPaymentBuilder) EXPECT() *MockPaymentBuilderMockRecorder {
	return m.recorder
}

// SetAmount mocks base method
func (m *MockPaymentBuilder) SetAmount(amount float64) handler.PaymentBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAmount", amount)
	ret0, _ := ret[0].(handler.PaymentBuilder)
	return ret0
}

// SetAmount indicates an expected call of SetAmount
func (mr *MockPaymentBuilderMockRecorder) SetAmount(amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAmount", reflect.TypeOf((*MockPaymentBuilder)(nil).SetAmount), amount)
}

// SetPayer mocks base method
func (m *MockPaymentBuilder) SetPayer(uid string) handler.PaymentBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPayer", uid)
	ret0, _ := ret[0].(handler.PaymentBuilder)
	return ret0
}

// SetPayer indicates an expected call of SetPayer
func (mr *MockPaymentBuilderMockRecorder) SetPayer(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPayer", reflect.TypeOf((*MockPaymentBuilder)(nil).SetPayer), uid)
}

// SetRecipient mocks base method
func (m *MockPaymentBuilder) SetRecipient(uid string) handler.PaymentBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRecipient", uid)
	ret0, _ := ret[0].(handler.PaymentBuilder)
	return ret0
}

// SetRecipient indicates an expected call of SetRecipient
func (mr *MockPaymentBuilderMockRecorder) SetRecipient(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRecipient", reflect.TypeOf((*MockPaymentBuilder)(nil).SetRecipient), uid)
}

// Build mocks base method
func (m *MockPaymentBuilder) Build(ctx context.Context) (*handler.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", ctx)
	ret0, _ := ret[0].(*handler.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Build indicates an expected call of Build
func (mr *MockPaymentBuilderMockRecorder) Build(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockPaymentBuilder)(nil).Build), ctx)
}

// MockAccountManager is a mock of AccountManager interface
type MockAccountManager struct {
	ctrl     *gomock.Controller
	recorder *MockAccountManagerMockRecorder
}

// MockAccountManagerMockRecorder is the mock recorder for MockAccountManager
type MockAccountManagerMockRecorder struct {
	mock *MockAccountManager
}

// NewMockAccountManager creates a new mock instance
func NewMockAccountManager(ctrl *gomock.Controller) *MockAccountManager {
	mock := &MockAccountManager{ctrl: ctrl}
	mock.recorder = &MockAccountManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountManager) EXPECT() *MockAccountManagerMockRecorder {
	return m.recorder
}

// AllAccounts mocks base method
func (m *MockAccountManager) AllAccounts(ctx context.Context) ([]handler.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllAccounts", ctx)
	ret0, _ := ret[0].([]handler.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllAccounts indicates an expected call of AllAccounts
func (mr *MockAccountManagerMockRecorder) AllAccounts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllAccounts", reflect.TypeOf((*MockAccountManager)(nil).AllAccounts), ctx)
}

// AllPayments mocks base method
func (m *MockAccountManager) AllPayments(ctx context.Context) ([]handler.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllPayments", ctx)
	ret0, _ := ret[0].([]handler.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllPayments indicates an expected call of AllPayments
func (mr *MockAccountManagerMockRecorder) AllPayments(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllPayments", reflect.TypeOf((*MockAccountManager)(nil).AllPayments), ctx)
}

// AccountBuilder mocks base method
func (m *MockAccountManager) AccountBuilder() handler.AccountBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccountBuilder")
	ret0, _ := ret[0].(handler.AccountBuilder)
	return ret0
}

// AccountBuilder indicates an expected call of AccountBuilder
func (mr *MockAccountManagerMockRecorder) AccountBuilder() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountBuilder", reflect.TypeOf((*MockAccountManager)(nil).AccountBuilder))
}

// PaymentBuilder mocks base method
func (m *MockAccountManager) PaymentBuilder() handler.PaymentBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaymentBuilder")
	ret0, _ := ret[0].(handler.PaymentBuilder)
	return ret0
}

// PaymentBuilder indicates an expected call of PaymentBuilder
func (mr *MockAccountManagerMockRecorder) PaymentBuilder() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaymentBuilder", reflect.TypeOf((*MockAccountManager)(nil).PaymentBuilder))
}
