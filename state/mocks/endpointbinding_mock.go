// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/state (interfaces: EndpointBinding)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	network "github.com/juju/juju/core/network"
	state "github.com/juju/juju/state"
)

// MockEndpointBinding is a mock of EndpointBinding interface.
type MockEndpointBinding struct {
	ctrl     *gomock.Controller
	recorder *MockEndpointBindingMockRecorder
}

// MockEndpointBindingMockRecorder is the mock recorder for MockEndpointBinding.
type MockEndpointBindingMockRecorder struct {
	mock *MockEndpointBinding
}

// NewMockEndpointBinding creates a new mock instance.
func NewMockEndpointBinding(ctrl *gomock.Controller) *MockEndpointBinding {
	mock := &MockEndpointBinding{ctrl: ctrl}
	mock.recorder = &MockEndpointBindingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEndpointBinding) EXPECT() *MockEndpointBindingMockRecorder {
	return m.recorder
}

// AllSpaceInfos mocks base method.
func (m *MockEndpointBinding) AllSpaceInfos() (network.SpaceInfos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllSpaceInfos")
	ret0, _ := ret[0].(network.SpaceInfos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllSpaceInfos indicates an expected call of AllSpaceInfos.
func (mr *MockEndpointBindingMockRecorder) AllSpaceInfos() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllSpaceInfos", reflect.TypeOf((*MockEndpointBinding)(nil).AllSpaceInfos))
}

// DefaultEndpointBindingSpace mocks base method.
func (m *MockEndpointBinding) DefaultEndpointBindingSpace() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultEndpointBindingSpace")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DefaultEndpointBindingSpace indicates an expected call of DefaultEndpointBindingSpace.
func (mr *MockEndpointBindingMockRecorder) DefaultEndpointBindingSpace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultEndpointBindingSpace", reflect.TypeOf((*MockEndpointBinding)(nil).DefaultEndpointBindingSpace))
}

// Space mocks base method.
func (m *MockEndpointBinding) Space(arg0 string) (*state.Space, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Space", arg0)
	ret0, _ := ret[0].(*state.Space)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Space indicates an expected call of Space.
func (mr *MockEndpointBindingMockRecorder) Space(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Space", reflect.TypeOf((*MockEndpointBinding)(nil).Space), arg0)
}
