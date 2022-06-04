package datastore

import (
	"github.com/golang/mock/gomock"
	"io"
	"reflect"
)

type MockPokemonFileReader struct {
	ctrl     *gomock.Controller
	recorder *MockPokemonFileReaderMockRecorder
}

type MockPokemonFileReaderMockRecorder struct {
	mock *MockPokemonFileReader
}

func NewMockPokemonFileReader(ctrl *gomock.Controller) *MockPokemonFileReader {
	mock := &MockPokemonFileReader{ctrl: ctrl}
	mock.recorder = &MockPokemonFileReaderMockRecorder{mock}
	return mock
}

func (m *MockPokemonFileReader) EXPECT() *MockPokemonFileReaderMockRecorder {
	return m.recorder
}

func (m *MockPokemonFileReader) OpenToRead() (io.ReadWriter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenToRead")
	ret0, _ := ret[0].(io.ReadWriter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPokemonFileReaderMockRecorder) OpenToRead() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenToRead", reflect.TypeOf((*MockPokemonFileReader)(nil).OpenToRead))
}

func (m *MockPokemonFileReader) OpenToWrite() (io.ReadWriter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenToWrite")
	ret0, _ := ret[0].(io.ReadWriter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPokemonFileReaderMockRecorder) OpenToWrite() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenToWrite", reflect.TypeOf((*MockPokemonFileReader)(nil).OpenToWrite))
}

func (m *MockPokemonFileReader) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockPokemonFileReaderMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPokemonFileReader)(nil).Close))
}
