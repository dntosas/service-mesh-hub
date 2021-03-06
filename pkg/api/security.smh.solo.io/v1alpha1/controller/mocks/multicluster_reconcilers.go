// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/security.smh.solo.io/v1alpha1"
	controller "github.com/solo-io/service-mesh-hub/pkg/api/security.smh.solo.io/v1alpha1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterVirtualMeshCertificateSigningRequestReconciler is a mock of MulticlusterVirtualMeshCertificateSigningRequestReconciler interface.
type MockMulticlusterVirtualMeshCertificateSigningRequestReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterVirtualMeshCertificateSigningRequestReconcilerMockRecorder
}

// MockMulticlusterVirtualMeshCertificateSigningRequestReconcilerMockRecorder is the mock recorder for MockMulticlusterVirtualMeshCertificateSigningRequestReconciler.
type MockMulticlusterVirtualMeshCertificateSigningRequestReconcilerMockRecorder struct {
	mock *MockMulticlusterVirtualMeshCertificateSigningRequestReconciler
}

// NewMockMulticlusterVirtualMeshCertificateSigningRequestReconciler creates a new mock instance.
func NewMockMulticlusterVirtualMeshCertificateSigningRequestReconciler(ctrl *gomock.Controller) *MockMulticlusterVirtualMeshCertificateSigningRequestReconciler {
	mock := &MockMulticlusterVirtualMeshCertificateSigningRequestReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterVirtualMeshCertificateSigningRequestReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterVirtualMeshCertificateSigningRequestReconciler) EXPECT() *MockMulticlusterVirtualMeshCertificateSigningRequestReconcilerMockRecorder {
	return m.recorder
}

// ReconcileVirtualMeshCertificateSigningRequest mocks base method.
func (m *MockMulticlusterVirtualMeshCertificateSigningRequestReconciler) ReconcileVirtualMeshCertificateSigningRequest(clusterName string, obj *v1alpha1.VirtualMeshCertificateSigningRequest) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualMeshCertificateSigningRequest", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileVirtualMeshCertificateSigningRequest indicates an expected call of ReconcileVirtualMeshCertificateSigningRequest.
func (mr *MockMulticlusterVirtualMeshCertificateSigningRequestReconcilerMockRecorder) ReconcileVirtualMeshCertificateSigningRequest(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualMeshCertificateSigningRequest", reflect.TypeOf((*MockMulticlusterVirtualMeshCertificateSigningRequestReconciler)(nil).ReconcileVirtualMeshCertificateSigningRequest), clusterName, obj)
}

// MockMulticlusterVirtualMeshCertificateSigningRequestDeletionReconciler is a mock of MulticlusterVirtualMeshCertificateSigningRequestDeletionReconciler interface.
type MockMulticlusterVirtualMeshCertificateSigningRequestDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterVirtualMeshCertificateSigningRequestDeletionReconcilerMockRecorder
}

// MockMulticlusterVirtualMeshCertificateSigningRequestDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterVirtualMeshCertificateSigningRequestDeletionReconciler.
type MockMulticlusterVirtualMeshCertificateSigningRequestDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterVirtualMeshCertificateSigningRequestDeletionReconciler
}

// NewMockMulticlusterVirtualMeshCertificateSigningRequestDeletionReconciler creates a new mock instance.
func NewMockMulticlusterVirtualMeshCertificateSigningRequestDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterVirtualMeshCertificateSigningRequestDeletionReconciler {
	mock := &MockMulticlusterVirtualMeshCertificateSigningRequestDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterVirtualMeshCertificateSigningRequestDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterVirtualMeshCertificateSigningRequestDeletionReconciler) EXPECT() *MockMulticlusterVirtualMeshCertificateSigningRequestDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileVirtualMeshCertificateSigningRequestDeletion mocks base method.
func (m *MockMulticlusterVirtualMeshCertificateSigningRequestDeletionReconciler) ReconcileVirtualMeshCertificateSigningRequestDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualMeshCertificateSigningRequestDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileVirtualMeshCertificateSigningRequestDeletion indicates an expected call of ReconcileVirtualMeshCertificateSigningRequestDeletion.
func (mr *MockMulticlusterVirtualMeshCertificateSigningRequestDeletionReconcilerMockRecorder) ReconcileVirtualMeshCertificateSigningRequestDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualMeshCertificateSigningRequestDeletion", reflect.TypeOf((*MockMulticlusterVirtualMeshCertificateSigningRequestDeletionReconciler)(nil).ReconcileVirtualMeshCertificateSigningRequestDeletion), clusterName, req)
}

// MockMulticlusterVirtualMeshCertificateSigningRequestReconcileLoop is a mock of MulticlusterVirtualMeshCertificateSigningRequestReconcileLoop interface.
type MockMulticlusterVirtualMeshCertificateSigningRequestReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterVirtualMeshCertificateSigningRequestReconcileLoopMockRecorder
}

// MockMulticlusterVirtualMeshCertificateSigningRequestReconcileLoopMockRecorder is the mock recorder for MockMulticlusterVirtualMeshCertificateSigningRequestReconcileLoop.
type MockMulticlusterVirtualMeshCertificateSigningRequestReconcileLoopMockRecorder struct {
	mock *MockMulticlusterVirtualMeshCertificateSigningRequestReconcileLoop
}

// NewMockMulticlusterVirtualMeshCertificateSigningRequestReconcileLoop creates a new mock instance.
func NewMockMulticlusterVirtualMeshCertificateSigningRequestReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterVirtualMeshCertificateSigningRequestReconcileLoop {
	mock := &MockMulticlusterVirtualMeshCertificateSigningRequestReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterVirtualMeshCertificateSigningRequestReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterVirtualMeshCertificateSigningRequestReconcileLoop) EXPECT() *MockMulticlusterVirtualMeshCertificateSigningRequestReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterVirtualMeshCertificateSigningRequestReconciler mocks base method.
func (m *MockMulticlusterVirtualMeshCertificateSigningRequestReconcileLoop) AddMulticlusterVirtualMeshCertificateSigningRequestReconciler(ctx context.Context, rec controller.MulticlusterVirtualMeshCertificateSigningRequestReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterVirtualMeshCertificateSigningRequestReconciler", varargs...)
}

// AddMulticlusterVirtualMeshCertificateSigningRequestReconciler indicates an expected call of AddMulticlusterVirtualMeshCertificateSigningRequestReconciler.
func (mr *MockMulticlusterVirtualMeshCertificateSigningRequestReconcileLoopMockRecorder) AddMulticlusterVirtualMeshCertificateSigningRequestReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterVirtualMeshCertificateSigningRequestReconciler", reflect.TypeOf((*MockMulticlusterVirtualMeshCertificateSigningRequestReconcileLoop)(nil).AddMulticlusterVirtualMeshCertificateSigningRequestReconciler), varargs...)
}
