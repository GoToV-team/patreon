// Code generated by MockGen. DO NOT EDIT.
// Source: patreon/internal/app/usecase/subscribers (interfaces: Usecase)

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	models "patreon/internal/app/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// SubscribersUsecase is a mock of Usecase interface.
type SubscribersUsecase struct {
	ctrl     *gomock.Controller
	recorder *SubscribersUsecaseMockRecorder
}

// SubscribersUsecaseMockRecorder is the mock recorder for SubscribersUsecase.
type SubscribersUsecaseMockRecorder struct {
	mock *SubscribersUsecase
}

// NewSubscribersUsecase creates a new mock instance.
func NewSubscribersUsecase(ctrl *gomock.Controller) *SubscribersUsecase {
	mock := &SubscribersUsecase{ctrl: ctrl}
	mock.recorder = &SubscribersUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *SubscribersUsecase) EXPECT() *SubscribersUsecaseMockRecorder {
	return m.recorder
}

// GetCreators mocks base method.
func (m *SubscribersUsecase) GetCreators(arg0 int64) ([]models.Creator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreators", arg0)
	ret0, _ := ret[0].([]models.Creator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCreators indicates an expected call of GetCreators.
func (mr *SubscribersUsecaseMockRecorder) GetCreators(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreators", reflect.TypeOf((*SubscribersUsecase)(nil).GetCreators), arg0)
}

// GetSubscribers mocks base method.
func (m *SubscribersUsecase) GetSubscribers(arg0 int64) ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscribers", arg0)
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscribers indicates an expected call of GetSubscribers.
func (mr *SubscribersUsecaseMockRecorder) GetSubscribers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscribers", reflect.TypeOf((*SubscribersUsecase)(nil).GetSubscribers), arg0)
}

// Subscribe mocks base method.
func (m *SubscribersUsecase) Subscribe(arg0 *models.Subscriber) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *SubscribersUsecaseMockRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*SubscribersUsecase)(nil).Subscribe), arg0)
}

// UnSubscribe mocks base method.
func (m *SubscribersUsecase) UnSubscribe(arg0 *models.Subscriber) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnSubscribe", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnSubscribe indicates an expected call of UnSubscribe.
func (mr *SubscribersUsecaseMockRecorder) UnSubscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnSubscribe", reflect.TypeOf((*SubscribersUsecase)(nil).UnSubscribe), arg0)
}
