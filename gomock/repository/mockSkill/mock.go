// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/mysql/skill/repository.go

// Package mockskill is a generated GoMock package.
package mockskill

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/ramasuryananda/dummy-cv/internal/dto/entity"
)

// MockRepositoryProvider is a mock of RepositoryProvider interface.
type MockRepositoryProvider struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryProviderMockRecorder
}

// MockRepositoryProviderMockRecorder is the mock recorder for MockRepositoryProvider.
type MockRepositoryProviderMockRecorder struct {
	mock *MockRepositoryProvider
}

// NewMockRepositoryProvider creates a new mock instance.
func NewMockRepositoryProvider(ctrl *gomock.Controller) *MockRepositoryProvider {
	mock := &MockRepositoryProvider{ctrl: ctrl}
	mock.recorder = &MockRepositoryProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryProvider) EXPECT() *MockRepositoryProviderMockRecorder {
	return m.recorder
}

// CreateSkillData mocks base method.
func (m *MockRepositoryProvider) CreateSkillData(ctx context.Context, data entity.Skill) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSkillData", ctx, data)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSkillData indicates an expected call of CreateSkillData.
func (mr *MockRepositoryProviderMockRecorder) CreateSkillData(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSkillData", reflect.TypeOf((*MockRepositoryProvider)(nil).CreateSkillData), ctx, data)
}

// DeleteSkillData mocks base method.
func (m *MockRepositoryProvider) DeleteSkillData(ctx context.Context, profileCode, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSkillData", ctx, profileCode, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSkillData indicates an expected call of DeleteSkillData.
func (mr *MockRepositoryProviderMockRecorder) DeleteSkillData(ctx, profileCode, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSkillData", reflect.TypeOf((*MockRepositoryProvider)(nil).DeleteSkillData), ctx, profileCode, id)
}

// GetFirstSkillByProfileCodeandID mocks base method.
func (m *MockRepositoryProvider) GetFirstSkillByProfileCodeandID(ctx context.Context, profileCode, id uint64) (entity.Skill, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFirstSkillByProfileCodeandID", ctx, profileCode, id)
	ret0, _ := ret[0].(entity.Skill)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFirstSkillByProfileCodeandID indicates an expected call of GetFirstSkillByProfileCodeandID.
func (mr *MockRepositoryProviderMockRecorder) GetFirstSkillByProfileCodeandID(ctx, profileCode, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFirstSkillByProfileCodeandID", reflect.TypeOf((*MockRepositoryProvider)(nil).GetFirstSkillByProfileCodeandID), ctx, profileCode, id)
}

// GetSkillByProfileCode mocks base method.
func (m *MockRepositoryProvider) GetSkillByProfileCode(ctx context.Context, profileCode uint64) ([]entity.Skill, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSkillByProfileCode", ctx, profileCode)
	ret0, _ := ret[0].([]entity.Skill)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSkillByProfileCode indicates an expected call of GetSkillByProfileCode.
func (mr *MockRepositoryProviderMockRecorder) GetSkillByProfileCode(ctx, profileCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSkillByProfileCode", reflect.TypeOf((*MockRepositoryProvider)(nil).GetSkillByProfileCode), ctx, profileCode)
}
