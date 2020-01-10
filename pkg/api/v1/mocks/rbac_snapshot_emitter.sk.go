// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/api/v1/rbac_snapshot_emitter.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/mesh-projects/pkg/api/v1"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

// MockRbacSnapshotEmitter is a mock of RbacSnapshotEmitter interface
type MockRbacSnapshotEmitter struct {
	ctrl     *gomock.Controller
	recorder *MockRbacSnapshotEmitterMockRecorder
}

// MockRbacSnapshotEmitterMockRecorder is the mock recorder for MockRbacSnapshotEmitter
type MockRbacSnapshotEmitterMockRecorder struct {
	mock *MockRbacSnapshotEmitter
}

// NewMockRbacSnapshotEmitter creates a new mock instance
func NewMockRbacSnapshotEmitter(ctrl *gomock.Controller) *MockRbacSnapshotEmitter {
	mock := &MockRbacSnapshotEmitter{ctrl: ctrl}
	mock.recorder = &MockRbacSnapshotEmitterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRbacSnapshotEmitter) EXPECT() *MockRbacSnapshotEmitterMockRecorder {
	return m.recorder
}

// Snapshots mocks base method
func (m *MockRbacSnapshotEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *v1.RbacSnapshot, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshots", watchNamespaces, opts)
	ret0, _ := ret[0].(<-chan *v1.RbacSnapshot)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Snapshots indicates an expected call of Snapshots
func (mr *MockRbacSnapshotEmitterMockRecorder) Snapshots(watchNamespaces, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshots", reflect.TypeOf((*MockRbacSnapshotEmitter)(nil).Snapshots), watchNamespaces, opts)
}

// MockRbacEmitter is a mock of RbacEmitter interface
type MockRbacEmitter struct {
	ctrl     *gomock.Controller
	recorder *MockRbacEmitterMockRecorder
}

// MockRbacEmitterMockRecorder is the mock recorder for MockRbacEmitter
type MockRbacEmitterMockRecorder struct {
	mock *MockRbacEmitter
}

// NewMockRbacEmitter creates a new mock instance
func NewMockRbacEmitter(ctrl *gomock.Controller) *MockRbacEmitter {
	mock := &MockRbacEmitter{ctrl: ctrl}
	mock.recorder = &MockRbacEmitterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRbacEmitter) EXPECT() *MockRbacEmitterMockRecorder {
	return m.recorder
}

// Snapshots mocks base method
func (m *MockRbacEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *v1.RbacSnapshot, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshots", watchNamespaces, opts)
	ret0, _ := ret[0].(<-chan *v1.RbacSnapshot)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Snapshots indicates an expected call of Snapshots
func (mr *MockRbacEmitterMockRecorder) Snapshots(watchNamespaces, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshots", reflect.TypeOf((*MockRbacEmitter)(nil).Snapshots), watchNamespaces, opts)
}

// Register mocks base method
func (m *MockRbacEmitter) Register() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register")
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockRbacEmitterMockRecorder) Register() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRbacEmitter)(nil).Register))
}

// Mesh mocks base method
func (m *MockRbacEmitter) Mesh() v1.MeshClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mesh")
	ret0, _ := ret[0].(v1.MeshClient)
	return ret0
}

// Mesh indicates an expected call of Mesh
func (mr *MockRbacEmitterMockRecorder) Mesh() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mesh", reflect.TypeOf((*MockRbacEmitter)(nil).Mesh))
}
