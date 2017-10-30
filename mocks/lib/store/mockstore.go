// Code generated by MockGen. DO NOT EDIT.
// Source: code.uber.internal/infra/kraken/lib/store (interfaces: FileStore,FileReadWriter)

// Package mockstore is a generated GoMock package.
package mockstore

import (
	store "code.uber.internal/infra/kraken/lib/store"
	base "code.uber.internal/infra/kraken/lib/store/base"
	gomock "github.com/golang/mock/gomock"
	io "io"
	os "os"
	reflect "reflect"
)

// MockFileStore is a mock of FileStore interface
type MockFileStore struct {
	ctrl     *gomock.Controller
	recorder *MockFileStoreMockRecorder
}

// MockFileStoreMockRecorder is the mock recorder for MockFileStore
type MockFileStoreMockRecorder struct {
	mock *MockFileStore
}

// NewMockFileStore creates a new mock instance
func NewMockFileStore(ctrl *gomock.Controller) *MockFileStore {
	mock := &MockFileStore{ctrl: ctrl}
	mock.recorder = &MockFileStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileStore) EXPECT() *MockFileStoreMockRecorder {
	return m.recorder
}

// Config mocks base method
func (m *MockFileStore) Config() store.Config {
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(store.Config)
	return ret0
}

// Config indicates an expected call of Config
func (mr *MockFileStoreMockRecorder) Config() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockFileStore)(nil).Config))
}

// CreateCacheFile mocks base method
func (m *MockFileStore) CreateCacheFile(arg0 string, arg1 io.Reader) error {
	ret := m.ctrl.Call(m, "CreateCacheFile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCacheFile indicates an expected call of CreateCacheFile
func (mr *MockFileStoreMockRecorder) CreateCacheFile(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCacheFile", reflect.TypeOf((*MockFileStore)(nil).CreateCacheFile), arg0, arg1)
}

// CreateDownloadFile mocks base method
func (m *MockFileStore) CreateDownloadFile(arg0 string, arg1 int64) error {
	ret := m.ctrl.Call(m, "CreateDownloadFile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDownloadFile indicates an expected call of CreateDownloadFile
func (mr *MockFileStoreMockRecorder) CreateDownloadFile(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDownloadFile", reflect.TypeOf((*MockFileStore)(nil).CreateDownloadFile), arg0, arg1)
}

// CreateUploadFile mocks base method
func (m *MockFileStore) CreateUploadFile(arg0 string, arg1 int64) error {
	ret := m.ctrl.Call(m, "CreateUploadFile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUploadFile indicates an expected call of CreateUploadFile
func (mr *MockFileStoreMockRecorder) CreateUploadFile(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUploadFile", reflect.TypeOf((*MockFileStore)(nil).CreateUploadFile), arg0, arg1)
}

// DeleteAllTrashFiles mocks base method
func (m *MockFileStore) DeleteAllTrashFiles() error {
	ret := m.ctrl.Call(m, "DeleteAllTrashFiles")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllTrashFiles indicates an expected call of DeleteAllTrashFiles
func (mr *MockFileStoreMockRecorder) DeleteAllTrashFiles() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllTrashFiles", reflect.TypeOf((*MockFileStore)(nil).DeleteAllTrashFiles))
}

// DeleteUploadFileStartedAt mocks base method
func (m *MockFileStore) DeleteUploadFileStartedAt(arg0 string) error {
	ret := m.ctrl.Call(m, "DeleteUploadFileStartedAt", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUploadFileStartedAt indicates an expected call of DeleteUploadFileStartedAt
func (mr *MockFileStoreMockRecorder) DeleteUploadFileStartedAt(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUploadFileStartedAt", reflect.TypeOf((*MockFileStore)(nil).DeleteUploadFileStartedAt), arg0)
}

// DerefCacheFile mocks base method
func (m *MockFileStore) DerefCacheFile(arg0 string) (int64, error) {
	ret := m.ctrl.Call(m, "DerefCacheFile", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DerefCacheFile indicates an expected call of DerefCacheFile
func (mr *MockFileStoreMockRecorder) DerefCacheFile(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DerefCacheFile", reflect.TypeOf((*MockFileStore)(nil).DerefCacheFile), arg0)
}

// GetCacheFilePath mocks base method
func (m *MockFileStore) GetCacheFilePath(arg0 string) (string, error) {
	ret := m.ctrl.Call(m, "GetCacheFilePath", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCacheFilePath indicates an expected call of GetCacheFilePath
func (mr *MockFileStoreMockRecorder) GetCacheFilePath(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCacheFilePath", reflect.TypeOf((*MockFileStore)(nil).GetCacheFilePath), arg0)
}

// GetCacheFileReader mocks base method
func (m *MockFileStore) GetCacheFileReader(arg0 string) (base.FileReader, error) {
	ret := m.ctrl.Call(m, "GetCacheFileReader", arg0)
	ret0, _ := ret[0].(base.FileReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCacheFileReader indicates an expected call of GetCacheFileReader
func (mr *MockFileStoreMockRecorder) GetCacheFileReader(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCacheFileReader", reflect.TypeOf((*MockFileStore)(nil).GetCacheFileReader), arg0)
}

// GetCacheFileStat mocks base method
func (m *MockFileStore) GetCacheFileStat(arg0 string) (os.FileInfo, error) {
	ret := m.ctrl.Call(m, "GetCacheFileStat", arg0)
	ret0, _ := ret[0].(os.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCacheFileStat indicates an expected call of GetCacheFileStat
func (mr *MockFileStoreMockRecorder) GetCacheFileStat(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCacheFileStat", reflect.TypeOf((*MockFileStore)(nil).GetCacheFileStat), arg0)
}

// GetDownloadFileReadWriter mocks base method
func (m *MockFileStore) GetDownloadFileReadWriter(arg0 string) (base.FileReadWriter, error) {
	ret := m.ctrl.Call(m, "GetDownloadFileReadWriter", arg0)
	ret0, _ := ret[0].(base.FileReadWriter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDownloadFileReadWriter indicates an expected call of GetDownloadFileReadWriter
func (mr *MockFileStoreMockRecorder) GetDownloadFileReadWriter(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDownloadFileReadWriter", reflect.TypeOf((*MockFileStore)(nil).GetDownloadFileReadWriter), arg0)
}

// GetDownloadOrCacheFileMeta mocks base method
func (m *MockFileStore) GetDownloadOrCacheFileMeta(arg0 string) ([]byte, error) {
	ret := m.ctrl.Call(m, "GetDownloadOrCacheFileMeta", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDownloadOrCacheFileMeta indicates an expected call of GetDownloadOrCacheFileMeta
func (mr *MockFileStoreMockRecorder) GetDownloadOrCacheFileMeta(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDownloadOrCacheFileMeta", reflect.TypeOf((*MockFileStore)(nil).GetDownloadOrCacheFileMeta), arg0)
}

// GetDownloadOrCacheFileReader mocks base method
func (m *MockFileStore) GetDownloadOrCacheFileReader(arg0 string) (base.FileReader, error) {
	ret := m.ctrl.Call(m, "GetDownloadOrCacheFileReader", arg0)
	ret0, _ := ret[0].(base.FileReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDownloadOrCacheFileReader indicates an expected call of GetDownloadOrCacheFileReader
func (mr *MockFileStoreMockRecorder) GetDownloadOrCacheFileReader(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDownloadOrCacheFileReader", reflect.TypeOf((*MockFileStore)(nil).GetDownloadOrCacheFileReader), arg0)
}

// GetFilePieceStatus mocks base method
func (m *MockFileStore) GetFilePieceStatus(arg0 string, arg1, arg2 int) ([]byte, error) {
	ret := m.ctrl.Call(m, "GetFilePieceStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilePieceStatus indicates an expected call of GetFilePieceStatus
func (mr *MockFileStoreMockRecorder) GetFilePieceStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilePieceStatus", reflect.TypeOf((*MockFileStore)(nil).GetFilePieceStatus), arg0, arg1, arg2)
}

// GetUploadFileHashState mocks base method
func (m *MockFileStore) GetUploadFileHashState(arg0, arg1, arg2 string) ([]byte, error) {
	ret := m.ctrl.Call(m, "GetUploadFileHashState", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUploadFileHashState indicates an expected call of GetUploadFileHashState
func (mr *MockFileStoreMockRecorder) GetUploadFileHashState(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUploadFileHashState", reflect.TypeOf((*MockFileStore)(nil).GetUploadFileHashState), arg0, arg1, arg2)
}

// GetUploadFileReadWriter mocks base method
func (m *MockFileStore) GetUploadFileReadWriter(arg0 string) (base.FileReadWriter, error) {
	ret := m.ctrl.Call(m, "GetUploadFileReadWriter", arg0)
	ret0, _ := ret[0].(base.FileReadWriter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUploadFileReadWriter indicates an expected call of GetUploadFileReadWriter
func (mr *MockFileStoreMockRecorder) GetUploadFileReadWriter(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUploadFileReadWriter", reflect.TypeOf((*MockFileStore)(nil).GetUploadFileReadWriter), arg0)
}

// GetUploadFileReader mocks base method
func (m *MockFileStore) GetUploadFileReader(arg0 string) (base.FileReader, error) {
	ret := m.ctrl.Call(m, "GetUploadFileReader", arg0)
	ret0, _ := ret[0].(base.FileReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUploadFileReader indicates an expected call of GetUploadFileReader
func (mr *MockFileStoreMockRecorder) GetUploadFileReader(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUploadFileReader", reflect.TypeOf((*MockFileStore)(nil).GetUploadFileReader), arg0)
}

// GetUploadFileStartedAt mocks base method
func (m *MockFileStore) GetUploadFileStartedAt(arg0 string) ([]byte, error) {
	ret := m.ctrl.Call(m, "GetUploadFileStartedAt", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUploadFileStartedAt indicates an expected call of GetUploadFileStartedAt
func (mr *MockFileStoreMockRecorder) GetUploadFileStartedAt(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUploadFileStartedAt", reflect.TypeOf((*MockFileStore)(nil).GetUploadFileStartedAt), arg0)
}

// GetUploadFileStat mocks base method
func (m *MockFileStore) GetUploadFileStat(arg0 string) (os.FileInfo, error) {
	ret := m.ctrl.Call(m, "GetUploadFileStat", arg0)
	ret0, _ := ret[0].(os.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUploadFileStat indicates an expected call of GetUploadFileStat
func (mr *MockFileStoreMockRecorder) GetUploadFileStat(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUploadFileStat", reflect.TypeOf((*MockFileStore)(nil).GetUploadFileStat), arg0)
}

// ListCacheFilesByShardID mocks base method
func (m *MockFileStore) ListCacheFilesByShardID(arg0 string) ([]string, error) {
	ret := m.ctrl.Call(m, "ListCacheFilesByShardID", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCacheFilesByShardID indicates an expected call of ListCacheFilesByShardID
func (mr *MockFileStoreMockRecorder) ListCacheFilesByShardID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCacheFilesByShardID", reflect.TypeOf((*MockFileStore)(nil).ListCacheFilesByShardID), arg0)
}

// ListPopulatedShardIDs mocks base method
func (m *MockFileStore) ListPopulatedShardIDs() ([]string, error) {
	ret := m.ctrl.Call(m, "ListPopulatedShardIDs")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPopulatedShardIDs indicates an expected call of ListPopulatedShardIDs
func (mr *MockFileStoreMockRecorder) ListPopulatedShardIDs() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPopulatedShardIDs", reflect.TypeOf((*MockFileStore)(nil).ListPopulatedShardIDs))
}

// ListUploadFileHashStatePaths mocks base method
func (m *MockFileStore) ListUploadFileHashStatePaths(arg0 string) ([]string, error) {
	ret := m.ctrl.Call(m, "ListUploadFileHashStatePaths", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUploadFileHashStatePaths indicates an expected call of ListUploadFileHashStatePaths
func (mr *MockFileStoreMockRecorder) ListUploadFileHashStatePaths(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUploadFileHashStatePaths", reflect.TypeOf((*MockFileStore)(nil).ListUploadFileHashStatePaths), arg0)
}

// MoveCacheFileToTrash mocks base method
func (m *MockFileStore) MoveCacheFileToTrash(arg0 string) error {
	ret := m.ctrl.Call(m, "MoveCacheFileToTrash", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveCacheFileToTrash indicates an expected call of MoveCacheFileToTrash
func (mr *MockFileStoreMockRecorder) MoveCacheFileToTrash(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveCacheFileToTrash", reflect.TypeOf((*MockFileStore)(nil).MoveCacheFileToTrash), arg0)
}

// MoveDownloadFileToCache mocks base method
func (m *MockFileStore) MoveDownloadFileToCache(arg0 string) error {
	ret := m.ctrl.Call(m, "MoveDownloadFileToCache", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveDownloadFileToCache indicates an expected call of MoveDownloadFileToCache
func (mr *MockFileStoreMockRecorder) MoveDownloadFileToCache(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveDownloadFileToCache", reflect.TypeOf((*MockFileStore)(nil).MoveDownloadFileToCache), arg0)
}

// MoveDownloadOrCacheFileToTrash mocks base method
func (m *MockFileStore) MoveDownloadOrCacheFileToTrash(arg0 string) error {
	ret := m.ctrl.Call(m, "MoveDownloadOrCacheFileToTrash", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveDownloadOrCacheFileToTrash indicates an expected call of MoveDownloadOrCacheFileToTrash
func (mr *MockFileStoreMockRecorder) MoveDownloadOrCacheFileToTrash(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveDownloadOrCacheFileToTrash", reflect.TypeOf((*MockFileStore)(nil).MoveDownloadOrCacheFileToTrash), arg0)
}

// MoveUploadFileToCache mocks base method
func (m *MockFileStore) MoveUploadFileToCache(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "MoveUploadFileToCache", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveUploadFileToCache indicates an expected call of MoveUploadFileToCache
func (mr *MockFileStoreMockRecorder) MoveUploadFileToCache(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveUploadFileToCache", reflect.TypeOf((*MockFileStore)(nil).MoveUploadFileToCache), arg0, arg1)
}

// RefCacheFile mocks base method
func (m *MockFileStore) RefCacheFile(arg0 string) (int64, error) {
	ret := m.ctrl.Call(m, "RefCacheFile", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefCacheFile indicates an expected call of RefCacheFile
func (mr *MockFileStoreMockRecorder) RefCacheFile(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefCacheFile", reflect.TypeOf((*MockFileStore)(nil).RefCacheFile), arg0)
}

// SetDownloadOrCacheFileMeta mocks base method
func (m *MockFileStore) SetDownloadOrCacheFileMeta(arg0 string, arg1 []byte) (bool, error) {
	ret := m.ctrl.Call(m, "SetDownloadOrCacheFileMeta", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetDownloadOrCacheFileMeta indicates an expected call of SetDownloadOrCacheFileMeta
func (mr *MockFileStoreMockRecorder) SetDownloadOrCacheFileMeta(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDownloadOrCacheFileMeta", reflect.TypeOf((*MockFileStore)(nil).SetDownloadOrCacheFileMeta), arg0, arg1)
}

// SetUploadFileHashState mocks base method
func (m *MockFileStore) SetUploadFileHashState(arg0 string, arg1 []byte, arg2, arg3 string) error {
	ret := m.ctrl.Call(m, "SetUploadFileHashState", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUploadFileHashState indicates an expected call of SetUploadFileHashState
func (mr *MockFileStoreMockRecorder) SetUploadFileHashState(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUploadFileHashState", reflect.TypeOf((*MockFileStore)(nil).SetUploadFileHashState), arg0, arg1, arg2, arg3)
}

// SetUploadFileStartedAt mocks base method
func (m *MockFileStore) SetUploadFileStartedAt(arg0 string, arg1 []byte) error {
	ret := m.ctrl.Call(m, "SetUploadFileStartedAt", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUploadFileStartedAt indicates an expected call of SetUploadFileStartedAt
func (mr *MockFileStoreMockRecorder) SetUploadFileStartedAt(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUploadFileStartedAt", reflect.TypeOf((*MockFileStore)(nil).SetUploadFileStartedAt), arg0, arg1)
}

// Stop mocks base method
func (m *MockFileStore) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockFileStoreMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockFileStore)(nil).Stop))
}

// WriteDownloadFilePieceStatus mocks base method
func (m *MockFileStore) WriteDownloadFilePieceStatus(arg0 string, arg1 []byte) (bool, error) {
	ret := m.ctrl.Call(m, "WriteDownloadFilePieceStatus", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteDownloadFilePieceStatus indicates an expected call of WriteDownloadFilePieceStatus
func (mr *MockFileStoreMockRecorder) WriteDownloadFilePieceStatus(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteDownloadFilePieceStatus", reflect.TypeOf((*MockFileStore)(nil).WriteDownloadFilePieceStatus), arg0, arg1)
}

// WriteDownloadFilePieceStatusAt mocks base method
func (m *MockFileStore) WriteDownloadFilePieceStatusAt(arg0 string, arg1 []byte, arg2 int) (bool, error) {
	ret := m.ctrl.Call(m, "WriteDownloadFilePieceStatusAt", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteDownloadFilePieceStatusAt indicates an expected call of WriteDownloadFilePieceStatusAt
func (mr *MockFileStoreMockRecorder) WriteDownloadFilePieceStatusAt(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteDownloadFilePieceStatusAt", reflect.TypeOf((*MockFileStore)(nil).WriteDownloadFilePieceStatusAt), arg0, arg1, arg2)
}

// MockFileReadWriter is a mock of FileReadWriter interface
type MockFileReadWriter struct {
	ctrl     *gomock.Controller
	recorder *MockFileReadWriterMockRecorder
}

// MockFileReadWriterMockRecorder is the mock recorder for MockFileReadWriter
type MockFileReadWriterMockRecorder struct {
	mock *MockFileReadWriter
}

// NewMockFileReadWriter creates a new mock instance
func NewMockFileReadWriter(ctrl *gomock.Controller) *MockFileReadWriter {
	mock := &MockFileReadWriter{ctrl: ctrl}
	mock.recorder = &MockFileReadWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileReadWriter) EXPECT() *MockFileReadWriterMockRecorder {
	return m.recorder
}

// Cancel mocks base method
func (m *MockFileReadWriter) Cancel() error {
	ret := m.ctrl.Call(m, "Cancel")
	ret0, _ := ret[0].(error)
	return ret0
}

// Cancel indicates an expected call of Cancel
func (mr *MockFileReadWriterMockRecorder) Cancel() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockFileReadWriter)(nil).Cancel))
}

// Close mocks base method
func (m *MockFileReadWriter) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockFileReadWriterMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockFileReadWriter)(nil).Close))
}

// Commit mocks base method
func (m *MockFileReadWriter) Commit() error {
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockFileReadWriterMockRecorder) Commit() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockFileReadWriter)(nil).Commit))
}

// Read mocks base method
func (m *MockFileReadWriter) Read(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockFileReadWriterMockRecorder) Read(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockFileReadWriter)(nil).Read), arg0)
}

// ReadAt mocks base method
func (m *MockFileReadWriter) ReadAt(arg0 []byte, arg1 int64) (int, error) {
	ret := m.ctrl.Call(m, "ReadAt", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAt indicates an expected call of ReadAt
func (mr *MockFileReadWriterMockRecorder) ReadAt(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAt", reflect.TypeOf((*MockFileReadWriter)(nil).ReadAt), arg0, arg1)
}

// Seek mocks base method
func (m *MockFileReadWriter) Seek(arg0 int64, arg1 int) (int64, error) {
	ret := m.ctrl.Call(m, "Seek", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Seek indicates an expected call of Seek
func (mr *MockFileReadWriterMockRecorder) Seek(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seek", reflect.TypeOf((*MockFileReadWriter)(nil).Seek), arg0, arg1)
}

// Size mocks base method
func (m *MockFileReadWriter) Size() int64 {
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockFileReadWriterMockRecorder) Size() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockFileReadWriter)(nil).Size))
}

// Write mocks base method
func (m *MockFileReadWriter) Write(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockFileReadWriterMockRecorder) Write(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockFileReadWriter)(nil).Write), arg0)
}

// WriteAt mocks base method
func (m *MockFileReadWriter) WriteAt(arg0 []byte, arg1 int64) (int, error) {
	ret := m.ctrl.Call(m, "WriteAt", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteAt indicates an expected call of WriteAt
func (mr *MockFileReadWriterMockRecorder) WriteAt(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteAt", reflect.TypeOf((*MockFileReadWriter)(nil).WriteAt), arg0, arg1)
}
