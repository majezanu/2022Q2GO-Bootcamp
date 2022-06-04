package repository

import (
	"majezanu/capstone/domain/model"
)

type PokemonRepository interface {
	FindByField(field string, value interface{}) (*model.Pokemon, error)
	FindAll() ([]model.Pokemon, error)
	Save(pokemon *model.Pokemon) error
}
