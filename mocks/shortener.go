// Code generated by MockGen. DO NOT EDIT.
// Source: shortener/internal/app (interfaces: Shortener)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockShortener is a mock of Shortener interface.
type MockShortener struct {
	ctrl     *gomock.Controller
	recorder *MockShortenerMockRecorder
}

// MockShortenerMockRecorder is the mock recorder for MockShortener.
type MockShortenerMockRecorder struct {
	mock *MockShortener
}

// NewMockShortener creates a new mock instance.
func NewMockShortener(ctrl *gomock.Controller) *MockShortener {
	mock := &MockShortener{ctrl: ctrl}
	mock.recorder = &MockShortenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShortener) EXPECT() *MockShortenerMockRecorder {
	return m.recorder
}

// CreateUrl mocks base method.
func (m *MockShortener) CreateUrl(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUrl", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUrl indicates an expected call of CreateUrl.
func (mr *MockShortenerMockRecorder) CreateUrl(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUrl", reflect.TypeOf((*MockShortener)(nil).CreateUrl), arg0, arg1)
}

// GetUrl mocks base method.
func (m *MockShortener) GetUrl(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUrl", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUrl indicates an expected call of GetUrl.
func (mr *MockShortenerMockRecorder) GetUrl(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUrl", reflect.TypeOf((*MockShortener)(nil).GetUrl), arg0, arg1)
}
