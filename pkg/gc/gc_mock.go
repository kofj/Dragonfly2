// Code generated by MockGen. DO NOT EDIT.
// Source: gc.go

// Package gc is a generated GoMock package.
package gc

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGC is a mock of GC interface.
type MockGC struct {
	ctrl     *gomock.Controller
	recorder *MockGCMockRecorder
}

// MockGCMockRecorder is the mock recorder for MockGC.
type MockGCMockRecorder struct {
	mock *MockGC
}

// NewMockGC creates a new mock instance.
func NewMockGC(ctrl *gomock.Controller) *MockGC {
	mock := &MockGC{ctrl: ctrl}
	mock.recorder = &MockGCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGC) EXPECT() *MockGCMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockGC) Add(arg0 Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockGCMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockGC)(nil).Add), arg0)
}

// Run mocks base method.
func (m *MockGC) Run(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockGCMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockGC)(nil).Run), arg0)
}

// RunAll mocks base method.
func (m *MockGC) RunAll() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RunAll")
}

// RunAll indicates an expected call of RunAll.
func (mr *MockGCMockRecorder) RunAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunAll", reflect.TypeOf((*MockGC)(nil).RunAll))
}

// Start mocks base method.
func (m *MockGC) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockGCMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockGC)(nil).Start))
}

// Stop mocks base method.
func (m *MockGC) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockGCMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockGC)(nil).Stop))
}
