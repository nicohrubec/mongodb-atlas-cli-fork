// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store (interfaces: CloudProviderRegionsLister)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20230201001/admin"
)

// MockCloudProviderRegionsLister is a mock of CloudProviderRegionsLister interface.
type MockCloudProviderRegionsLister struct {
	ctrl     *gomock.Controller
	recorder *MockCloudProviderRegionsListerMockRecorder
}

// MockCloudProviderRegionsListerMockRecorder is the mock recorder for MockCloudProviderRegionsLister.
type MockCloudProviderRegionsListerMockRecorder struct {
	mock *MockCloudProviderRegionsLister
}

// NewMockCloudProviderRegionsLister creates a new mock instance.
func NewMockCloudProviderRegionsLister(ctrl *gomock.Controller) *MockCloudProviderRegionsLister {
	mock := &MockCloudProviderRegionsLister{ctrl: ctrl}
	mock.recorder = &MockCloudProviderRegionsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudProviderRegionsLister) EXPECT() *MockCloudProviderRegionsListerMockRecorder {
	return m.recorder
}

// CloudProviderRegions mocks base method.
func (m *MockCloudProviderRegionsLister) CloudProviderRegions(arg0, arg1 string, arg2 []string) (*admin.PaginatedApiAtlasProviderRegions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudProviderRegions", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.PaginatedApiAtlasProviderRegions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudProviderRegions indicates an expected call of CloudProviderRegions.
func (mr *MockCloudProviderRegionsListerMockRecorder) CloudProviderRegions(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudProviderRegions", reflect.TypeOf((*MockCloudProviderRegionsLister)(nil).CloudProviderRegions), arg0, arg1, arg2)
}
