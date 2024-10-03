// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/pravin/Desktop/CODESAGE/internal/domain/interfaces/user_service_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	roles "cli-project/internal/config/roles"
	dto "cli-project/internal/domain/dto"
	models "cli-project/internal/domain/models"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// CountActiveUserInLast24Hours mocks base method.
func (m *MockUserService) CountActiveUserInLast24Hours(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountActiveUserInLast24Hours", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountActiveUserInLast24Hours indicates an expected call of CountActiveUserInLast24Hours.
func (mr *MockUserServiceMockRecorder) CountActiveUserInLast24Hours(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountActiveUserInLast24Hours", reflect.TypeOf((*MockUserService)(nil).CountActiveUserInLast24Hours), ctx)
}

// GetAllUsers mocks base method.
func (m *MockUserService) GetAllUsers(ctx context.Context) ([]dto.StandardUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", ctx)
	ret0, _ := ret[0].([]dto.StandardUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockUserServiceMockRecorder) GetAllUsers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUserService)(nil).GetAllUsers), ctx)
}

// GetPlatformStats mocks base method.
func (m *MockUserService) GetPlatformStats(ctx context.Context) (*models.PlatformStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlatformStats", ctx)
	ret0, _ := ret[0].(*models.PlatformStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlatformStats indicates an expected call of GetPlatformStats.
func (mr *MockUserServiceMockRecorder) GetPlatformStats(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlatformStats", reflect.TypeOf((*MockUserService)(nil).GetPlatformStats), ctx)
}

// GetUserByID mocks base method.
func (m *MockUserService) GetUserByID(ctx context.Context, userID string) (*models.StandardUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, userID)
	ret0, _ := ret[0].(*models.StandardUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserServiceMockRecorder) GetUserByID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserService)(nil).GetUserByID), ctx, userID)
}

// GetUserByUsername mocks base method.
func (m *MockUserService) GetUserByUsername(ctx context.Context, username string) (*models.StandardUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", ctx, username)
	ret0, _ := ret[0].(*models.StandardUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockUserServiceMockRecorder) GetUserByUsername(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockUserService)(nil).GetUserByUsername), ctx, username)
}

// GetUserCodesageStats mocks base method.
func (m *MockUserService) GetUserCodesageStats(ctx context.Context, userID string) (*models.CodesageStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserCodesageStats", ctx, userID)
	ret0, _ := ret[0].(*models.CodesageStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserCodesageStats indicates an expected call of GetUserCodesageStats.
func (mr *MockUserServiceMockRecorder) GetUserCodesageStats(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserCodesageStats", reflect.TypeOf((*MockUserService)(nil).GetUserCodesageStats), ctx, userID)
}

// GetUserID mocks base method.
func (m *MockUserService) GetUserID(ctx context.Context, username string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserID", ctx, username)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserID indicates an expected call of GetUserID.
func (mr *MockUserServiceMockRecorder) GetUserID(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserID", reflect.TypeOf((*MockUserService)(nil).GetUserID), ctx, username)
}

// GetUserLeetcodeStats mocks base method.
func (m *MockUserService) GetUserLeetcodeStats(userID string) (*models.LeetcodeStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserLeetcodeStats", userID)
	ret0, _ := ret[0].(*models.LeetcodeStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserLeetcodeStats indicates an expected call of GetUserLeetcodeStats.
func (mr *MockUserServiceMockRecorder) GetUserLeetcodeStats(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserLeetcodeStats", reflect.TypeOf((*MockUserService)(nil).GetUserLeetcodeStats), userID)
}

// GetUserProgress mocks base method.
func (m *MockUserService) GetUserProgress(ctx context.Context, userID string) (*[]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserProgress", ctx, userID)
	ret0, _ := ret[0].(*[]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProgress indicates an expected call of GetUserProgress.
func (mr *MockUserServiceMockRecorder) GetUserProgress(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProgress", reflect.TypeOf((*MockUserService)(nil).GetUserProgress), ctx, userID)
}

// GetUserRole mocks base method.
func (m *MockUserService) GetUserRole(ctx context.Context, userID string) (roles.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRole", ctx, userID)
	ret0, _ := ret[0].(roles.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRole indicates an expected call of GetUserRole.
func (mr *MockUserServiceMockRecorder) GetUserRole(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRole", reflect.TypeOf((*MockUserService)(nil).GetUserRole), ctx, userID)
}

// IsUserBanned mocks base method.
func (m *MockUserService) IsUserBanned(ctx context.Context, userID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserBanned", ctx, userID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUserBanned indicates an expected call of IsUserBanned.
func (mr *MockUserServiceMockRecorder) IsUserBanned(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserBanned", reflect.TypeOf((*MockUserService)(nil).IsUserBanned), ctx, userID)
}

// UpdateUserBanState mocks base method.
func (m *MockUserService) UpdateUserBanState(ctx context.Context, username string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserBanState", ctx, username)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserBanState indicates an expected call of UpdateUserBanState.
func (mr *MockUserServiceMockRecorder) UpdateUserBanState(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserBanState", reflect.TypeOf((*MockUserService)(nil).UpdateUserBanState), ctx, username)
}

// UpdateUserProgress mocks base method.
func (m *MockUserService) UpdateUserProgress(ctx context.Context, userID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserProgress", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserProgress indicates an expected call of UpdateUserProgress.
func (mr *MockUserServiceMockRecorder) UpdateUserProgress(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserProgress", reflect.TypeOf((*MockUserService)(nil).UpdateUserProgress), ctx, userID)
}

// ViewDashboard mocks base method.
func (m *MockUserService) ViewDashboard(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewDashboard", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ViewDashboard indicates an expected call of ViewDashboard.
func (mr *MockUserServiceMockRecorder) ViewDashboard(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewDashboard", reflect.TypeOf((*MockUserService)(nil).ViewDashboard), ctx)
}
