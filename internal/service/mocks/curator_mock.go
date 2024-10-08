// Code generated by MockGen. DO NOT EDIT.
// Source: db_cp_6/internal/repo (interfaces: CuratorRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	entity "db_cp_6/internal/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCuratorRepo is a mock of CuratorRepo interface.
type MockCuratorRepo struct {
	ctrl     *gomock.Controller
	recorder *MockCuratorRepoMockRecorder
}

// MockCuratorRepoMockRecorder is the mock recorder for MockCuratorRepo.
type MockCuratorRepoMockRecorder struct {
	mock *MockCuratorRepo
}

// NewMockCuratorRepo creates a new mock instance.
func NewMockCuratorRepo(ctrl *gomock.Controller) *MockCuratorRepo {
	mock := &MockCuratorRepo{ctrl: ctrl}
	mock.recorder = &MockCuratorRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCuratorRepo) EXPECT() *MockCuratorRepoMockRecorder {
	return m.recorder
}

// CreateCurator mocks base method.
func (m *MockCuratorRepo) CreateCurator(arg0 context.Context, arg1 interface{}, arg2 *entity.Curator) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCurator", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCurator indicates an expected call of CreateCurator.
func (mr *MockCuratorRepoMockRecorder) CreateCurator(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCurator", reflect.TypeOf((*MockCuratorRepo)(nil).CreateCurator), arg0, arg1, arg2)
}

// CreateCuratorExpedition mocks base method.
func (m *MockCuratorRepo) CreateCuratorExpedition(arg0 context.Context, arg1 interface{}, arg2, arg3 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCuratorExpedition", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCuratorExpedition indicates an expected call of CreateCuratorExpedition.
func (mr *MockCuratorRepoMockRecorder) CreateCuratorExpedition(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCuratorExpedition", reflect.TypeOf((*MockCuratorRepo)(nil).CreateCuratorExpedition), arg0, arg1, arg2, arg3)
}

// DeleteCurator mocks base method.
func (m *MockCuratorRepo) DeleteCurator(arg0 context.Context, arg1 interface{}, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCurator", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCurator indicates an expected call of DeleteCurator.
func (mr *MockCuratorRepoMockRecorder) DeleteCurator(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCurator", reflect.TypeOf((*MockCuratorRepo)(nil).DeleteCurator), arg0, arg1, arg2)
}

// GetAllCurators mocks base method.
func (m *MockCuratorRepo) GetAllCurators(arg0 context.Context, arg1 interface{}) (entity.Curators, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCurators", arg0, arg1)
	ret0, _ := ret[0].(entity.Curators)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCurators indicates an expected call of GetAllCurators.
func (mr *MockCuratorRepoMockRecorder) GetAllCurators(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCurators", reflect.TypeOf((*MockCuratorRepo)(nil).GetAllCurators), arg0, arg1)
}

// GetExpeditionCurators mocks base method.
func (m *MockCuratorRepo) GetExpeditionCurators(arg0 context.Context, arg1 interface{}, arg2 int) (entity.Curators, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExpeditionCurators", arg0, arg1, arg2)
	ret0, _ := ret[0].(entity.Curators)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExpeditionCurators indicates an expected call of GetExpeditionCurators.
func (mr *MockCuratorRepoMockRecorder) GetExpeditionCurators(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExpeditionCurators", reflect.TypeOf((*MockCuratorRepo)(nil).GetExpeditionCurators), arg0, arg1, arg2)
}
