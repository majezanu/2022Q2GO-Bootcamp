package controller

type PokemonController interface {
	GetById(c Context) error
	GetByName(c Context) error
	GetAll(c Context) error
}
