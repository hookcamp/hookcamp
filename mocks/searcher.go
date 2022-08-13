// Code generated by MockGen. DO NOT EDIT.
// Source: searcher/searcher.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	datastore "github.com/frain-dev/convoy/datastore"
	gomock "github.com/golang/mock/gomock"
)

// MockSearcher is a mock of Searcher interface.
type MockSearcher struct {
	ctrl     *gomock.Controller
	recorder *MockSearcherMockRecorder
}

// MockSearcherMockRecorder is the mock recorder for MockSearcher.
type MockSearcherMockRecorder struct {
	mock *MockSearcher
}

// NewMockSearcher creates a new mock instance.
func NewMockSearcher(ctrl *gomock.Controller) *MockSearcher {
	mock := &MockSearcher{ctrl: ctrl}
	mock.recorder = &MockSearcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearcher) EXPECT() *MockSearcherMockRecorder {
	return m.recorder
}

// Index mocks base method.
func (m *MockSearcher) Index(collection string, document interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Index", collection, document)
	ret0, _ := ret[0].(error)
	return ret0
}

// Index indicates an expected call of Index.
func (mr *MockSearcherMockRecorder) Index(collection, document interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Index", reflect.TypeOf((*MockSearcher)(nil).Index), collection, document)
}

// Remove mocks base method.
func (m *MockSearcher) Remove(collection string, filter *datastore.Filter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", collection, filter)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockSearcherMockRecorder) Remove(collection, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockSearcher)(nil).Remove), collection, filter)
}

// Search mocks base method.
func (m *MockSearcher) Search(collection string, filter *datastore.Filter) ([]string, datastore.PaginationData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", collection, filter)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(datastore.PaginationData)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Search indicates an expected call of Search.
func (mr *MockSearcherMockRecorder) Search(collection, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockSearcher)(nil).Search), collection, filter)
}
