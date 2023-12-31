// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/serod11/gofiber-boilerplate/pkg/repo (interfaces: BookRepo)
//
// Generated by this command:
//
//	mockgen -package mock -destination test/mock/mock_book_repo.go github.com/serod11/gofiber-boilerplate/pkg/repo BookRepo
//
// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	model "github.com/serod11/gofiber-boilerplate/pkg/model"
	gomock "go.uber.org/mock/gomock"
)

// MockBookRepo is a mock of BookRepo interface.
type MockBookRepo struct {
	ctrl     *gomock.Controller
	recorder *MockBookRepoMockRecorder
}

// MockBookRepoMockRecorder is the mock recorder for MockBookRepo.
type MockBookRepoMockRecorder struct {
	mock *MockBookRepo
}

// NewMockBookRepo creates a new mock instance.
func NewMockBookRepo(ctrl *gomock.Controller) *MockBookRepo {
	mock := &MockBookRepo{ctrl: ctrl}
	mock.recorder = &MockBookRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookRepo) EXPECT() *MockBookRepoMockRecorder {
	return m.recorder
}

// FindAll mocks base method.
func (m *MockBookRepo) FindAll() []model.Book {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]model.Book)
	return ret0
}

// FindAll indicates an expected call of FindAll.
func (mr *MockBookRepoMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockBookRepo)(nil).FindAll))
}

// FindById mocks base method.
func (m *MockBookRepo) FindById(arg0 uint) model.Book {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0)
	ret0, _ := ret[0].(model.Book)
	return ret0
}

// FindById indicates an expected call of FindById.
func (mr *MockBookRepoMockRecorder) FindById(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockBookRepo)(nil).FindById), arg0)
}
