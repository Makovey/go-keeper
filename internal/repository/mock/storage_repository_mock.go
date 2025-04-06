// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go

// Package mock is a generated GoMock package.
package mock

import (
	bufio "bufio"
	context "context"
	reflect "reflect"

	entity "github.com/Makovey/go-keeper/internal/repository/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockFileStorager is a mock of FileStorager interface.
type MockFileStorager struct {
	ctrl     *gomock.Controller
	recorder *MockFileStoragerMockRecorder
}

// MockFileStoragerMockRecorder is the mock recorder for MockFileStorager.
type MockFileStoragerMockRecorder struct {
	mock *MockFileStorager
}

// NewMockFileStorager creates a new mock instance.
func NewMockFileStorager(ctrl *gomock.Controller) *MockFileStorager {
	mock := &MockFileStorager{ctrl: ctrl}
	mock.recorder = &MockFileStoragerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileStorager) EXPECT() *MockFileStoragerMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockFileStorager) Get(path string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", path)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockFileStoragerMockRecorder) Get(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockFileStorager)(nil).Get), path)
}

// Save mocks base method.
func (m *MockFileStorager) Save(path, fileName string, data *bufio.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", path, fileName, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockFileStoragerMockRecorder) Save(path, fileName, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockFileStorager)(nil).Save), path, fileName, data)
}

// MockRepositoryStorage is a mock of RepositoryStorage interface.
type MockRepositoryStorage struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryStorageMockRecorder
}

// MockRepositoryStorageMockRecorder is the mock recorder for MockRepositoryStorage.
type MockRepositoryStorageMockRecorder struct {
	mock *MockRepositoryStorage
}

// NewMockRepositoryStorage creates a new mock instance.
func NewMockRepositoryStorage(ctrl *gomock.Controller) *MockRepositoryStorage {
	mock := &MockRepositoryStorage{ctrl: ctrl}
	mock.recorder = &MockRepositoryStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryStorage) EXPECT() *MockRepositoryStorageMockRecorder {
	return m.recorder
}

// GetFileMetadata mocks base method.
func (m *MockRepositoryStorage) GetFileMetadata(ctx context.Context, userID, fileID string) (*entity.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileMetadata", ctx, userID, fileID)
	ret0, _ := ret[0].(*entity.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileMetadata indicates an expected call of GetFileMetadata.
func (mr *MockRepositoryStorageMockRecorder) GetFileMetadata(ctx, userID, fileID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileMetadata", reflect.TypeOf((*MockRepositoryStorage)(nil).GetFileMetadata), ctx, userID, fileID)
}

// GetUsersFiles mocks base method.
func (m *MockRepositoryStorage) GetUsersFiles(ctx context.Context, userID string) ([]*entity.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersFiles", ctx, userID)
	ret0, _ := ret[0].([]*entity.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersFiles indicates an expected call of GetUsersFiles.
func (mr *MockRepositoryStorageMockRecorder) GetUsersFiles(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersFiles", reflect.TypeOf((*MockRepositoryStorage)(nil).GetUsersFiles), ctx, userID)
}

// SaveFileMetadata mocks base method.
func (m *MockRepositoryStorage) SaveFileMetadata(ctx context.Context, fileData *entity.File) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveFileMetadata", ctx, fileData)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveFileMetadata indicates an expected call of SaveFileMetadata.
func (mr *MockRepositoryStorageMockRecorder) SaveFileMetadata(ctx, fileData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveFileMetadata", reflect.TypeOf((*MockRepositoryStorage)(nil).SaveFileMetadata), ctx, fileData)
}
