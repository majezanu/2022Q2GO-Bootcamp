package client

import (
	"github.com/golang/mock/gomock"
	"majezanu/capstone/domain/model"
	"reflect"
)

type MockPokemonClient struct {
	ctrl     *gomock.Controller
	recorder *MockPokemonClientRecorder
}

type MockPokemonClientRecorder struct {
	mock *MockPokemonClient
}

func NewMockPokemonClient(ctrl *gomock.Controller) *MockPokemonClient {
	mock := &MockPokemonClient{ctrl: ctrl}
	mock.recorder = &MockPokemonClientRecorder{mock}
	return mock
}

func (m *MockPokemonClient) EXPECT() *MockPokemonClientRecorder {
	return m.recorder
}

func (m *MockPokemonClient) GetById(id int) (*model.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*model.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPokemonClientRecorder) GetById(id int) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockPokemonClient)(nil).GetById), id)
}
