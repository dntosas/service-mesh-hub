// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1alpha1sets is a generated GoMock package.
package mock_v1alpha1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/security.smh.solo.io/v1alpha1"
	v1alpha1sets "github.com/solo-io/service-mesh-hub/pkg/api/security.smh.solo.io/v1alpha1/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// MockVirtualMeshCertificateSigningRequestSet is a mock of VirtualMeshCertificateSigningRequestSet interface.
type MockVirtualMeshCertificateSigningRequestSet struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMeshCertificateSigningRequestSetMockRecorder
}

// MockVirtualMeshCertificateSigningRequestSetMockRecorder is the mock recorder for MockVirtualMeshCertificateSigningRequestSet.
type MockVirtualMeshCertificateSigningRequestSetMockRecorder struct {
	mock *MockVirtualMeshCertificateSigningRequestSet
}

// NewMockVirtualMeshCertificateSigningRequestSet creates a new mock instance.
func NewMockVirtualMeshCertificateSigningRequestSet(ctrl *gomock.Controller) *MockVirtualMeshCertificateSigningRequestSet {
	mock := &MockVirtualMeshCertificateSigningRequestSet{ctrl: ctrl}
	mock.recorder = &MockVirtualMeshCertificateSigningRequestSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMeshCertificateSigningRequestSet) EXPECT() *MockVirtualMeshCertificateSigningRequestSetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockVirtualMeshCertificateSigningRequestSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestSet)(nil).Keys))
}

// List mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestSet) List() []*v1alpha1.VirtualMeshCertificateSigningRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1alpha1.VirtualMeshCertificateSigningRequest)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockVirtualMeshCertificateSigningRequestSetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestSet)(nil).List))
}

// Map mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestSet) Map() map[string]*v1alpha1.VirtualMeshCertificateSigningRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha1.VirtualMeshCertificateSigningRequest)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockVirtualMeshCertificateSigningRequestSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestSet)(nil).Map))
}

// Insert mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestSet) Insert(virtualMeshCertificateSigningRequest ...*v1alpha1.VirtualMeshCertificateSigningRequest) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range virtualMeshCertificateSigningRequest {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockVirtualMeshCertificateSigningRequestSetMockRecorder) Insert(virtualMeshCertificateSigningRequest ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestSet)(nil).Insert), virtualMeshCertificateSigningRequest...)
}

// Equal mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestSet) Equal(virtualMeshCertificateSigningRequestSet v1alpha1sets.VirtualMeshCertificateSigningRequestSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", virtualMeshCertificateSigningRequestSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockVirtualMeshCertificateSigningRequestSetMockRecorder) Equal(virtualMeshCertificateSigningRequestSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestSet)(nil).Equal), virtualMeshCertificateSigningRequestSet)
}

// Has mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestSet) Has(virtualMeshCertificateSigningRequest *v1alpha1.VirtualMeshCertificateSigningRequest) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", virtualMeshCertificateSigningRequest)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockVirtualMeshCertificateSigningRequestSetMockRecorder) Has(virtualMeshCertificateSigningRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestSet)(nil).Has), virtualMeshCertificateSigningRequest)
}

// Delete mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestSet) Delete(virtualMeshCertificateSigningRequest *v1alpha1.VirtualMeshCertificateSigningRequest) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", virtualMeshCertificateSigningRequest)
}

// Delete indicates an expected call of Delete.
func (mr *MockVirtualMeshCertificateSigningRequestSetMockRecorder) Delete(virtualMeshCertificateSigningRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestSet)(nil).Delete), virtualMeshCertificateSigningRequest)
}

// Union mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestSet) Union(set v1alpha1sets.VirtualMeshCertificateSigningRequestSet) v1alpha1sets.VirtualMeshCertificateSigningRequestSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha1sets.VirtualMeshCertificateSigningRequestSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockVirtualMeshCertificateSigningRequestSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestSet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestSet) Difference(set v1alpha1sets.VirtualMeshCertificateSigningRequestSet) v1alpha1sets.VirtualMeshCertificateSigningRequestSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha1sets.VirtualMeshCertificateSigningRequestSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockVirtualMeshCertificateSigningRequestSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestSet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestSet) Intersection(set v1alpha1sets.VirtualMeshCertificateSigningRequestSet) v1alpha1sets.VirtualMeshCertificateSigningRequestSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha1sets.VirtualMeshCertificateSigningRequestSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockVirtualMeshCertificateSigningRequestSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestSet)(nil).Intersection), set)
}

// Find mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestSet) Find(id ezkube.ResourceId) (*v1alpha1.VirtualMeshCertificateSigningRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1alpha1.VirtualMeshCertificateSigningRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockVirtualMeshCertificateSigningRequestSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestSet)(nil).Find), id)
}

// Length mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockVirtualMeshCertificateSigningRequestSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestSet)(nil).Length))
}
