// Code generated by MockGen. DO NOT EDIT.
// Source: base_service/internal/domain/interfaces/user (interfaces: UserCommandRepository,UserQueryRepository,CacheRepository)

// Package mock_user is a generated GoMock package.
package mock_user

import (
	entities "base_service/internal/domain/entities"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserCommandRepository is a mock of UserCommandRepository interface.
type MockUserCommandRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserCommandRepositoryMockRecorder
}

// MockUserCommandRepositoryMockRecorder is the mock recorder for MockUserCommandRepository.
type MockUserCommandRepositoryMockRecorder struct {
	mock *MockUserCommandRepository
}

// NewMockUserCommandRepository creates a new mock instance.
func NewMockUserCommandRepository(ctrl *gomock.Controller) *MockUserCommandRepository {
	mock := &MockUserCommandRepository{ctrl: ctrl}
	mock.recorder = &MockUserCommandRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserCommandRepository) EXPECT() *MockUserCommandRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserCommandRepository) CreateUser(arg0 context.Context, arg1 *entities.User) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserCommandRepositoryMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserCommandRepository)(nil).CreateUser), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockUserCommandRepository) UpdateUser(arg0 context.Context, arg1 *entities.User) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserCommandRepositoryMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserCommandRepository)(nil).UpdateUser), arg0, arg1)
}

// MockUserQueryRepository is a mock of UserQueryRepository interface.
type MockUserQueryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserQueryRepositoryMockRecorder
}

// MockUserQueryRepositoryMockRecorder is the mock recorder for MockUserQueryRepository.
type MockUserQueryRepositoryMockRecorder struct {
	mock *MockUserQueryRepository
}

// NewMockUserQueryRepository creates a new mock instance.
func NewMockUserQueryRepository(ctrl *gomock.Controller) *MockUserQueryRepository {
	mock := &MockUserQueryRepository{ctrl: ctrl}
	mock.recorder = &MockUserQueryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserQueryRepository) EXPECT() *MockUserQueryRepositoryMockRecorder {
	return m.recorder
}

// GetUser mocks base method.
func (m *MockUserQueryRepository) GetUser(arg0 context.Context, arg1 string) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserQueryRepositoryMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserQueryRepository)(nil).GetUser), arg0, arg1)
}

// MockCacheRepository is a mock of CacheRepository interface.
type MockCacheRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCacheRepositoryMockRecorder
}

// MockCacheRepositoryMockRecorder is the mock recorder for MockCacheRepository.
type MockCacheRepositoryMockRecorder struct {
	mock *MockCacheRepository
}

// NewMockCacheRepository creates a new mock instance.
func NewMockCacheRepository(ctrl *gomock.Controller) *MockCacheRepository {
	mock := &MockCacheRepository{ctrl: ctrl}
	mock.recorder = &MockCacheRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheRepository) EXPECT() *MockCacheRepositoryMockRecorder {
	return m.recorder
}

// GetUser mocks base method.
func (m *MockCacheRepository) GetUser(arg0 context.Context, arg1 string) *entities.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(*entities.User)
	return ret0
}

// GetUser indicates an expected call of GetUser.
func (mr *MockCacheRepositoryMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockCacheRepository)(nil).GetUser), arg0, arg1)
}

// SetUser mocks base method.
func (m *MockCacheRepository) SetUser(arg0 context.Context, arg1 *entities.User, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUser indicates an expected call of SetUser.
func (mr *MockCacheRepositoryMockRecorder) SetUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUser", reflect.TypeOf((*MockCacheRepository)(nil).SetUser), arg0, arg1, arg2)
}
