// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Projects-WG\CLI-Project\internal\domain\interfaces\user_service_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "cli-project/internal/domain/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
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

// BanUser mocks base method.
func (m *MockUserService) BanUser(username string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BanUser", username)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BanUser indicates an expected call of BanUser.
func (mr *MockUserServiceMockRecorder) BanUser(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BanUser", reflect.TypeOf((*MockUserService)(nil).BanUser), username)
}

// CountActiveUserInLast24Hours mocks base method.
func (m *MockUserService) CountActiveUserInLast24Hours() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountActiveUserInLast24Hours")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountActiveUserInLast24Hours indicates an expected call of CountActiveUserInLast24Hours.
func (mr *MockUserServiceMockRecorder) CountActiveUserInLast24Hours() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountActiveUserInLast24Hours", reflect.TypeOf((*MockUserService)(nil).CountActiveUserInLast24Hours))
}

// GetAllUsers mocks base method.
func (m *MockUserService) GetAllUsers() (*[]models.StandardUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers")
	ret0, _ := ret[0].(*[]models.StandardUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockUserServiceMockRecorder) GetAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUserService)(nil).GetAllUsers))
}

// GetLeetcodeStats mocks base method.
func (m *MockUserService) GetLeetcodeStats(userID string) (*models.LeetcodeStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLeetcodeStats", userID)
	ret0, _ := ret[0].(*models.LeetcodeStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLeetcodeStats indicates an expected call of GetLeetcodeStats.
func (mr *MockUserServiceMockRecorder) GetLeetcodeStats(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLeetcodeStats", reflect.TypeOf((*MockUserService)(nil).GetLeetcodeStats), userID)
}

// GetUserByID mocks base method.
func (m *MockUserService) GetUserByID(userID string) (*models.StandardUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", userID)
	ret0, _ := ret[0].(*models.StandardUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserServiceMockRecorder) GetUserByID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserService)(nil).GetUserByID), userID)
}

// GetUserByUsername mocks base method.
func (m *MockUserService) GetUserByUsername(username string) (*models.StandardUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", username)
	ret0, _ := ret[0].(*models.StandardUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockUserServiceMockRecorder) GetUserByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockUserService)(nil).GetUserByUsername), username)
}

// GetUserID mocks base method.
func (m *MockUserService) GetUserID(username string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserID", username)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserID indicates an expected call of GetUserID.
func (mr *MockUserServiceMockRecorder) GetUserID(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserID", reflect.TypeOf((*MockUserService)(nil).GetUserID), username)
}

// GetUserRole mocks base method.
func (m *MockUserService) GetUserRole(userID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRole", userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRole indicates an expected call of GetUserRole.
func (mr *MockUserServiceMockRecorder) GetUserRole(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRole", reflect.TypeOf((*MockUserService)(nil).GetUserRole), userID)
}

// IsUserBanned mocks base method.
func (m *MockUserService) IsUserBanned(userID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserBanned", userID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUserBanned indicates an expected call of IsUserBanned.
func (mr *MockUserServiceMockRecorder) IsUserBanned(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserBanned", reflect.TypeOf((*MockUserService)(nil).IsUserBanned), userID)
}

// Login mocks base method.
func (m *MockUserService) Login(username, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", username, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// Login indicates an expected call of Login.
func (mr *MockUserServiceMockRecorder) Login(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserService)(nil).Login), username, password)
}

// Logout mocks base method.
func (m *MockUserService) Logout() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout")
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout.
func (mr *MockUserServiceMockRecorder) Logout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockUserService)(nil).Logout))
}

// Signup mocks base method.
func (m *MockUserService) Signup(user *models.StandardUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signup", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Signup indicates an expected call of Signup.
func (mr *MockUserServiceMockRecorder) Signup(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signup", reflect.TypeOf((*MockUserService)(nil).Signup), user)
}

// UnbanUser mocks base method.
func (m *MockUserService) UnbanUser(username string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnbanUser", username)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnbanUser indicates an expected call of UnbanUser.
func (mr *MockUserServiceMockRecorder) UnbanUser(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnbanUser", reflect.TypeOf((*MockUserService)(nil).UnbanUser), username)
}

// UpdateUserProgress mocks base method.
func (m *MockUserService) UpdateUserProgress(solvedQuestionID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserProgress", solvedQuestionID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserProgress indicates an expected call of UpdateUserProgress.
func (mr *MockUserServiceMockRecorder) UpdateUserProgress(solvedQuestionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserProgress", reflect.TypeOf((*MockUserService)(nil).UpdateUserProgress), solvedQuestionID)
}

// ViewDashboard mocks base method.
func (m *MockUserService) ViewDashboard() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewDashboard")
	ret0, _ := ret[0].(error)
	return ret0
}

// ViewDashboard indicates an expected call of ViewDashboard.
func (mr *MockUserServiceMockRecorder) ViewDashboard() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewDashboard", reflect.TypeOf((*MockUserService)(nil).ViewDashboard))
}
