package interactor

import (
	"majezanu/capstone/domain/custom_error"
	"majezanu/capstone/domain/model"
	"majezanu/capstone/internal/contracts/client"
	usecase "majezanu/capstone/internal/contracts/interactor"
	"majezanu/capstone/internal/contracts/repository"
)

type pokemonUseCase struct {
	Repository repository.PokemonRepository
	Client     client.PokemonClient
}

func NewPokemonUseCase(repo repository.PokemonRepository, client client.PokemonClient) usecase.PokemonUseCase {
	return &pokemonUseCase{repo, client}
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

func (useCase *pokemonUseCase) GetFromApiAndSave(id int) (err error) {
	oldPokemon, err := useCase.GetById(id)
	if oldPokemon != nil {
		return custom_error.PokemonAlreadyExistError
	}
	pokemon, err := useCase.Client.GetById(id)
	if err != nil {
		return
	}
	return useCase.Repository.Save(pokemon)
}
