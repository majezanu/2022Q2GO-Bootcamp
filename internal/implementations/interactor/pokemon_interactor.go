package interactor

import (
	"majezanu/capstone/domain/custom_error"
	"majezanu/capstone/domain/model"
	usecase "majezanu/capstone/internal/contracts/interactor"
	"majezanu/capstone/internal/contracts/repository"
)

type pokemonUseCase struct {
	Repository repository.PokemonRepository
}

func NewPokemonUseCase(repo repository.PokemonRepository) usecase.PokemonUseCase {
	return &pokemonUseCase{repo}
}

func processResult(input *model.Pokemon, errInput error) (pokemon *model.Pokemon, err error) {
	err = errInput
	pokemon = input
	if err == nil && pokemon == nil {
		err = custom_error.UnexpectedError
		return
	}

	if errInput != nil {
		pokemon = nil
	}

	return
}

func (useCase *pokemonUseCase) GetById(id int) (pokemon *model.Pokemon, err error) {
	pokemon, err = useCase.Repository.FindByField("id", id)
	pokemon, err = processResult(pokemon, err)
	return
}

func (useCase *pokemonUseCase) GetByName(name string) (pokemon *model.Pokemon, err error) {
	pokemon, err = useCase.Repository.FindByField("name", name)
	return processResult(pokemon, err)
}

func (useCase *pokemonUseCase) GetAll() (pokemonList []model.Pokemon, err error) {
	pokemonList, err = useCase.Repository.FindAll()

	if err != nil {
		pokemonList = nil
	}

	return
}
