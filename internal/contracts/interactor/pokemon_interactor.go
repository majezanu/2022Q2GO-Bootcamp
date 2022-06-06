package interactor

import (
	"majezanu/capstone/domain/model"
)

type PokemonUseCase interface {
	GetById(id int) (*model.Pokemon, error)
	GetMultiple(idType string, items int, itemsPerWorker int) ([]model.Pokemon, error)
	GetByName(name string) (*model.Pokemon, error)
	GetAll() ([]model.Pokemon, error)
	GetFromApiAndSave(id int) (*model.Pokemon, error)
}
