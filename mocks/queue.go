// Code generated by MockGen. DO NOT EDIT.
// Source: queue/queue.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	convoy "github.com/frain-dev/convoy"
	queue "github.com/frain-dev/convoy/queue"
	gomock "github.com/golang/mock/gomock"
)

// MockQueuer is a mock of Queuer interface.
type MockQueuer struct {
	ctrl     *gomock.Controller
	recorder *MockQueuerMockRecorder
}

// MockQueuerMockRecorder is the mock recorder for MockQueuer.
type MockQueuerMockRecorder struct {
	mock *MockQueuer
}

// NewMockQueuer creates a new mock instance.
func NewMockQueuer(ctrl *gomock.Controller) *MockQueuer {
	mock := &MockQueuer{ctrl: ctrl}
	mock.recorder = &MockQueuerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueuer) EXPECT() *MockQueuerMockRecorder {
	return m.recorder
}

// Options mocks base method.
func (m *MockQueuer) Options() queue.QueueOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Options")
	ret0, _ := ret[0].(queue.QueueOptions)
	return ret0
}

// Options indicates an expected call of Options.
func (mr *MockQueuerMockRecorder) Options() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Options", reflect.TypeOf((*MockQueuer)(nil).Options))
}

// Write mocks base method.
func (m *MockQueuer) Write(arg0 convoy.TaskName, arg1 convoy.QueueName, arg2 *queue.Job) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockQueuerMockRecorder) Write(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockQueuer)(nil).Write), arg0, arg1, arg2)
}
