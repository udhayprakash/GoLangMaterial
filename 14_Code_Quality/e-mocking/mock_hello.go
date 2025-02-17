// Code generated by MockGen. DO NOT EDIT.
// Source: a_hello_world.go
//
// Generated by this command:
//
//	mockgen -source=a_hello_world.go -destination=mock_hello.go -package=main
//

// Package main is a generated GoMock package.
package main

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockHelloService is a mock of HelloService interface.
type MockHelloService struct {
	ctrl     *gomock.Controller
	recorder *MockHelloServiceMockRecorder
	isgomock struct{}
}

// MockHelloServiceMockRecorder is the mock recorder for MockHelloService.
type MockHelloServiceMockRecorder struct {
	mock *MockHelloService
}

// NewMockHelloService creates a new mock instance.
func NewMockHelloService(ctrl *gomock.Controller) *MockHelloService {
	mock := &MockHelloService{ctrl: ctrl}
	mock.recorder = &MockHelloServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHelloService) EXPECT() *MockHelloServiceMockRecorder {
	return m.recorder
}

// Hello mocks base method.
func (m *MockHelloService) Hello() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Hello")
}

// Hello indicates an expected call of Hello.
func (mr *MockHelloServiceMockRecorder) Hello() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hello", reflect.TypeOf((*MockHelloService)(nil).Hello))
}
