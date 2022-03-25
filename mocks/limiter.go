// Code generated by MockGen. DO NOT EDIT.
// Source: limiter/limiter.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	redis_rate "github.com/go-redis/redis_rate/v9"
	gomock "github.com/golang/mock/gomock"
)

// MockRateLimiter is a mock of RateLimiter interface.
type MockRateLimiter struct {
	ctrl     *gomock.Controller
	recorder *MockRateLimiterMockRecorder
}

// MockRateLimiterMockRecorder is the mock recorder for MockRateLimiter.
type MockRateLimiterMockRecorder struct {
	mock *MockRateLimiter
}

// NewMockRateLimiter creates a new mock instance.
func NewMockRateLimiter(ctrl *gomock.Controller) *MockRateLimiter {
	mock := &MockRateLimiter{ctrl: ctrl}
	mock.recorder = &MockRateLimiterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRateLimiter) EXPECT() *MockRateLimiterMockRecorder {
	return m.recorder
}

// Allow mocks base method.
func (m *MockRateLimiter) Allow(ctx context.Context, key string, limit, duration int) (*redis_rate.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Allow", ctx, key, limit, duration)
	ret0, _ := ret[0].(*redis_rate.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Allow indicates an expected call of Allow.
func (mr *MockRateLimiterMockRecorder) Allow(ctx, key, limit, duration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Allow", reflect.TypeOf((*MockRateLimiter)(nil).Allow), ctx, key, limit, duration)
}

// ShouldAllow mocks base method.
func (m *MockRateLimiter) ShouldAllow(ctx context.Context, key string, limit, duration int) (*redis_rate.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldAllow", ctx, key, limit, duration)
	ret0, _ := ret[0].(*redis_rate.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShouldAllow indicates an expected call of ShouldAllow.
func (mr *MockRateLimiterMockRecorder) ShouldAllow(ctx, key, limit, duration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldAllow", reflect.TypeOf((*MockRateLimiter)(nil).ShouldAllow), ctx, key, limit, duration)
}
