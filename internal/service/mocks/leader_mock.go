// Code generated by MockGen. DO NOT EDIT.
// Source: db_cp_6/internal/repo (interfaces: LeaderRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	entity "db_cp_6/internal/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLeaderRepo is a mock of LeaderRepo interface.
type MockLeaderRepo struct {
	ctrl     *gomock.Controller
	recorder *MockLeaderRepoMockRecorder
}

// MockLeaderRepoMockRecorder is the mock recorder for MockLeaderRepo.
type MockLeaderRepoMockRecorder struct {
	mock *MockLeaderRepo
}

// NewMockLeaderRepo creates a new mock instance.
func NewMockLeaderRepo(ctrl *gomock.Controller) *MockLeaderRepo {
	mock := &MockLeaderRepo{ctrl: ctrl}
	mock.recorder = &MockLeaderRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLeaderRepo) EXPECT() *MockLeaderRepoMockRecorder {
	return m.recorder
}

// CreateLeader mocks base method.
func (m *MockLeaderRepo) CreateLeader(arg0 context.Context, arg1 interface{}, arg2 *entity.Leader) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLeader", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLeader indicates an expected call of CreateLeader.
func (mr *MockLeaderRepoMockRecorder) CreateLeader(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLeader", reflect.TypeOf((*MockLeaderRepo)(nil).CreateLeader), arg0, arg1, arg2)
}

// CreateLeaderExpedition mocks base method.
func (m *MockLeaderRepo) CreateLeaderExpedition(arg0 context.Context, arg1 interface{}, arg2, arg3 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLeaderExpedition", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLeaderExpedition indicates an expected call of CreateLeaderExpedition.
func (mr *MockLeaderRepoMockRecorder) CreateLeaderExpedition(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLeaderExpedition", reflect.TypeOf((*MockLeaderRepo)(nil).CreateLeaderExpedition), arg0, arg1, arg2, arg3)
}

// DeleteLeader mocks base method.
func (m *MockLeaderRepo) DeleteLeader(arg0 context.Context, arg1 interface{}, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLeader", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLeader indicates an expected call of DeleteLeader.
func (mr *MockLeaderRepoMockRecorder) DeleteLeader(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLeader", reflect.TypeOf((*MockLeaderRepo)(nil).DeleteLeader), arg0, arg1, arg2)
}

// GetAllLeaders mocks base method.
func (m *MockLeaderRepo) GetAllLeaders(arg0 context.Context, arg1 interface{}) (entity.Leaders, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllLeaders", arg0, arg1)
	ret0, _ := ret[0].(entity.Leaders)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllLeaders indicates an expected call of GetAllLeaders.
func (mr *MockLeaderRepoMockRecorder) GetAllLeaders(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllLeaders", reflect.TypeOf((*MockLeaderRepo)(nil).GetAllLeaders), arg0, arg1)
}

// GetExpeditionLeaders mocks base method.
func (m *MockLeaderRepo) GetExpeditionLeaders(arg0 context.Context, arg1 interface{}, arg2 int) (entity.Leaders, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExpeditionLeaders", arg0, arg1, arg2)
	ret0, _ := ret[0].(entity.Leaders)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExpeditionLeaders indicates an expected call of GetExpeditionLeaders.
func (mr *MockLeaderRepoMockRecorder) GetExpeditionLeaders(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExpeditionLeaders", reflect.TypeOf((*MockLeaderRepo)(nil).GetExpeditionLeaders), arg0, arg1, arg2)
}

// GetLeaderByLogin mocks base method.
func (m *MockLeaderRepo) GetLeaderByLogin(arg0 context.Context, arg1 interface{}, arg2 string) (*entity.Leader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLeaderByLogin", arg0, arg1, arg2)
	ret0, _ := ret[0].(*entity.Leader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLeaderByLogin indicates an expected call of GetLeaderByLogin.
func (mr *MockLeaderRepoMockRecorder) GetLeaderByLogin(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLeaderByLogin", reflect.TypeOf((*MockLeaderRepo)(nil).GetLeaderByLogin), arg0, arg1, arg2)
}
