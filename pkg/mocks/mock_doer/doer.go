// Code generated by MockGen. DO NOT EDIT.
// Source: doer.go

// Package mock_doer is a generated GoMock package.
package mock_doer

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDoer is a mock of Doer interface
type MockDoer struct {
	ctrl     *gomock.Controller
	recorder *MockDoerMockRecorder
}

// MockDoerMockRecorder is the mock recorder for MockDoer
type MockDoerMockRecorder struct {
	mock *MockDoer
}

// NewMockDoer creates a new mock instance
func NewMockDoer(ctrl *gomock.Controller) *MockDoer {
	mock := &MockDoer{ctrl: ctrl}
	mock.recorder = &MockDoerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDoer) EXPECT() *MockDoerMockRecorder {
	return m.recorder
}

// DoSomething mocks base method
func (m *MockDoer) DoSomething(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoSomething", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoSomething indicates an expected call of DoSomething
func (mr *MockDoerMockRecorder) DoSomething(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoSomething", reflect.TypeOf((*MockDoer)(nil).DoSomething), arg0, arg1)
}

// SaySomething mocks base method
func (m *MockDoer) SaySomething(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaySomething", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaySomething indicates an expected call of SaySomething
func (mr *MockDoerMockRecorder) SaySomething(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaySomething", reflect.TypeOf((*MockDoer)(nil).SaySomething), arg0)
}

// KickSomething mocks base method
func (m *MockDoer) KickSomething(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KickSomething", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// KickSomething indicates an expected call of KickSomething
func (mr *MockDoerMockRecorder) KickSomething(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KickSomething", reflect.TypeOf((*MockDoer)(nil).KickSomething), arg0)
}

// SingSomething mocks base method
func (m *MockDoer) SingSomething(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SingSomething", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SingSomething indicates an expected call of SingSomething
func (mr *MockDoerMockRecorder) SingSomething(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SingSomething", reflect.TypeOf((*MockDoer)(nil).SingSomething), arg0)
}
