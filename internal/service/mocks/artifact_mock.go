// Code generated by MockGen. DO NOT EDIT.
// Source: db_cp_6/internal/repo (interfaces: ArtifactRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	entity "db_cp_6/internal/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockArtifactRepo is a mock of ArtifactRepo interface.
type MockArtifactRepo struct {
	ctrl     *gomock.Controller
	recorder *MockArtifactRepoMockRecorder
}

// MockArtifactRepoMockRecorder is the mock recorder for MockArtifactRepo.
type MockArtifactRepoMockRecorder struct {
	mock *MockArtifactRepo
}

// NewMockArtifactRepo creates a new mock instance.
func NewMockArtifactRepo(ctrl *gomock.Controller) *MockArtifactRepo {
	mock := &MockArtifactRepo{ctrl: ctrl}
	mock.recorder = &MockArtifactRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArtifactRepo) EXPECT() *MockArtifactRepoMockRecorder {
	return m.recorder
}

// CreateArtifact mocks base method.
func (m *MockArtifactRepo) CreateArtifact(arg0 context.Context, arg1 interface{}, arg2 *entity.Artifact) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateArtifact", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateArtifact indicates an expected call of CreateArtifact.
func (mr *MockArtifactRepoMockRecorder) CreateArtifact(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateArtifact", reflect.TypeOf((*MockArtifactRepo)(nil).CreateArtifact), arg0, arg1, arg2)
}

// GetAllArtifacts mocks base method.
func (m *MockArtifactRepo) GetAllArtifacts(arg0 context.Context, arg1 interface{}) (entity.Artifacts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllArtifacts", arg0, arg1)
	ret0, _ := ret[0].(entity.Artifacts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllArtifacts indicates an expected call of GetAllArtifacts.
func (mr *MockArtifactRepoMockRecorder) GetAllArtifacts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllArtifacts", reflect.TypeOf((*MockArtifactRepo)(nil).GetAllArtifacts), arg0, arg1)
}

// GetLocationArtifacts mocks base method.
func (m *MockArtifactRepo) GetLocationArtifacts(arg0 context.Context, arg1 interface{}, arg2 int) (entity.Artifacts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocationArtifacts", arg0, arg1, arg2)
	ret0, _ := ret[0].(entity.Artifacts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLocationArtifacts indicates an expected call of GetLocationArtifacts.
func (mr *MockArtifactRepoMockRecorder) GetLocationArtifacts(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocationArtifacts", reflect.TypeOf((*MockArtifactRepo)(nil).GetLocationArtifacts), arg0, arg1, arg2)
}
