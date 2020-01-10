// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/api/v1/network_bridge_snapshot_simple_emitter.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/mesh-projects/pkg/api/v1"
)

// MockNetworkBridgeSimpleEmitter is a mock of NetworkBridgeSimpleEmitter interface
type MockNetworkBridgeSimpleEmitter struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkBridgeSimpleEmitterMockRecorder
}

// MockNetworkBridgeSimpleEmitterMockRecorder is the mock recorder for MockNetworkBridgeSimpleEmitter
type MockNetworkBridgeSimpleEmitterMockRecorder struct {
	mock *MockNetworkBridgeSimpleEmitter
}

// NewMockNetworkBridgeSimpleEmitter creates a new mock instance
func NewMockNetworkBridgeSimpleEmitter(ctrl *gomock.Controller) *MockNetworkBridgeSimpleEmitter {
	mock := &MockNetworkBridgeSimpleEmitter{ctrl: ctrl}
	mock.recorder = &MockNetworkBridgeSimpleEmitterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetworkBridgeSimpleEmitter) EXPECT() *MockNetworkBridgeSimpleEmitterMockRecorder {
	return m.recorder
}

// Snapshots mocks base method
func (m *MockNetworkBridgeSimpleEmitter) Snapshots(ctx context.Context) (<-chan *v1.NetworkBridgeSnapshot, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshots", ctx)
	ret0, _ := ret[0].(<-chan *v1.NetworkBridgeSnapshot)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Snapshots indicates an expected call of Snapshots
func (mr *MockNetworkBridgeSimpleEmitterMockRecorder) Snapshots(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshots", reflect.TypeOf((*MockNetworkBridgeSimpleEmitter)(nil).Snapshots), ctx)
}
