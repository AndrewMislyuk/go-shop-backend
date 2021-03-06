// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	domain "github.com/AndrewMislyuk/go-shop-backend/internal/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUser) CreateUser(user domain.UserSignUp) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUser)(nil).CreateUser), user)
}

// GenerateToken mocks base method.
func (m *MockUser) GenerateToken(email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockUserMockRecorder) GenerateToken(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockUser)(nil).GenerateToken), email, password)
}

// GetMe mocks base method.
func (m *MockUser) GetMe(token string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMe", token)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMe indicates an expected call of GetMe.
func (mr *MockUserMockRecorder) GetMe(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMe", reflect.TypeOf((*MockUser)(nil).GetMe), token)
}

// MockProductsList is a mock of ProductsList interface.
type MockProductsList struct {
	ctrl     *gomock.Controller
	recorder *MockProductsListMockRecorder
}

// MockProductsListMockRecorder is the mock recorder for MockProductsList.
type MockProductsListMockRecorder struct {
	mock *MockProductsList
}

// NewMockProductsList creates a new mock instance.
func NewMockProductsList(ctrl *gomock.Controller) *MockProductsList {
	mock := &MockProductsList{ctrl: ctrl}
	mock.recorder = &MockProductsListMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductsList) EXPECT() *MockProductsListMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductsList) Create(list domain.CreateProductInput) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", list)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductsListMockRecorder) Create(list interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductsList)(nil).Create), list)
}

// Delete mocks base method.
func (m *MockProductsList) Delete(itemId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", itemId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockProductsListMockRecorder) Delete(itemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProductsList)(nil).Delete), itemId)
}

// GetAll mocks base method.
func (m *MockProductsList) GetAll() ([]domain.ProductsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]domain.ProductsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockProductsListMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockProductsList)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockProductsList) GetById(listId string) (domain.ProductsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", listId)
	ret0, _ := ret[0].(domain.ProductsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockProductsListMockRecorder) GetById(listId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockProductsList)(nil).GetById), listId)
}

// Update mocks base method.
func (m *MockProductsList) Update(itemId string, input domain.UpdateProductInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", itemId, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockProductsListMockRecorder) Update(itemId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductsList)(nil).Update), itemId, input)
}

// MockFiles is a mock of Files interface.
type MockFiles struct {
	ctrl     *gomock.Controller
	recorder *MockFilesMockRecorder
}

// MockFilesMockRecorder is the mock recorder for MockFiles.
type MockFilesMockRecorder struct {
	mock *MockFiles
}

// NewMockFiles creates a new mock instance.
func NewMockFiles(ctrl *gomock.Controller) *MockFiles {
	mock := &MockFiles{ctrl: ctrl}
	mock.recorder = &MockFilesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFiles) EXPECT() *MockFilesMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockFiles) Delete() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockFilesMockRecorder) Delete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFiles)(nil).Delete))
}

// Get mocks base method.
func (m *MockFiles) Get() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockFilesMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockFiles)(nil).Get))
}

// Upload mocks base method.
func (m *MockFiles) Upload() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload")
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload.
func (mr *MockFilesMockRecorder) Upload() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockFiles)(nil).Upload))
}
