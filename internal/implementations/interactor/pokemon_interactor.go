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

const EVEN = "even"
const ODD = "odd"

func NewPokemonUseCase(repo repository.PokemonRepository, client client.PokemonClient) usecase.PokemonUseCase {
	return &pokemonUseCase{repo, client}
}

func (useCase *pokemonUseCase) GetMultiple(idType string, items int, itemsPerWorker int) (result []model.Pokemon, err error) {
	if idType != EVEN && idType != ODD {
		err = custom_error.PokemonIdTypeError
		return
	}
	if items == 0 {
		err = custom_error.PokemonItemsError
		return
	}
	if itemsPerWorker == 0 {
		err = custom_error.PokemonItemsPerWorkerError
		return
	}

	result, err = useCase.Repository.FindAllByIdType(idType, items, itemsPerWorker)

	return
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

func (useCase *pokemonUseCase) GetFromApiAndSave(id int) (pokemon *model.Pokemon, err error) {
	oldPokemon, _ := useCase.GetById(id)
	if oldPokemon != nil {
		pokemon = oldPokemon
		err = custom_error.PokemonAlreadyExistError
		return
	}
	pokemon, err = useCase.Client.GetById(id)
	if err != nil {
		return
	}
	err = useCase.Repository.Save(pokemon)
	return
}
