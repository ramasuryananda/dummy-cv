// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/mysql/profile/repository.go

// Package mockprofile is a generated GoMock package.
package mockprofile

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

// GetUserByProfileCode mocks base method.
func (m *MockRepositoryProvider) GetUserByProfileCode(ctx context.Context, profileCode uint64) (entity.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByProfileCode", ctx, profileCode)
	ret0, _ := ret[0].(entity.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByProfileCode indicates an expected call of GetUserByProfileCode.
func (mr *MockRepositoryProviderMockRecorder) GetUserByProfileCode(ctx, profileCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByProfileCode", reflect.TypeOf((*MockRepositoryProvider)(nil).GetUserByProfileCode), ctx, profileCode)
}

// InsertProfile mocks base method.
func (m *MockRepositoryProvider) InsertProfile(ctx context.Context, profileData entity.Profile) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertProfile", ctx, profileData)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertProfile indicates an expected call of InsertProfile.
func (mr *MockRepositoryProviderMockRecorder) InsertProfile(ctx, profileData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertProfile", reflect.TypeOf((*MockRepositoryProvider)(nil).InsertProfile), ctx, profileData)
}

// UpdateProfile mocks base method.
func (m *MockRepositoryProvider) UpdateProfile(ctx context.Context, profileData entity.Profile) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", ctx, profileData)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockRepositoryProviderMockRecorder) UpdateProfile(ctx, profileData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockRepositoryProvider)(nil).UpdateProfile), ctx, profileData)
}
