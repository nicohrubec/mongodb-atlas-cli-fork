// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store (interfaces: PrivateEndpointLister,PrivateEndpointDescriber,PrivateEndpointCreator,PrivateEndpointDeleter,InterfaceEndpointDescriber,InterfaceEndpointCreator,InterfaceEndpointDeleter,RegionalizedPrivateEndpointSettingUpdater,RegionalizedPrivateEndpointSettingDescriber,DataLakePrivateEndpointLister,DataLakePrivateEndpointCreator,DataLakePrivateEndpointDeleter,DataLakePrivateEndpointDescriber)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20230201001/admin"
)

// MockPrivateEndpointLister is a mock of PrivateEndpointLister interface.
type MockPrivateEndpointLister struct {
	ctrl     *gomock.Controller
	recorder *MockPrivateEndpointListerMockRecorder
}

// MockPrivateEndpointListerMockRecorder is the mock recorder for MockPrivateEndpointLister.
type MockPrivateEndpointListerMockRecorder struct {
	mock *MockPrivateEndpointLister
}

// NewMockPrivateEndpointLister creates a new mock instance.
func NewMockPrivateEndpointLister(ctrl *gomock.Controller) *MockPrivateEndpointLister {
	mock := &MockPrivateEndpointLister{ctrl: ctrl}
	mock.recorder = &MockPrivateEndpointListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrivateEndpointLister) EXPECT() *MockPrivateEndpointListerMockRecorder {
	return m.recorder
}

// PrivateEndpoints mocks base method.
func (m *MockPrivateEndpointLister) PrivateEndpoints(arg0, arg1 string) ([]admin.EndpointService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrivateEndpoints", arg0, arg1)
	ret0, _ := ret[0].([]admin.EndpointService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrivateEndpoints indicates an expected call of PrivateEndpoints.
func (mr *MockPrivateEndpointListerMockRecorder) PrivateEndpoints(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrivateEndpoints", reflect.TypeOf((*MockPrivateEndpointLister)(nil).PrivateEndpoints), arg0, arg1)
}

// MockPrivateEndpointDescriber is a mock of PrivateEndpointDescriber interface.
type MockPrivateEndpointDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockPrivateEndpointDescriberMockRecorder
}

// MockPrivateEndpointDescriberMockRecorder is the mock recorder for MockPrivateEndpointDescriber.
type MockPrivateEndpointDescriberMockRecorder struct {
	mock *MockPrivateEndpointDescriber
}

// NewMockPrivateEndpointDescriber creates a new mock instance.
func NewMockPrivateEndpointDescriber(ctrl *gomock.Controller) *MockPrivateEndpointDescriber {
	mock := &MockPrivateEndpointDescriber{ctrl: ctrl}
	mock.recorder = &MockPrivateEndpointDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrivateEndpointDescriber) EXPECT() *MockPrivateEndpointDescriberMockRecorder {
	return m.recorder
}

// PrivateEndpoint mocks base method.
func (m *MockPrivateEndpointDescriber) PrivateEndpoint(arg0, arg1, arg2 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrivateEndpoint", arg0, arg1, arg2)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrivateEndpoint indicates an expected call of PrivateEndpoint.
func (mr *MockPrivateEndpointDescriberMockRecorder) PrivateEndpoint(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrivateEndpoint", reflect.TypeOf((*MockPrivateEndpointDescriber)(nil).PrivateEndpoint), arg0, arg1, arg2)
}

// MockPrivateEndpointCreator is a mock of PrivateEndpointCreator interface.
type MockPrivateEndpointCreator struct {
	ctrl     *gomock.Controller
	recorder *MockPrivateEndpointCreatorMockRecorder
}

// MockPrivateEndpointCreatorMockRecorder is the mock recorder for MockPrivateEndpointCreator.
type MockPrivateEndpointCreatorMockRecorder struct {
	mock *MockPrivateEndpointCreator
}

// NewMockPrivateEndpointCreator creates a new mock instance.
func NewMockPrivateEndpointCreator(ctrl *gomock.Controller) *MockPrivateEndpointCreator {
	mock := &MockPrivateEndpointCreator{ctrl: ctrl}
	mock.recorder = &MockPrivateEndpointCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrivateEndpointCreator) EXPECT() *MockPrivateEndpointCreatorMockRecorder {
	return m.recorder
}

// CreatePrivateEndpoint mocks base method.
func (m *MockPrivateEndpointCreator) CreatePrivateEndpoint(arg0 string, arg1 *admin.CloudProviderEndpointServiceRequest) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePrivateEndpoint", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePrivateEndpoint indicates an expected call of CreatePrivateEndpoint.
func (mr *MockPrivateEndpointCreatorMockRecorder) CreatePrivateEndpoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePrivateEndpoint", reflect.TypeOf((*MockPrivateEndpointCreator)(nil).CreatePrivateEndpoint), arg0, arg1)
}

// MockPrivateEndpointDeleter is a mock of PrivateEndpointDeleter interface.
type MockPrivateEndpointDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockPrivateEndpointDeleterMockRecorder
}

// MockPrivateEndpointDeleterMockRecorder is the mock recorder for MockPrivateEndpointDeleter.
type MockPrivateEndpointDeleterMockRecorder struct {
	mock *MockPrivateEndpointDeleter
}

// NewMockPrivateEndpointDeleter creates a new mock instance.
func NewMockPrivateEndpointDeleter(ctrl *gomock.Controller) *MockPrivateEndpointDeleter {
	mock := &MockPrivateEndpointDeleter{ctrl: ctrl}
	mock.recorder = &MockPrivateEndpointDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrivateEndpointDeleter) EXPECT() *MockPrivateEndpointDeleterMockRecorder {
	return m.recorder
}

// DeletePrivateEndpoint mocks base method.
func (m *MockPrivateEndpointDeleter) DeletePrivateEndpoint(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePrivateEndpoint", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePrivateEndpoint indicates an expected call of DeletePrivateEndpoint.
func (mr *MockPrivateEndpointDeleterMockRecorder) DeletePrivateEndpoint(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePrivateEndpoint", reflect.TypeOf((*MockPrivateEndpointDeleter)(nil).DeletePrivateEndpoint), arg0, arg1, arg2)
}

// MockInterfaceEndpointDescriber is a mock of InterfaceEndpointDescriber interface.
type MockInterfaceEndpointDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceEndpointDescriberMockRecorder
}

// MockInterfaceEndpointDescriberMockRecorder is the mock recorder for MockInterfaceEndpointDescriber.
type MockInterfaceEndpointDescriberMockRecorder struct {
	mock *MockInterfaceEndpointDescriber
}

// NewMockInterfaceEndpointDescriber creates a new mock instance.
func NewMockInterfaceEndpointDescriber(ctrl *gomock.Controller) *MockInterfaceEndpointDescriber {
	mock := &MockInterfaceEndpointDescriber{ctrl: ctrl}
	mock.recorder = &MockInterfaceEndpointDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterfaceEndpointDescriber) EXPECT() *MockInterfaceEndpointDescriberMockRecorder {
	return m.recorder
}

// InterfaceEndpoint mocks base method.
func (m *MockInterfaceEndpointDescriber) InterfaceEndpoint(arg0, arg1, arg2, arg3 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InterfaceEndpoint", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InterfaceEndpoint indicates an expected call of InterfaceEndpoint.
func (mr *MockInterfaceEndpointDescriberMockRecorder) InterfaceEndpoint(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InterfaceEndpoint", reflect.TypeOf((*MockInterfaceEndpointDescriber)(nil).InterfaceEndpoint), arg0, arg1, arg2, arg3)
}

// MockInterfaceEndpointCreator is a mock of InterfaceEndpointCreator interface.
type MockInterfaceEndpointCreator struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceEndpointCreatorMockRecorder
}

// MockInterfaceEndpointCreatorMockRecorder is the mock recorder for MockInterfaceEndpointCreator.
type MockInterfaceEndpointCreatorMockRecorder struct {
	mock *MockInterfaceEndpointCreator
}

// NewMockInterfaceEndpointCreator creates a new mock instance.
func NewMockInterfaceEndpointCreator(ctrl *gomock.Controller) *MockInterfaceEndpointCreator {
	mock := &MockInterfaceEndpointCreator{ctrl: ctrl}
	mock.recorder = &MockInterfaceEndpointCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterfaceEndpointCreator) EXPECT() *MockInterfaceEndpointCreatorMockRecorder {
	return m.recorder
}

// CreateInterfaceEndpoint mocks base method.
func (m *MockInterfaceEndpointCreator) CreateInterfaceEndpoint(arg0, arg1, arg2 string, arg3 *admin.CreateEndpointRequest) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInterfaceEndpoint", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInterfaceEndpoint indicates an expected call of CreateInterfaceEndpoint.
func (mr *MockInterfaceEndpointCreatorMockRecorder) CreateInterfaceEndpoint(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInterfaceEndpoint", reflect.TypeOf((*MockInterfaceEndpointCreator)(nil).CreateInterfaceEndpoint), arg0, arg1, arg2, arg3)
}

// MockInterfaceEndpointDeleter is a mock of InterfaceEndpointDeleter interface.
type MockInterfaceEndpointDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceEndpointDeleterMockRecorder
}

// MockInterfaceEndpointDeleterMockRecorder is the mock recorder for MockInterfaceEndpointDeleter.
type MockInterfaceEndpointDeleterMockRecorder struct {
	mock *MockInterfaceEndpointDeleter
}

// NewMockInterfaceEndpointDeleter creates a new mock instance.
func NewMockInterfaceEndpointDeleter(ctrl *gomock.Controller) *MockInterfaceEndpointDeleter {
	mock := &MockInterfaceEndpointDeleter{ctrl: ctrl}
	mock.recorder = &MockInterfaceEndpointDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterfaceEndpointDeleter) EXPECT() *MockInterfaceEndpointDeleterMockRecorder {
	return m.recorder
}

// DeleteInterfaceEndpoint mocks base method.
func (m *MockInterfaceEndpointDeleter) DeleteInterfaceEndpoint(arg0, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInterfaceEndpoint", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInterfaceEndpoint indicates an expected call of DeleteInterfaceEndpoint.
func (mr *MockInterfaceEndpointDeleterMockRecorder) DeleteInterfaceEndpoint(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInterfaceEndpoint", reflect.TypeOf((*MockInterfaceEndpointDeleter)(nil).DeleteInterfaceEndpoint), arg0, arg1, arg2, arg3)
}

// MockRegionalizedPrivateEndpointSettingUpdater is a mock of RegionalizedPrivateEndpointSettingUpdater interface.
type MockRegionalizedPrivateEndpointSettingUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockRegionalizedPrivateEndpointSettingUpdaterMockRecorder
}

// MockRegionalizedPrivateEndpointSettingUpdaterMockRecorder is the mock recorder for MockRegionalizedPrivateEndpointSettingUpdater.
type MockRegionalizedPrivateEndpointSettingUpdaterMockRecorder struct {
	mock *MockRegionalizedPrivateEndpointSettingUpdater
}

// NewMockRegionalizedPrivateEndpointSettingUpdater creates a new mock instance.
func NewMockRegionalizedPrivateEndpointSettingUpdater(ctrl *gomock.Controller) *MockRegionalizedPrivateEndpointSettingUpdater {
	mock := &MockRegionalizedPrivateEndpointSettingUpdater{ctrl: ctrl}
	mock.recorder = &MockRegionalizedPrivateEndpointSettingUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegionalizedPrivateEndpointSettingUpdater) EXPECT() *MockRegionalizedPrivateEndpointSettingUpdaterMockRecorder {
	return m.recorder
}

// UpdateRegionalizedPrivateEndpointSetting mocks base method.
func (m *MockRegionalizedPrivateEndpointSettingUpdater) UpdateRegionalizedPrivateEndpointSetting(arg0 string, arg1 bool) (*admin.ProjectSettingItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRegionalizedPrivateEndpointSetting", arg0, arg1)
	ret0, _ := ret[0].(*admin.ProjectSettingItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRegionalizedPrivateEndpointSetting indicates an expected call of UpdateRegionalizedPrivateEndpointSetting.
func (mr *MockRegionalizedPrivateEndpointSettingUpdaterMockRecorder) UpdateRegionalizedPrivateEndpointSetting(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRegionalizedPrivateEndpointSetting", reflect.TypeOf((*MockRegionalizedPrivateEndpointSettingUpdater)(nil).UpdateRegionalizedPrivateEndpointSetting), arg0, arg1)
}

// MockRegionalizedPrivateEndpointSettingDescriber is a mock of RegionalizedPrivateEndpointSettingDescriber interface.
type MockRegionalizedPrivateEndpointSettingDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockRegionalizedPrivateEndpointSettingDescriberMockRecorder
}

// MockRegionalizedPrivateEndpointSettingDescriberMockRecorder is the mock recorder for MockRegionalizedPrivateEndpointSettingDescriber.
type MockRegionalizedPrivateEndpointSettingDescriberMockRecorder struct {
	mock *MockRegionalizedPrivateEndpointSettingDescriber
}

// NewMockRegionalizedPrivateEndpointSettingDescriber creates a new mock instance.
func NewMockRegionalizedPrivateEndpointSettingDescriber(ctrl *gomock.Controller) *MockRegionalizedPrivateEndpointSettingDescriber {
	mock := &MockRegionalizedPrivateEndpointSettingDescriber{ctrl: ctrl}
	mock.recorder = &MockRegionalizedPrivateEndpointSettingDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegionalizedPrivateEndpointSettingDescriber) EXPECT() *MockRegionalizedPrivateEndpointSettingDescriberMockRecorder {
	return m.recorder
}

// RegionalizedPrivateEndpointSetting mocks base method.
func (m *MockRegionalizedPrivateEndpointSettingDescriber) RegionalizedPrivateEndpointSetting(arg0 string) (*admin.ProjectSettingItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegionalizedPrivateEndpointSetting", arg0)
	ret0, _ := ret[0].(*admin.ProjectSettingItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegionalizedPrivateEndpointSetting indicates an expected call of RegionalizedPrivateEndpointSetting.
func (mr *MockRegionalizedPrivateEndpointSettingDescriberMockRecorder) RegionalizedPrivateEndpointSetting(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegionalizedPrivateEndpointSetting", reflect.TypeOf((*MockRegionalizedPrivateEndpointSettingDescriber)(nil).RegionalizedPrivateEndpointSetting), arg0)
}

// MockDataLakePrivateEndpointLister is a mock of DataLakePrivateEndpointLister interface.
type MockDataLakePrivateEndpointLister struct {
	ctrl     *gomock.Controller
	recorder *MockDataLakePrivateEndpointListerMockRecorder
}

// MockDataLakePrivateEndpointListerMockRecorder is the mock recorder for MockDataLakePrivateEndpointLister.
type MockDataLakePrivateEndpointListerMockRecorder struct {
	mock *MockDataLakePrivateEndpointLister
}

// NewMockDataLakePrivateEndpointLister creates a new mock instance.
func NewMockDataLakePrivateEndpointLister(ctrl *gomock.Controller) *MockDataLakePrivateEndpointLister {
	mock := &MockDataLakePrivateEndpointLister{ctrl: ctrl}
	mock.recorder = &MockDataLakePrivateEndpointListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataLakePrivateEndpointLister) EXPECT() *MockDataLakePrivateEndpointListerMockRecorder {
	return m.recorder
}

// DataLakePrivateEndpoints mocks base method.
func (m *MockDataLakePrivateEndpointLister) DataLakePrivateEndpoints(arg0 *admin.ListDataFederationPrivateEndpointsApiParams) (*admin.PaginatedPrivateNetworkEndpointIdEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataLakePrivateEndpoints", arg0)
	ret0, _ := ret[0].(*admin.PaginatedPrivateNetworkEndpointIdEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataLakePrivateEndpoints indicates an expected call of DataLakePrivateEndpoints.
func (mr *MockDataLakePrivateEndpointListerMockRecorder) DataLakePrivateEndpoints(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataLakePrivateEndpoints", reflect.TypeOf((*MockDataLakePrivateEndpointLister)(nil).DataLakePrivateEndpoints), arg0)
}

// MockDataLakePrivateEndpointCreator is a mock of DataLakePrivateEndpointCreator interface.
type MockDataLakePrivateEndpointCreator struct {
	ctrl     *gomock.Controller
	recorder *MockDataLakePrivateEndpointCreatorMockRecorder
}

// MockDataLakePrivateEndpointCreatorMockRecorder is the mock recorder for MockDataLakePrivateEndpointCreator.
type MockDataLakePrivateEndpointCreatorMockRecorder struct {
	mock *MockDataLakePrivateEndpointCreator
}

// NewMockDataLakePrivateEndpointCreator creates a new mock instance.
func NewMockDataLakePrivateEndpointCreator(ctrl *gomock.Controller) *MockDataLakePrivateEndpointCreator {
	mock := &MockDataLakePrivateEndpointCreator{ctrl: ctrl}
	mock.recorder = &MockDataLakePrivateEndpointCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataLakePrivateEndpointCreator) EXPECT() *MockDataLakePrivateEndpointCreatorMockRecorder {
	return m.recorder
}

// DataLakeCreatePrivateEndpoint mocks base method.
func (m *MockDataLakePrivateEndpointCreator) DataLakeCreatePrivateEndpoint(arg0 string, arg1 *admin.PrivateNetworkEndpointIdEntry) (*admin.PaginatedPrivateNetworkEndpointIdEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataLakeCreatePrivateEndpoint", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedPrivateNetworkEndpointIdEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataLakeCreatePrivateEndpoint indicates an expected call of DataLakeCreatePrivateEndpoint.
func (mr *MockDataLakePrivateEndpointCreatorMockRecorder) DataLakeCreatePrivateEndpoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataLakeCreatePrivateEndpoint", reflect.TypeOf((*MockDataLakePrivateEndpointCreator)(nil).DataLakeCreatePrivateEndpoint), arg0, arg1)
}

// MockDataLakePrivateEndpointDeleter is a mock of DataLakePrivateEndpointDeleter interface.
type MockDataLakePrivateEndpointDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockDataLakePrivateEndpointDeleterMockRecorder
}

// MockDataLakePrivateEndpointDeleterMockRecorder is the mock recorder for MockDataLakePrivateEndpointDeleter.
type MockDataLakePrivateEndpointDeleterMockRecorder struct {
	mock *MockDataLakePrivateEndpointDeleter
}

// NewMockDataLakePrivateEndpointDeleter creates a new mock instance.
func NewMockDataLakePrivateEndpointDeleter(ctrl *gomock.Controller) *MockDataLakePrivateEndpointDeleter {
	mock := &MockDataLakePrivateEndpointDeleter{ctrl: ctrl}
	mock.recorder = &MockDataLakePrivateEndpointDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataLakePrivateEndpointDeleter) EXPECT() *MockDataLakePrivateEndpointDeleterMockRecorder {
	return m.recorder
}

// DataLakeDeletePrivateEndpoint mocks base method.
func (m *MockDataLakePrivateEndpointDeleter) DataLakeDeletePrivateEndpoint(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataLakeDeletePrivateEndpoint", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DataLakeDeletePrivateEndpoint indicates an expected call of DataLakeDeletePrivateEndpoint.
func (mr *MockDataLakePrivateEndpointDeleterMockRecorder) DataLakeDeletePrivateEndpoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataLakeDeletePrivateEndpoint", reflect.TypeOf((*MockDataLakePrivateEndpointDeleter)(nil).DataLakeDeletePrivateEndpoint), arg0, arg1)
}

// MockDataLakePrivateEndpointDescriber is a mock of DataLakePrivateEndpointDescriber interface.
type MockDataLakePrivateEndpointDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockDataLakePrivateEndpointDescriberMockRecorder
}

// MockDataLakePrivateEndpointDescriberMockRecorder is the mock recorder for MockDataLakePrivateEndpointDescriber.
type MockDataLakePrivateEndpointDescriberMockRecorder struct {
	mock *MockDataLakePrivateEndpointDescriber
}

// NewMockDataLakePrivateEndpointDescriber creates a new mock instance.
func NewMockDataLakePrivateEndpointDescriber(ctrl *gomock.Controller) *MockDataLakePrivateEndpointDescriber {
	mock := &MockDataLakePrivateEndpointDescriber{ctrl: ctrl}
	mock.recorder = &MockDataLakePrivateEndpointDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataLakePrivateEndpointDescriber) EXPECT() *MockDataLakePrivateEndpointDescriberMockRecorder {
	return m.recorder
}

// DataLakePrivateEndpoint mocks base method.
func (m *MockDataLakePrivateEndpointDescriber) DataLakePrivateEndpoint(arg0, arg1 string) (*admin.PrivateNetworkEndpointIdEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataLakePrivateEndpoint", arg0, arg1)
	ret0, _ := ret[0].(*admin.PrivateNetworkEndpointIdEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataLakePrivateEndpoint indicates an expected call of DataLakePrivateEndpoint.
func (mr *MockDataLakePrivateEndpointDescriberMockRecorder) DataLakePrivateEndpoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataLakePrivateEndpoint", reflect.TypeOf((*MockDataLakePrivateEndpointDescriber)(nil).DataLakePrivateEndpoint), arg0, arg1)
}
