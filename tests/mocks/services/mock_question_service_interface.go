// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/pravin/Desktop/CODESAGE/internal/domain/interfaces/question_service_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	dto "cli-project/internal/domain/dto"
	models "cli-project/internal/domain/models"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockQuestionService is a mock of QuestionService interface.
type MockQuestionService struct {
	ctrl     *gomock.Controller
	recorder *MockQuestionServiceMockRecorder
}

// MockQuestionServiceMockRecorder is the mock recorder for MockQuestionService.
type MockQuestionServiceMockRecorder struct {
	mock *MockQuestionService
}

// NewMockQuestionService creates a new mock instance.
func NewMockQuestionService(ctrl *gomock.Controller) *MockQuestionService {
	mock := &MockQuestionService{ctrl: ctrl}
	mock.recorder = &MockQuestionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuestionService) EXPECT() *MockQuestionServiceMockRecorder {
	return m.recorder
}

// AddQuestionsFromFile mocks base method.
func (m *MockQuestionService) AddQuestionsFromFile(ctx context.Context, questionFilePath string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddQuestionsFromFile", ctx, questionFilePath)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddQuestionsFromFile indicates an expected call of AddQuestionsFromFile.
func (mr *MockQuestionServiceMockRecorder) AddQuestionsFromFile(ctx, questionFilePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddQuestionsFromFile", reflect.TypeOf((*MockQuestionService)(nil).AddQuestionsFromFile), ctx, questionFilePath)
}

// GetAllQuestions mocks base method.
func (m *MockQuestionService) GetAllQuestions(arg0 context.Context) ([]dto.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllQuestions", arg0)
	ret0, _ := ret[0].([]dto.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllQuestions indicates an expected call of GetAllQuestions.
func (mr *MockQuestionServiceMockRecorder) GetAllQuestions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllQuestions", reflect.TypeOf((*MockQuestionService)(nil).GetAllQuestions), arg0)
}

// GetQuestionByID mocks base method.
func (m *MockQuestionService) GetQuestionByID(ctx context.Context, questionID string) (*models.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuestionByID", ctx, questionID)
	ret0, _ := ret[0].(*models.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuestionByID indicates an expected call of GetQuestionByID.
func (mr *MockQuestionServiceMockRecorder) GetQuestionByID(ctx, questionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuestionByID", reflect.TypeOf((*MockQuestionService)(nil).GetQuestionByID), ctx, questionID)
}

// GetQuestionsByFilters mocks base method.
func (m *MockQuestionService) GetQuestionsByFilters(ctx context.Context, difficulty, company, topic string) ([]dto.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuestionsByFilters", ctx, difficulty, company, topic)
	ret0, _ := ret[0].([]dto.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuestionsByFilters indicates an expected call of GetQuestionsByFilters.
func (mr *MockQuestionServiceMockRecorder) GetQuestionsByFilters(ctx, difficulty, company, topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuestionsByFilters", reflect.TypeOf((*MockQuestionService)(nil).GetQuestionsByFilters), ctx, difficulty, company, topic)
}

// GetTotalQuestionsCount mocks base method.
func (m *MockQuestionService) GetTotalQuestionsCount(arg0 context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalQuestionsCount", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalQuestionsCount indicates an expected call of GetTotalQuestionsCount.
func (mr *MockQuestionServiceMockRecorder) GetTotalQuestionsCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalQuestionsCount", reflect.TypeOf((*MockQuestionService)(nil).GetTotalQuestionsCount), arg0)
}

// QuestionExistsByID mocks base method.
func (m *MockQuestionService) QuestionExistsByID(ctx context.Context, questionID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuestionExistsByID", ctx, questionID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuestionExistsByID indicates an expected call of QuestionExistsByID.
func (mr *MockQuestionServiceMockRecorder) QuestionExistsByID(ctx, questionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuestionExistsByID", reflect.TypeOf((*MockQuestionService)(nil).QuestionExistsByID), ctx, questionID)
}

// QuestionExistsByTitleSlug mocks base method.
func (m *MockQuestionService) QuestionExistsByTitleSlug(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuestionExistsByTitleSlug", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuestionExistsByTitleSlug indicates an expected call of QuestionExistsByTitleSlug.
func (mr *MockQuestionServiceMockRecorder) QuestionExistsByTitleSlug(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuestionExistsByTitleSlug", reflect.TypeOf((*MockQuestionService)(nil).QuestionExistsByTitleSlug), arg0, arg1)
}

// RemoveQuestionByID mocks base method.
func (m *MockQuestionService) RemoveQuestionByID(ctx context.Context, questionID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveQuestionByID", ctx, questionID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveQuestionByID indicates an expected call of RemoveQuestionByID.
func (mr *MockQuestionServiceMockRecorder) RemoveQuestionByID(ctx, questionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveQuestionByID", reflect.TypeOf((*MockQuestionService)(nil).RemoveQuestionByID), ctx, questionID)
}
