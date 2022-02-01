// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/caasapplicationprovisioner (interfaces: Runner)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	worker "github.com/juju/worker/v3"
)

// MockRunner is a mock of Runner interface.
type MockRunner struct {
	ctrl     *gomock.Controller
	recorder *MockRunnerMockRecorder
}

// MockRunnerMockRecorder is the mock recorder for MockRunner.
type MockRunnerMockRecorder struct {
	mock *MockRunner
}

// NewMockRunner creates a new mock instance.
func NewMockRunner(ctrl *gomock.Controller) *MockRunner {
	mock := &MockRunner{ctrl: ctrl}
	mock.recorder = &MockRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRunner) EXPECT() *MockRunnerMockRecorder {
	return m.recorder
}

// Kill mocks base method.
func (m *MockRunner) Kill() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Kill")
}

// Kill indicates an expected call of Kill.
func (mr *MockRunnerMockRecorder) Kill() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kill", reflect.TypeOf((*MockRunner)(nil).Kill))
}

// StartWorker mocks base method.
func (m *MockRunner) StartWorker(arg0 string, arg1 func() (worker.Worker, error)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartWorker", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartWorker indicates an expected call of StartWorker.
func (mr *MockRunnerMockRecorder) StartWorker(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWorker", reflect.TypeOf((*MockRunner)(nil).StartWorker), arg0, arg1)
}

// StopAndRemoveWorker mocks base method.
func (m *MockRunner) StopAndRemoveWorker(arg0 string, arg1 <-chan struct{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopAndRemoveWorker", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopAndRemoveWorker indicates an expected call of StopAndRemoveWorker.
func (mr *MockRunnerMockRecorder) StopAndRemoveWorker(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopAndRemoveWorker", reflect.TypeOf((*MockRunner)(nil).StopAndRemoveWorker), arg0, arg1)
}

// Wait mocks base method.
func (m *MockRunner) Wait() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait")
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockRunnerMockRecorder) Wait() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockRunner)(nil).Wait))
}

// Worker mocks base method.
func (m *MockRunner) Worker(arg0 string, arg1 <-chan struct{}) (worker.Worker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Worker", arg0, arg1)
	ret0, _ := ret[0].(worker.Worker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Worker indicates an expected call of Worker.
func (mr *MockRunnerMockRecorder) Worker(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Worker", reflect.TypeOf((*MockRunner)(nil).Worker), arg0, arg1)
}
