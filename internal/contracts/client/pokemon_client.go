package client

import "majezanu/capstone/domain/model"

type PokemonClient interface {
	GetById(id int) (*model.Pokemon, error)
}
