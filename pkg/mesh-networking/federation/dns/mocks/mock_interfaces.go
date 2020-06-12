// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_dns is a generated GoMock package.
package mock_dns

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dns "github.com/solo-io/service-mesh-hub/pkg/mesh-networking/federation/dns"
	v1 "k8s.io/api/core/v1"
)

// MockIpAssigner is a mock of IpAssigner interface.
type MockIpAssigner struct {
	ctrl     *gomock.Controller
	recorder *MockIpAssignerMockRecorder
}

// MockIpAssignerMockRecorder is the mock recorder for MockIpAssigner.
type MockIpAssignerMockRecorder struct {
	mock *MockIpAssigner
}

// NewMockIpAssigner creates a new mock instance.
func NewMockIpAssigner(ctrl *gomock.Controller) *MockIpAssigner {
	mock := &MockIpAssigner{ctrl: ctrl}
	mock.recorder = &MockIpAssignerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIpAssigner) EXPECT() *MockIpAssignerMockRecorder {
	return m.recorder
}

// AssignIPOnCluster mocks base method.
func (m *MockIpAssigner) AssignIPOnCluster(ctx context.Context, clusterName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignIPOnCluster", ctx, clusterName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignIPOnCluster indicates an expected call of AssignIPOnCluster.
func (mr *MockIpAssignerMockRecorder) AssignIPOnCluster(ctx, clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignIPOnCluster", reflect.TypeOf((*MockIpAssigner)(nil).AssignIPOnCluster), ctx, clusterName)
}

// UnAssignIPOnCluster mocks base method.
func (m *MockIpAssigner) UnAssignIPOnCluster(ctx context.Context, clusterName, ip string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnAssignIPOnCluster", ctx, clusterName, ip)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnAssignIPOnCluster indicates an expected call of UnAssignIPOnCluster.
func (mr *MockIpAssignerMockRecorder) UnAssignIPOnCluster(ctx, clusterName, ip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnAssignIPOnCluster", reflect.TypeOf((*MockIpAssigner)(nil).UnAssignIPOnCluster), ctx, clusterName, ip)
}

// MockExternalAccessPointGetter is a mock of ExternalAccessPointGetter interface.
type MockExternalAccessPointGetter struct {
	ctrl     *gomock.Controller
	recorder *MockExternalAccessPointGetterMockRecorder
}

// MockExternalAccessPointGetterMockRecorder is the mock recorder for MockExternalAccessPointGetter.
type MockExternalAccessPointGetterMockRecorder struct {
	mock *MockExternalAccessPointGetter
}

// NewMockExternalAccessPointGetter creates a new mock instance.
func NewMockExternalAccessPointGetter(ctrl *gomock.Controller) *MockExternalAccessPointGetter {
	mock := &MockExternalAccessPointGetter{ctrl: ctrl}
	mock.recorder = &MockExternalAccessPointGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExternalAccessPointGetter) EXPECT() *MockExternalAccessPointGetterMockRecorder {
	return m.recorder
}

// GetExternalAccessPointForService mocks base method.
func (m *MockExternalAccessPointGetter) GetExternalAccessPointForService(ctx context.Context, svc *v1.Service, portName, clusterName string) (dns.ExternalAccessPoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExternalAccessPointForService", ctx, svc, portName, clusterName)
	ret0, _ := ret[0].(dns.ExternalAccessPoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExternalAccessPointForService indicates an expected call of GetExternalAccessPointForService.
func (mr *MockExternalAccessPointGetterMockRecorder) GetExternalAccessPointForService(ctx, svc, portName, clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExternalAccessPointForService", reflect.TypeOf((*MockExternalAccessPointGetter)(nil).GetExternalAccessPointForService), ctx, svc, portName, clusterName)
}