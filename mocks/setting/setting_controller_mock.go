// Code generated by MockGen. DO NOT EDIT.
// Source: controllers/setting/c_setting.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	setting "go-template/controllers/setting"
	setting0 "go-template/entities/setting"
	services "go-template/services"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSettingService is a mock of SettingService interface.
type MockSettingService struct {
	ctrl     *gomock.Controller
	recorder *MockSettingServiceMockRecorder
}

// MockSettingServiceMockRecorder is the mock recorder for MockSettingService.
type MockSettingServiceMockRecorder struct {
	mock *MockSettingService
}

// NewMockSettingService creates a new mock instance.
func NewMockSettingService(ctrl *gomock.Controller) *MockSettingService {
	mock := &MockSettingService{ctrl: ctrl}
	mock.recorder = &MockSettingServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSettingService) EXPECT() *MockSettingServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSettingService) Create(ctx context.Context, input *setting0.Setting) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, input)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSettingServiceMockRecorder) Create(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSettingService)(nil).Create), ctx, input)
}

// Delete mocks base method.
func (m *MockSettingService) Delete(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSettingServiceMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSettingService)(nil).Delete), ctx, id)
}

// FindAll mocks base method.
func (m *MockSettingService) FindAll(arg0 context.Context, arg1 *setting.SettingParameter) ([]setting0.Setting, services.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0, arg1)
	ret0, _ := ret[0].([]setting0.Setting)
	ret1, _ := ret[1].(services.Pagination)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindAll indicates an expected call of FindAll.
func (mr *MockSettingServiceMockRecorder) FindAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockSettingService)(nil).FindAll), arg0, arg1)
}

// FindDefaultSetting mocks base method.
func (m *MockSettingService) FindDefaultSetting(ctx context.Context) (setting0.SettingDefault, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindDefaultSetting", ctx)
	ret0, _ := ret[0].(setting0.SettingDefault)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindDefaultSetting indicates an expected call of FindDefaultSetting.
func (mr *MockSettingServiceMockRecorder) FindDefaultSetting(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDefaultSetting", reflect.TypeOf((*MockSettingService)(nil).FindDefaultSetting), ctx)
}

// FindOne mocks base method.
func (m *MockSettingService) FindOne(arg0 context.Context, arg1 *setting.SettingParameter) (setting0.Setting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0, arg1)
	ret0, _ := ret[0].(setting0.Setting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockSettingServiceMockRecorder) FindOne(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockSettingService)(nil).FindOne), arg0, arg1)
}

// SelectAll mocks base method.
func (m *MockSettingService) SelectAll(arg0 context.Context, arg1 *setting.SettingParameter) ([]setting0.Setting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAll", arg0, arg1)
	ret0, _ := ret[0].([]setting0.Setting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectAll indicates an expected call of SelectAll.
func (mr *MockSettingServiceMockRecorder) SelectAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAll", reflect.TypeOf((*MockSettingService)(nil).SelectAll), arg0, arg1)
}

// Update mocks base method.
func (m *MockSettingService) Update(ctx context.Context, input *setting0.Setting) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockSettingServiceMockRecorder) Update(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSettingService)(nil).Update), ctx, input)
}