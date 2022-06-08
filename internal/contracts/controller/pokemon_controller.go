package controller

type PokemonController interface {
	GetById(c Context) error
	GetMultiple(c Context) error
	GetByName(c Context) error
	GetAll(c Context) error
	FetchByIdAndSave(c Context) error
}
