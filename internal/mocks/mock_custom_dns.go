// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store (interfaces: CustomDNSEnabler,CustomDNSDisabler,CustomDNSDescriber)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20230201001/admin"
)

// MockCustomDNSEnabler is a mock of CustomDNSEnabler interface.
type MockCustomDNSEnabler struct {
	ctrl     *gomock.Controller
	recorder *MockCustomDNSEnablerMockRecorder
}

// MockCustomDNSEnablerMockRecorder is the mock recorder for MockCustomDNSEnabler.
type MockCustomDNSEnablerMockRecorder struct {
	mock *MockCustomDNSEnabler
}

// NewMockCustomDNSEnabler creates a new mock instance.
func NewMockCustomDNSEnabler(ctrl *gomock.Controller) *MockCustomDNSEnabler {
	mock := &MockCustomDNSEnabler{ctrl: ctrl}
	mock.recorder = &MockCustomDNSEnablerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomDNSEnabler) EXPECT() *MockCustomDNSEnablerMockRecorder {
	return m.recorder
}

// EnableCustomDNS mocks base method.
func (m *MockCustomDNSEnabler) EnableCustomDNS(arg0 string) (*admin.AWSCustomDNSEnabled, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableCustomDNS", arg0)
	ret0, _ := ret[0].(*admin.AWSCustomDNSEnabled)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableCustomDNS indicates an expected call of EnableCustomDNS.
func (mr *MockCustomDNSEnablerMockRecorder) EnableCustomDNS(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableCustomDNS", reflect.TypeOf((*MockCustomDNSEnabler)(nil).EnableCustomDNS), arg0)
}

// MockCustomDNSDisabler is a mock of CustomDNSDisabler interface.
type MockCustomDNSDisabler struct {
	ctrl     *gomock.Controller
	recorder *MockCustomDNSDisablerMockRecorder
}

// MockCustomDNSDisablerMockRecorder is the mock recorder for MockCustomDNSDisabler.
type MockCustomDNSDisablerMockRecorder struct {
	mock *MockCustomDNSDisabler
}

// NewMockCustomDNSDisabler creates a new mock instance.
func NewMockCustomDNSDisabler(ctrl *gomock.Controller) *MockCustomDNSDisabler {
	mock := &MockCustomDNSDisabler{ctrl: ctrl}
	mock.recorder = &MockCustomDNSDisablerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomDNSDisabler) EXPECT() *MockCustomDNSDisablerMockRecorder {
	return m.recorder
}

// DisableCustomDNS mocks base method.
func (m *MockCustomDNSDisabler) DisableCustomDNS(arg0 string) (*admin.AWSCustomDNSEnabled, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableCustomDNS", arg0)
	ret0, _ := ret[0].(*admin.AWSCustomDNSEnabled)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableCustomDNS indicates an expected call of DisableCustomDNS.
func (mr *MockCustomDNSDisablerMockRecorder) DisableCustomDNS(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableCustomDNS", reflect.TypeOf((*MockCustomDNSDisabler)(nil).DisableCustomDNS), arg0)
}

// MockCustomDNSDescriber is a mock of CustomDNSDescriber interface.
type MockCustomDNSDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockCustomDNSDescriberMockRecorder
}

// MockCustomDNSDescriberMockRecorder is the mock recorder for MockCustomDNSDescriber.
type MockCustomDNSDescriberMockRecorder struct {
	mock *MockCustomDNSDescriber
}

// NewMockCustomDNSDescriber creates a new mock instance.
func NewMockCustomDNSDescriber(ctrl *gomock.Controller) *MockCustomDNSDescriber {
	mock := &MockCustomDNSDescriber{ctrl: ctrl}
	mock.recorder = &MockCustomDNSDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomDNSDescriber) EXPECT() *MockCustomDNSDescriberMockRecorder {
	return m.recorder
}

// DescribeCustomDNS mocks base method.
func (m *MockCustomDNSDescriber) DescribeCustomDNS(arg0 string) (*admin.AWSCustomDNSEnabled, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCustomDNS", arg0)
	ret0, _ := ret[0].(*admin.AWSCustomDNSEnabled)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCustomDNS indicates an expected call of DescribeCustomDNS.
func (mr *MockCustomDNSDescriberMockRecorder) DescribeCustomDNS(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCustomDNS", reflect.TypeOf((*MockCustomDNSDescriber)(nil).DescribeCustomDNS), arg0)
}
