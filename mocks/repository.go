// Code generated by MockGen. DO NOT EDIT.
// Source: datastore/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	datastore "github.com/frain-dev/convoy/datastore"
	gomock "github.com/golang/mock/gomock"
)

// MockAPIKeyRepository is a mock of APIKeyRepository interface.
type MockAPIKeyRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAPIKeyRepositoryMockRecorder
}

// MockAPIKeyRepositoryMockRecorder is the mock recorder for MockAPIKeyRepository.
type MockAPIKeyRepositoryMockRecorder struct {
	mock *MockAPIKeyRepository
}

// NewMockAPIKeyRepository creates a new mock instance.
func NewMockAPIKeyRepository(ctrl *gomock.Controller) *MockAPIKeyRepository {
	mock := &MockAPIKeyRepository{ctrl: ctrl}
	mock.recorder = &MockAPIKeyRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIKeyRepository) EXPECT() *MockAPIKeyRepositoryMockRecorder {
	return m.recorder
}

// CreateAPIKey mocks base method.
func (m *MockAPIKeyRepository) CreateAPIKey(arg0 context.Context, arg1 *datastore.APIKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAPIKey", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAPIKey indicates an expected call of CreateAPIKey.
func (mr *MockAPIKeyRepositoryMockRecorder) CreateAPIKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAPIKey", reflect.TypeOf((*MockAPIKeyRepository)(nil).CreateAPIKey), arg0, arg1)
}

// FindAPIKeyByHash mocks base method.
func (m *MockAPIKeyRepository) FindAPIKeyByHash(arg0 context.Context, arg1 string) (*datastore.APIKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAPIKeyByHash", arg0, arg1)
	ret0, _ := ret[0].(*datastore.APIKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAPIKeyByHash indicates an expected call of FindAPIKeyByHash.
func (mr *MockAPIKeyRepositoryMockRecorder) FindAPIKeyByHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAPIKeyByHash", reflect.TypeOf((*MockAPIKeyRepository)(nil).FindAPIKeyByHash), arg0, arg1)
}

// FindAPIKeyByID mocks base method.
func (m *MockAPIKeyRepository) FindAPIKeyByID(arg0 context.Context, arg1 string) (*datastore.APIKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAPIKeyByID", arg0, arg1)
	ret0, _ := ret[0].(*datastore.APIKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAPIKeyByID indicates an expected call of FindAPIKeyByID.
func (mr *MockAPIKeyRepositoryMockRecorder) FindAPIKeyByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAPIKeyByID", reflect.TypeOf((*MockAPIKeyRepository)(nil).FindAPIKeyByID), arg0, arg1)
}

// FindAPIKeyByMaskID mocks base method.
func (m *MockAPIKeyRepository) FindAPIKeyByMaskID(arg0 context.Context, arg1 string) (*datastore.APIKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAPIKeyByMaskID", arg0, arg1)
	ret0, _ := ret[0].(*datastore.APIKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAPIKeyByMaskID indicates an expected call of FindAPIKeyByMaskID.
func (mr *MockAPIKeyRepositoryMockRecorder) FindAPIKeyByMaskID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAPIKeyByMaskID", reflect.TypeOf((*MockAPIKeyRepository)(nil).FindAPIKeyByMaskID), arg0, arg1)
}

// LoadAPIKeysPaged mocks base method.
func (m *MockAPIKeyRepository) LoadAPIKeysPaged(arg0 context.Context, arg1 *datastore.Pageable) ([]datastore.APIKey, datastore.PaginationData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadAPIKeysPaged", arg0, arg1)
	ret0, _ := ret[0].([]datastore.APIKey)
	ret1, _ := ret[1].(datastore.PaginationData)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LoadAPIKeysPaged indicates an expected call of LoadAPIKeysPaged.
func (mr *MockAPIKeyRepositoryMockRecorder) LoadAPIKeysPaged(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadAPIKeysPaged", reflect.TypeOf((*MockAPIKeyRepository)(nil).LoadAPIKeysPaged), arg0, arg1)
}

// RevokeAPIKeys mocks base method.
func (m *MockAPIKeyRepository) RevokeAPIKeys(arg0 context.Context, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeAPIKeys", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeAPIKeys indicates an expected call of RevokeAPIKeys.
func (mr *MockAPIKeyRepositoryMockRecorder) RevokeAPIKeys(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeAPIKeys", reflect.TypeOf((*MockAPIKeyRepository)(nil).RevokeAPIKeys), arg0, arg1)
}

// UpdateAPIKey mocks base method.
func (m *MockAPIKeyRepository) UpdateAPIKey(arg0 context.Context, arg1 *datastore.APIKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAPIKey", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAPIKey indicates an expected call of UpdateAPIKey.
func (mr *MockAPIKeyRepositoryMockRecorder) UpdateAPIKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAPIKey", reflect.TypeOf((*MockAPIKeyRepository)(nil).UpdateAPIKey), arg0, arg1)
}

// MockEventDeliveryRepository is a mock of EventDeliveryRepository interface.
type MockEventDeliveryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEventDeliveryRepositoryMockRecorder
}

// MockEventDeliveryRepositoryMockRecorder is the mock recorder for MockEventDeliveryRepository.
type MockEventDeliveryRepositoryMockRecorder struct {
	mock *MockEventDeliveryRepository
}

// NewMockEventDeliveryRepository creates a new mock instance.
func NewMockEventDeliveryRepository(ctrl *gomock.Controller) *MockEventDeliveryRepository {
	mock := &MockEventDeliveryRepository{ctrl: ctrl}
	mock.recorder = &MockEventDeliveryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventDeliveryRepository) EXPECT() *MockEventDeliveryRepositoryMockRecorder {
	return m.recorder
}

// CreateEventDelivery mocks base method.
func (m *MockEventDeliveryRepository) CreateEventDelivery(arg0 context.Context, arg1 *datastore.EventDelivery) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEventDelivery", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEventDelivery indicates an expected call of CreateEventDelivery.
func (mr *MockEventDeliveryRepositoryMockRecorder) CreateEventDelivery(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEventDelivery", reflect.TypeOf((*MockEventDeliveryRepository)(nil).CreateEventDelivery), arg0, arg1)
}

// FindEventDeliveriesByEventID mocks base method.
func (m *MockEventDeliveryRepository) FindEventDeliveriesByEventID(arg0 context.Context, arg1 string) ([]datastore.EventDelivery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEventDeliveriesByEventID", arg0, arg1)
	ret0, _ := ret[0].([]datastore.EventDelivery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEventDeliveriesByEventID indicates an expected call of FindEventDeliveriesByEventID.
func (mr *MockEventDeliveryRepositoryMockRecorder) FindEventDeliveriesByEventID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEventDeliveriesByEventID", reflect.TypeOf((*MockEventDeliveryRepository)(nil).FindEventDeliveriesByEventID), arg0, arg1)
}

// FindEventDeliveriesByIDs mocks base method.
func (m *MockEventDeliveryRepository) FindEventDeliveriesByIDs(arg0 context.Context, arg1 []string) ([]datastore.EventDelivery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEventDeliveriesByIDs", arg0, arg1)
	ret0, _ := ret[0].([]datastore.EventDelivery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEventDeliveriesByIDs indicates an expected call of FindEventDeliveriesByIDs.
func (mr *MockEventDeliveryRepositoryMockRecorder) FindEventDeliveriesByIDs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEventDeliveriesByIDs", reflect.TypeOf((*MockEventDeliveryRepository)(nil).FindEventDeliveriesByIDs), arg0, arg1)
}

// FindEventDeliveryByID mocks base method.
func (m *MockEventDeliveryRepository) FindEventDeliveryByID(arg0 context.Context, arg1 string) (*datastore.EventDelivery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEventDeliveryByID", arg0, arg1)
	ret0, _ := ret[0].(*datastore.EventDelivery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEventDeliveryByID indicates an expected call of FindEventDeliveryByID.
func (mr *MockEventDeliveryRepositoryMockRecorder) FindEventDeliveryByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEventDeliveryByID", reflect.TypeOf((*MockEventDeliveryRepository)(nil).FindEventDeliveryByID), arg0, arg1)
}

// LoadEventDeliveriesPaged mocks base method.
func (m *MockEventDeliveryRepository) LoadEventDeliveriesPaged(arg0 context.Context, arg1, arg2, arg3 string, arg4 []datastore.EventDeliveryStatus, arg5 datastore.SearchParams, arg6 datastore.Pageable) ([]datastore.EventDelivery, datastore.PaginationData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadEventDeliveriesPaged", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].([]datastore.EventDelivery)
	ret1, _ := ret[1].(datastore.PaginationData)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LoadEventDeliveriesPaged indicates an expected call of LoadEventDeliveriesPaged.
func (mr *MockEventDeliveryRepositoryMockRecorder) LoadEventDeliveriesPaged(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadEventDeliveriesPaged", reflect.TypeOf((*MockEventDeliveryRepository)(nil).LoadEventDeliveriesPaged), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// UpdateEventDeliveryWithAttempt mocks base method.
func (m *MockEventDeliveryRepository) UpdateEventDeliveryWithAttempt(arg0 context.Context, arg1 datastore.EventDelivery, arg2 datastore.DeliveryAttempt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEventDeliveryWithAttempt", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEventDeliveryWithAttempt indicates an expected call of UpdateEventDeliveryWithAttempt.
func (mr *MockEventDeliveryRepositoryMockRecorder) UpdateEventDeliveryWithAttempt(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEventDeliveryWithAttempt", reflect.TypeOf((*MockEventDeliveryRepository)(nil).UpdateEventDeliveryWithAttempt), arg0, arg1, arg2)
}

// UpdateStatusOfEventDelivery mocks base method.
func (m *MockEventDeliveryRepository) UpdateStatusOfEventDelivery(arg0 context.Context, arg1 datastore.EventDelivery, arg2 datastore.EventDeliveryStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatusOfEventDelivery", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatusOfEventDelivery indicates an expected call of UpdateStatusOfEventDelivery.
func (mr *MockEventDeliveryRepositoryMockRecorder) UpdateStatusOfEventDelivery(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatusOfEventDelivery", reflect.TypeOf((*MockEventDeliveryRepository)(nil).UpdateStatusOfEventDelivery), arg0, arg1, arg2)
}

// MockEventRepository is a mock of EventRepository interface.
type MockEventRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEventRepositoryMockRecorder
}

// MockEventRepositoryMockRecorder is the mock recorder for MockEventRepository.
type MockEventRepositoryMockRecorder struct {
	mock *MockEventRepository
}

// NewMockEventRepository creates a new mock instance.
func NewMockEventRepository(ctrl *gomock.Controller) *MockEventRepository {
	mock := &MockEventRepository{ctrl: ctrl}
	mock.recorder = &MockEventRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventRepository) EXPECT() *MockEventRepositoryMockRecorder {
	return m.recorder
}

// CountGroupMessages mocks base method.
func (m *MockEventRepository) CountGroupMessages(ctx context.Context, groupID string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountGroupMessages", ctx, groupID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountGroupMessages indicates an expected call of CountGroupMessages.
func (mr *MockEventRepositoryMockRecorder) CountGroupMessages(ctx, groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountGroupMessages", reflect.TypeOf((*MockEventRepository)(nil).CountGroupMessages), ctx, groupID)
}

// CreateEvent mocks base method.
func (m *MockEventRepository) CreateEvent(arg0 context.Context, arg1 *datastore.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEvent indicates an expected call of CreateEvent.
func (mr *MockEventRepositoryMockRecorder) CreateEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockEventRepository)(nil).CreateEvent), arg0, arg1)
}

// DeleteGroupEvents mocks base method.
func (m *MockEventRepository) DeleteGroupEvents(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGroupEvents", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGroupEvents indicates an expected call of DeleteGroupEvents.
func (mr *MockEventRepositoryMockRecorder) DeleteGroupEvents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroupEvents", reflect.TypeOf((*MockEventRepository)(nil).DeleteGroupEvents), arg0, arg1)
}

// FindEventByID mocks base method.
func (m *MockEventRepository) FindEventByID(ctx context.Context, id string) (*datastore.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEventByID", ctx, id)
	ret0, _ := ret[0].(*datastore.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEventByID indicates an expected call of FindEventByID.
func (mr *MockEventRepositoryMockRecorder) FindEventByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEventByID", reflect.TypeOf((*MockEventRepository)(nil).FindEventByID), ctx, id)
}

// LoadAbandonedEventsForPostingRetry mocks base method.
func (m *MockEventRepository) LoadAbandonedEventsForPostingRetry(arg0 context.Context) ([]datastore.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadAbandonedEventsForPostingRetry", arg0)
	ret0, _ := ret[0].([]datastore.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadAbandonedEventsForPostingRetry indicates an expected call of LoadAbandonedEventsForPostingRetry.
func (mr *MockEventRepositoryMockRecorder) LoadAbandonedEventsForPostingRetry(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadAbandonedEventsForPostingRetry", reflect.TypeOf((*MockEventRepository)(nil).LoadAbandonedEventsForPostingRetry), arg0)
}

// LoadEventIntervals mocks base method.
func (m *MockEventRepository) LoadEventIntervals(arg0 context.Context, arg1 string, arg2 datastore.SearchParams, arg3 datastore.Period, arg4 int) ([]datastore.EventInterval, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadEventIntervals", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]datastore.EventInterval)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadEventIntervals indicates an expected call of LoadEventIntervals.
func (mr *MockEventRepositoryMockRecorder) LoadEventIntervals(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadEventIntervals", reflect.TypeOf((*MockEventRepository)(nil).LoadEventIntervals), arg0, arg1, arg2, arg3, arg4)
}

// LoadEventsForPostingRetry mocks base method.
func (m *MockEventRepository) LoadEventsForPostingRetry(arg0 context.Context) ([]datastore.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadEventsForPostingRetry", arg0)
	ret0, _ := ret[0].([]datastore.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadEventsForPostingRetry indicates an expected call of LoadEventsForPostingRetry.
func (mr *MockEventRepositoryMockRecorder) LoadEventsForPostingRetry(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadEventsForPostingRetry", reflect.TypeOf((*MockEventRepository)(nil).LoadEventsForPostingRetry), arg0)
}

// LoadEventsPaged mocks base method.
func (m *MockEventRepository) LoadEventsPaged(arg0 context.Context, arg1, arg2 string, arg3 datastore.SearchParams, arg4 datastore.Pageable) ([]datastore.Event, datastore.PaginationData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadEventsPaged", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]datastore.Event)
	ret1, _ := ret[1].(datastore.PaginationData)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LoadEventsPaged indicates an expected call of LoadEventsPaged.
func (mr *MockEventRepositoryMockRecorder) LoadEventsPaged(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadEventsPaged", reflect.TypeOf((*MockEventRepository)(nil).LoadEventsPaged), arg0, arg1, arg2, arg3, arg4)
}

// LoadEventsPagedByAppId mocks base method.
func (m *MockEventRepository) LoadEventsPagedByAppId(arg0 context.Context, arg1 string, arg2 datastore.SearchParams, arg3 datastore.Pageable) ([]datastore.Event, datastore.PaginationData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadEventsPagedByAppId", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]datastore.Event)
	ret1, _ := ret[1].(datastore.PaginationData)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LoadEventsPagedByAppId indicates an expected call of LoadEventsPagedByAppId.
func (mr *MockEventRepositoryMockRecorder) LoadEventsPagedByAppId(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadEventsPagedByAppId", reflect.TypeOf((*MockEventRepository)(nil).LoadEventsPagedByAppId), arg0, arg1, arg2, arg3)
}

// LoadEventsScheduledForPosting mocks base method.
func (m *MockEventRepository) LoadEventsScheduledForPosting(arg0 context.Context) ([]datastore.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadEventsScheduledForPosting", arg0)
	ret0, _ := ret[0].([]datastore.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadEventsScheduledForPosting indicates an expected call of LoadEventsScheduledForPosting.
func (mr *MockEventRepositoryMockRecorder) LoadEventsScheduledForPosting(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadEventsScheduledForPosting", reflect.TypeOf((*MockEventRepository)(nil).LoadEventsScheduledForPosting), arg0)
}

// MockGroupRepository is a mock of GroupRepository interface.
type MockGroupRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGroupRepositoryMockRecorder
}

// MockGroupRepositoryMockRecorder is the mock recorder for MockGroupRepository.
type MockGroupRepositoryMockRecorder struct {
	mock *MockGroupRepository
}

// NewMockGroupRepository creates a new mock instance.
func NewMockGroupRepository(ctrl *gomock.Controller) *MockGroupRepository {
	mock := &MockGroupRepository{ctrl: ctrl}
	mock.recorder = &MockGroupRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupRepository) EXPECT() *MockGroupRepositoryMockRecorder {
	return m.recorder
}

// CreateGroup mocks base method.
func (m *MockGroupRepository) CreateGroup(arg0 context.Context, arg1 *datastore.Group) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateGroup indicates an expected call of CreateGroup.
func (mr *MockGroupRepositoryMockRecorder) CreateGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroup", reflect.TypeOf((*MockGroupRepository)(nil).CreateGroup), arg0, arg1)
}

// DeleteGroup mocks base method.
func (m *MockGroupRepository) DeleteGroup(ctx context.Context, uid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGroup", ctx, uid)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGroup indicates an expected call of DeleteGroup.
func (mr *MockGroupRepositoryMockRecorder) DeleteGroup(ctx, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroup", reflect.TypeOf((*MockGroupRepository)(nil).DeleteGroup), ctx, uid)
}

// FetchGroupByID mocks base method.
func (m *MockGroupRepository) FetchGroupByID(arg0 context.Context, arg1 string) (*datastore.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchGroupByID", arg0, arg1)
	ret0, _ := ret[0].(*datastore.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchGroupByID indicates an expected call of FetchGroupByID.
func (mr *MockGroupRepositoryMockRecorder) FetchGroupByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchGroupByID", reflect.TypeOf((*MockGroupRepository)(nil).FetchGroupByID), arg0, arg1)
}

// FetchGroupsByIDs mocks base method.
func (m *MockGroupRepository) FetchGroupsByIDs(arg0 context.Context, arg1 []string) ([]datastore.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchGroupsByIDs", arg0, arg1)
	ret0, _ := ret[0].([]datastore.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchGroupsByIDs indicates an expected call of FetchGroupsByIDs.
func (mr *MockGroupRepositoryMockRecorder) FetchGroupsByIDs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchGroupsByIDs", reflect.TypeOf((*MockGroupRepository)(nil).FetchGroupsByIDs), arg0, arg1)
}

// LoadGroups mocks base method.
func (m *MockGroupRepository) LoadGroups(arg0 context.Context, arg1 *datastore.GroupFilter) ([]*datastore.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadGroups", arg0, arg1)
	ret0, _ := ret[0].([]*datastore.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadGroups indicates an expected call of LoadGroups.
func (mr *MockGroupRepositoryMockRecorder) LoadGroups(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadGroups", reflect.TypeOf((*MockGroupRepository)(nil).LoadGroups), arg0, arg1)
}

// UpdateGroup mocks base method.
func (m *MockGroupRepository) UpdateGroup(arg0 context.Context, arg1 *datastore.Group) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGroup indicates an expected call of UpdateGroup.
func (mr *MockGroupRepositoryMockRecorder) UpdateGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGroup", reflect.TypeOf((*MockGroupRepository)(nil).UpdateGroup), arg0, arg1)
}

// MockApplicationRepository is a mock of ApplicationRepository interface.
type MockApplicationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationRepositoryMockRecorder
}

// MockApplicationRepositoryMockRecorder is the mock recorder for MockApplicationRepository.
type MockApplicationRepositoryMockRecorder struct {
	mock *MockApplicationRepository
}

// NewMockApplicationRepository creates a new mock instance.
func NewMockApplicationRepository(ctrl *gomock.Controller) *MockApplicationRepository {
	mock := &MockApplicationRepository{ctrl: ctrl}
	mock.recorder = &MockApplicationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationRepository) EXPECT() *MockApplicationRepositoryMockRecorder {
	return m.recorder
}

// CountGroupApplications mocks base method.
func (m *MockApplicationRepository) CountGroupApplications(ctx context.Context, groupID string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountGroupApplications", ctx, groupID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountGroupApplications indicates an expected call of CountGroupApplications.
func (mr *MockApplicationRepositoryMockRecorder) CountGroupApplications(ctx, groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountGroupApplications", reflect.TypeOf((*MockApplicationRepository)(nil).CountGroupApplications), ctx, groupID)
}

// CreateApplication mocks base method.
func (m *MockApplicationRepository) CreateApplication(arg0 context.Context, arg1 *datastore.Application) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateApplication", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateApplication indicates an expected call of CreateApplication.
func (mr *MockApplicationRepositoryMockRecorder) CreateApplication(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApplication", reflect.TypeOf((*MockApplicationRepository)(nil).CreateApplication), arg0, arg1)
}

// DeleteApplication mocks base method.
func (m *MockApplicationRepository) DeleteApplication(arg0 context.Context, arg1 *datastore.Application) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteApplication", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteApplication indicates an expected call of DeleteApplication.
func (mr *MockApplicationRepositoryMockRecorder) DeleteApplication(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApplication", reflect.TypeOf((*MockApplicationRepository)(nil).DeleteApplication), arg0, arg1)
}

// DeleteGroupApps mocks base method.
func (m *MockApplicationRepository) DeleteGroupApps(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGroupApps", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGroupApps indicates an expected call of DeleteGroupApps.
func (mr *MockApplicationRepositoryMockRecorder) DeleteGroupApps(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroupApps", reflect.TypeOf((*MockApplicationRepository)(nil).DeleteGroupApps), arg0, arg1)
}

// FindApplicationByID mocks base method.
func (m *MockApplicationRepository) FindApplicationByID(arg0 context.Context, arg1 string) (*datastore.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindApplicationByID", arg0, arg1)
	ret0, _ := ret[0].(*datastore.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindApplicationByID indicates an expected call of FindApplicationByID.
func (mr *MockApplicationRepositoryMockRecorder) FindApplicationByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindApplicationByID", reflect.TypeOf((*MockApplicationRepository)(nil).FindApplicationByID), arg0, arg1)
}

// FindApplicationEndpointByID mocks base method.
func (m *MockApplicationRepository) FindApplicationEndpointByID(arg0 context.Context, arg1, arg2 string) (*datastore.Endpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindApplicationEndpointByID", arg0, arg1, arg2)
	ret0, _ := ret[0].(*datastore.Endpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindApplicationEndpointByID indicates an expected call of FindApplicationEndpointByID.
func (mr *MockApplicationRepositoryMockRecorder) FindApplicationEndpointByID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindApplicationEndpointByID", reflect.TypeOf((*MockApplicationRepository)(nil).FindApplicationEndpointByID), arg0, arg1, arg2)
}

// LoadApplicationsPaged mocks base method.
func (m *MockApplicationRepository) LoadApplicationsPaged(arg0 context.Context, arg1, arg2 string, arg3 datastore.Pageable) ([]datastore.Application, datastore.PaginationData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadApplicationsPaged", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]datastore.Application)
	ret1, _ := ret[1].(datastore.PaginationData)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LoadApplicationsPaged indicates an expected call of LoadApplicationsPaged.
func (mr *MockApplicationRepositoryMockRecorder) LoadApplicationsPaged(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadApplicationsPaged", reflect.TypeOf((*MockApplicationRepository)(nil).LoadApplicationsPaged), arg0, arg1, arg2, arg3)
}

// LoadApplicationsPagedByGroupId mocks base method.
func (m *MockApplicationRepository) LoadApplicationsPagedByGroupId(arg0 context.Context, arg1 string, arg2 datastore.Pageable) ([]datastore.Application, datastore.PaginationData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadApplicationsPagedByGroupId", arg0, arg1, arg2)
	ret0, _ := ret[0].([]datastore.Application)
	ret1, _ := ret[1].(datastore.PaginationData)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LoadApplicationsPagedByGroupId indicates an expected call of LoadApplicationsPagedByGroupId.
func (mr *MockApplicationRepositoryMockRecorder) LoadApplicationsPagedByGroupId(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadApplicationsPagedByGroupId", reflect.TypeOf((*MockApplicationRepository)(nil).LoadApplicationsPagedByGroupId), arg0, arg1, arg2)
}

// SearchApplicationsByGroupId mocks base method.
func (m *MockApplicationRepository) SearchApplicationsByGroupId(arg0 context.Context, arg1 string, arg2 datastore.SearchParams) ([]datastore.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchApplicationsByGroupId", arg0, arg1, arg2)
	ret0, _ := ret[0].([]datastore.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchApplicationsByGroupId indicates an expected call of SearchApplicationsByGroupId.
func (mr *MockApplicationRepositoryMockRecorder) SearchApplicationsByGroupId(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchApplicationsByGroupId", reflect.TypeOf((*MockApplicationRepository)(nil).SearchApplicationsByGroupId), arg0, arg1, arg2)
}

// UpdateApplication mocks base method.
func (m *MockApplicationRepository) UpdateApplication(arg0 context.Context, arg1 *datastore.Application) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateApplication", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateApplication indicates an expected call of UpdateApplication.
func (mr *MockApplicationRepositoryMockRecorder) UpdateApplication(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApplication", reflect.TypeOf((*MockApplicationRepository)(nil).UpdateApplication), arg0, arg1)
}

// UpdateApplicationEndpointsStatus mocks base method.
func (m *MockApplicationRepository) UpdateApplicationEndpointsStatus(arg0 context.Context, arg1 string, arg2 []string, arg3 datastore.EndpointStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateApplicationEndpointsStatus", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateApplicationEndpointsStatus indicates an expected call of UpdateApplicationEndpointsStatus.
func (mr *MockApplicationRepositoryMockRecorder) UpdateApplicationEndpointsStatus(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApplicationEndpointsStatus", reflect.TypeOf((*MockApplicationRepository)(nil).UpdateApplicationEndpointsStatus), arg0, arg1, arg2, arg3)
}
