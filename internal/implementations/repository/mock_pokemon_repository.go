package repository

import (
	"github.com/golang/mock/gomock"
	"majezanu/capstone/domain/model"
	"reflect"
)

type MockPokemonRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPokemonRepositoryMockRecorder
}

func (m *MockPokemonRepository) FindAllByIdType(idType string, items int, itemsPerWorker int) ([]model.Pokemon, error) {
	//TODO implement me
	panic("implement me")
}

type MockPokemonRepositoryMockRecorder struct {
	mock *MockPokemonRepository
}

func NewMockPokemonRepository(ctrl *gomock.Controller) *MockPokemonRepository {
	mock := &MockPokemonRepository{ctrl: ctrl}
	mock.recorder = &MockPokemonRepositoryMockRecorder{mock}
	return mock
}

func (m *MockPokemonRepository) EXPECT() *MockPokemonRepositoryMockRecorder {
	return m.recorder
}

func (m *MockPokemonRepository) FindByField(field string, value interface{}) (*model.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByField", field, value)
	ret0, _ := ret[0].(*model.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPokemonRepositoryMockRecorder) FindByField(field string, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByField", reflect.TypeOf((*MockPokemonRepository)(nil).FindByField), field, value)
}

func (m *MockPokemonRepository) FindAll() ([]model.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]model.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPokemonRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockPokemonRepository)(nil).FindAll))
}

func (m *MockPokemonRepository) Save(pokemon *model.Pokemon) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", pokemon)
	ret1, _ := ret[0].(error)
	return ret1
}

func (mr *MockPokemonRepositoryMockRecorder) Save(pokemon *model.Pokemon) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockPokemonRepository)(nil).Save), pokemon)
}
