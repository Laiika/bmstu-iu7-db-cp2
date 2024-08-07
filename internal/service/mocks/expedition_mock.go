// Code generated by MockGen. DO NOT EDIT.
// Source: db_cp_6/internal/repo (interfaces: ExpeditionRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	entity "db_cp_6/internal/entity"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockExpeditionRepo is a mock of ExpeditionRepo interface.
type MockExpeditionRepo struct {
	ctrl     *gomock.Controller
	recorder *MockExpeditionRepoMockRecorder
}

// MockExpeditionRepoMockRecorder is the mock recorder for MockExpeditionRepo.
type MockExpeditionRepoMockRecorder struct {
	mock *MockExpeditionRepo
}

// NewMockExpeditionRepo creates a new mock instance.
func NewMockExpeditionRepo(ctrl *gomock.Controller) *MockExpeditionRepo {
	mock := &MockExpeditionRepo{ctrl: ctrl}
	mock.recorder = &MockExpeditionRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExpeditionRepo) EXPECT() *MockExpeditionRepoMockRecorder {
	return m.recorder
}

// CreateExpedition mocks base method.
func (m *MockExpeditionRepo) CreateExpedition(arg0 context.Context, arg1 interface{}, arg2 *entity.Expedition) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExpedition", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateExpedition indicates an expected call of CreateExpedition.
func (mr *MockExpeditionRepoMockRecorder) CreateExpedition(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExpedition", reflect.TypeOf((*MockExpeditionRepo)(nil).CreateExpedition), arg0, arg1, arg2)
}

// DeleteExpedition mocks base method.
func (m *MockExpeditionRepo) DeleteExpedition(arg0 context.Context, arg1 interface{}, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExpedition", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExpedition indicates an expected call of DeleteExpedition.
func (mr *MockExpeditionRepoMockRecorder) DeleteExpedition(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExpedition", reflect.TypeOf((*MockExpeditionRepo)(nil).DeleteExpedition), arg0, arg1, arg2)
}

// GetAllExpeditions mocks base method.
func (m *MockExpeditionRepo) GetAllExpeditions(arg0 context.Context, arg1 interface{}) (entity.Expeditions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllExpeditions", arg0, arg1)
	ret0, _ := ret[0].(entity.Expeditions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllExpeditions indicates an expected call of GetAllExpeditions.
func (mr *MockExpeditionRepoMockRecorder) GetAllExpeditions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllExpeditions", reflect.TypeOf((*MockExpeditionRepo)(nil).GetAllExpeditions), arg0, arg1)
}

// GetCuratorExpeditions mocks base method.
func (m *MockExpeditionRepo) GetCuratorExpeditions(arg0 context.Context, arg1 interface{}, arg2 int) (entity.Expeditions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCuratorExpeditions", arg0, arg1, arg2)
	ret0, _ := ret[0].(entity.Expeditions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCuratorExpeditions indicates an expected call of GetCuratorExpeditions.
func (mr *MockExpeditionRepoMockRecorder) GetCuratorExpeditions(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCuratorExpeditions", reflect.TypeOf((*MockExpeditionRepo)(nil).GetCuratorExpeditions), arg0, arg1, arg2)
}

// GetExpeditionById mocks base method.
func (m *MockExpeditionRepo) GetExpeditionById(arg0 context.Context, arg1 interface{}, arg2 int) (*entity.Expedition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExpeditionById", arg0, arg1, arg2)
	ret0, _ := ret[0].(*entity.Expedition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExpeditionById indicates an expected call of GetExpeditionById.
func (mr *MockExpeditionRepoMockRecorder) GetExpeditionById(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExpeditionById", reflect.TypeOf((*MockExpeditionRepo)(nil).GetExpeditionById), arg0, arg1, arg2)
}

// GetLeaderExpeditions mocks base method.
func (m *MockExpeditionRepo) GetLeaderExpeditions(arg0 context.Context, arg1 interface{}, arg2 int) (entity.Expeditions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLeaderExpeditions", arg0, arg1, arg2)
	ret0, _ := ret[0].(entity.Expeditions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLeaderExpeditions indicates an expected call of GetLeaderExpeditions.
func (mr *MockExpeditionRepoMockRecorder) GetLeaderExpeditions(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLeaderExpeditions", reflect.TypeOf((*MockExpeditionRepo)(nil).GetLeaderExpeditions), arg0, arg1, arg2)
}

// GetMemberExpeditions mocks base method.
func (m *MockExpeditionRepo) GetMemberExpeditions(arg0 context.Context, arg1 interface{}, arg2 int) (entity.Expeditions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMemberExpeditions", arg0, arg1, arg2)
	ret0, _ := ret[0].(entity.Expeditions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMemberExpeditions indicates an expected call of GetMemberExpeditions.
func (mr *MockExpeditionRepoMockRecorder) GetMemberExpeditions(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMemberExpeditions", reflect.TypeOf((*MockExpeditionRepo)(nil).GetMemberExpeditions), arg0, arg1, arg2)
}

// UpdateExpeditionDates mocks base method.
func (m *MockExpeditionRepo) UpdateExpeditionDates(arg0 context.Context, arg1 interface{}, arg2 int, arg3, arg4 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExpeditionDates", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExpeditionDates indicates an expected call of UpdateExpeditionDates.
func (mr *MockExpeditionRepoMockRecorder) UpdateExpeditionDates(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExpeditionDates", reflect.TypeOf((*MockExpeditionRepo)(nil).UpdateExpeditionDates), arg0, arg1, arg2, arg3, arg4)
}
