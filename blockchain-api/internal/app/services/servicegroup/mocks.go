// Code generated by MockGen. DO NOT EDIT.
// Source: ./repositories.go

// Package servicegroup is a generated GoMock package.
package servicegroup

import (
	context "context"
	reflect "reflect"

	models "github.com/cronnoss/blockchain-api/blockchain-api/internal/pkg/models"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetGroup mocks base method.
func (m *MockRepository) GetGroup(arg0 context.Context, arg1 int64) (*models.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroup", arg0, arg1)
	ret0, _ := ret[0].(*models.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroup indicates an expected call of GetGroup.
func (mr *MockRepositoryMockRecorder) GetGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroup", reflect.TypeOf((*MockRepository)(nil).GetGroup), arg0, arg1)
}

// GetGroupIDs mocks base method.
func (m *MockRepository) GetGroupIDs(arg0 context.Context) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupIDs", arg0)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupIDs indicates an expected call of GetGroupIDs.
func (mr *MockRepositoryMockRecorder) GetGroupIDs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupIDs", reflect.TypeOf((*MockRepository)(nil).GetGroupIDs), arg0)
}
