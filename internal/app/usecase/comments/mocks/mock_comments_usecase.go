// Code generated by MockGen. DO NOT EDIT.
// Source: patreon/internal/app/usecase/comments (interfaces: Usecase)

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	models "patreon/internal/app/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// CommentsUsecase is a mock of Usecase interface.
type CommentsUsecase struct {
	ctrl     *gomock.Controller
	recorder *CommentsUsecaseMockRecorder
}

// CommentsUsecaseMockRecorder is the mock recorder for CommentsUsecase.
type CommentsUsecaseMockRecorder struct {
	mock *CommentsUsecase
}

// NewCommentsUsecase creates a new mock instance.
func NewCommentsUsecase(ctrl *gomock.Controller) *CommentsUsecase {
	mock := &CommentsUsecase{ctrl: ctrl}
	mock.recorder = &CommentsUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *CommentsUsecase) EXPECT() *CommentsUsecaseMockRecorder {
	return m.recorder
}

// CheckExists mocks base method.
func (m *CommentsUsecase) CheckExists(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExists", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckExists indicates an expected call of CheckExists.
func (mr *CommentsUsecaseMockRecorder) CheckExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExists", reflect.TypeOf((*CommentsUsecase)(nil).CheckExists), arg0)
}

// Create mocks base method.
func (m *CommentsUsecase) Create(arg0 *models.Comment) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *CommentsUsecaseMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*CommentsUsecase)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *CommentsUsecase) Delete(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *CommentsUsecaseMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*CommentsUsecase)(nil).Delete), arg0)
}

// Get mocks base method.
func (m *CommentsUsecase) Get(arg0 int64) (*models.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*models.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *CommentsUsecaseMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*CommentsUsecase)(nil).Get), arg0)
}

// GetPostComments mocks base method.
func (m *CommentsUsecase) GetPostComments(arg0 int64, arg1 *models.Pagination) ([]models.PostComment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostComments", arg0, arg1)
	ret0, _ := ret[0].([]models.PostComment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostComments indicates an expected call of GetPostComments.
func (mr *CommentsUsecaseMockRecorder) GetPostComments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostComments", reflect.TypeOf((*CommentsUsecase)(nil).GetPostComments), arg0, arg1)
}

// GetUserComments mocks base method.
func (m *CommentsUsecase) GetUserComments(arg0 int64, arg1 *models.Pagination) ([]models.UserComment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserComments", arg0, arg1)
	ret0, _ := ret[0].([]models.UserComment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserComments indicates an expected call of GetUserComments.
func (mr *CommentsUsecaseMockRecorder) GetUserComments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserComments", reflect.TypeOf((*CommentsUsecase)(nil).GetUserComments), arg0, arg1)
}

// Update mocks base method.
func (m *CommentsUsecase) Update(arg0 *models.Comment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *CommentsUsecaseMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*CommentsUsecase)(nil).Update), arg0)
}
