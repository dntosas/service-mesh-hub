// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/api/external/istio/rbac/v1alpha1/rbac_config_client.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/mesh-projects/pkg/api/external/istio/rbac/v1alpha1"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

// MockRbacConfigWatcher is a mock of RbacConfigWatcher interface
type MockRbacConfigWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockRbacConfigWatcherMockRecorder
}

// MockRbacConfigWatcherMockRecorder is the mock recorder for MockRbacConfigWatcher
type MockRbacConfigWatcherMockRecorder struct {
	mock *MockRbacConfigWatcher
}

// NewMockRbacConfigWatcher creates a new mock instance
func NewMockRbacConfigWatcher(ctrl *gomock.Controller) *MockRbacConfigWatcher {
	mock := &MockRbacConfigWatcher{ctrl: ctrl}
	mock.recorder = &MockRbacConfigWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRbacConfigWatcher) EXPECT() *MockRbacConfigWatcherMockRecorder {
	return m.recorder
}

// Watch mocks base method
func (m *MockRbacConfigWatcher) Watch(namespace string, opts clients.WatchOpts) (<-chan v1alpha1.RbacConfigList, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", namespace, opts)
	ret0, _ := ret[0].(<-chan v1alpha1.RbacConfigList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockRbacConfigWatcherMockRecorder) Watch(namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockRbacConfigWatcher)(nil).Watch), namespace, opts)
}

// MockRbacConfigClient is a mock of RbacConfigClient interface
type MockRbacConfigClient struct {
	ctrl     *gomock.Controller
	recorder *MockRbacConfigClientMockRecorder
}

// MockRbacConfigClientMockRecorder is the mock recorder for MockRbacConfigClient
type MockRbacConfigClientMockRecorder struct {
	mock *MockRbacConfigClient
}

// NewMockRbacConfigClient creates a new mock instance
func NewMockRbacConfigClient(ctrl *gomock.Controller) *MockRbacConfigClient {
	mock := &MockRbacConfigClient{ctrl: ctrl}
	mock.recorder = &MockRbacConfigClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRbacConfigClient) EXPECT() *MockRbacConfigClientMockRecorder {
	return m.recorder
}

// BaseClient mocks base method
func (m *MockRbacConfigClient) BaseClient() clients.ResourceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseClient")
	ret0, _ := ret[0].(clients.ResourceClient)
	return ret0
}

// BaseClient indicates an expected call of BaseClient
func (mr *MockRbacConfigClientMockRecorder) BaseClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseClient", reflect.TypeOf((*MockRbacConfigClient)(nil).BaseClient))
}

// Register mocks base method
func (m *MockRbacConfigClient) Register() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register")
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockRbacConfigClientMockRecorder) Register() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRbacConfigClient)(nil).Register))
}

// Read mocks base method
func (m *MockRbacConfigClient) Read(namespace, name string, opts clients.ReadOpts) (*v1alpha1.RbacConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", namespace, name, opts)
	ret0, _ := ret[0].(*v1alpha1.RbacConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockRbacConfigClientMockRecorder) Read(namespace, name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockRbacConfigClient)(nil).Read), namespace, name, opts)
}

// Write mocks base method
func (m *MockRbacConfigClient) Write(resource *v1alpha1.RbacConfig, opts clients.WriteOpts) (*v1alpha1.RbacConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", resource, opts)
	ret0, _ := ret[0].(*v1alpha1.RbacConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockRbacConfigClientMockRecorder) Write(resource, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockRbacConfigClient)(nil).Write), resource, opts)
}

// Delete mocks base method
func (m *MockRbacConfigClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", namespace, name, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRbacConfigClientMockRecorder) Delete(namespace, name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRbacConfigClient)(nil).Delete), namespace, name, opts)
}

// List mocks base method
func (m *MockRbacConfigClient) List(namespace string, opts clients.ListOpts) (v1alpha1.RbacConfigList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", namespace, opts)
	ret0, _ := ret[0].(v1alpha1.RbacConfigList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockRbacConfigClientMockRecorder) List(namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRbacConfigClient)(nil).List), namespace, opts)
}

// Watch mocks base method
func (m *MockRbacConfigClient) Watch(namespace string, opts clients.WatchOpts) (<-chan v1alpha1.RbacConfigList, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", namespace, opts)
	ret0, _ := ret[0].(<-chan v1alpha1.RbacConfigList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockRbacConfigClientMockRecorder) Watch(namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockRbacConfigClient)(nil).Watch), namespace, opts)
}
