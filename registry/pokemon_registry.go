package registry

import (
	"majezanu/capstone/internal/contracts/controller"
	"majezanu/capstone/internal/contracts/datastore"
	"majezanu/capstone/internal/contracts/interactor"
	"majezanu/capstone/internal/contracts/repository"
	controllerImpl "majezanu/capstone/internal/implementations/controller"
	datastoreImpl "majezanu/capstone/internal/implementations/datastore"
	interactorImpl "majezanu/capstone/internal/implementations/interactor"
	repositoryImpl "majezanu/capstone/internal/implementations/repository"
)

func (r *registry) NewPokemonController() controller.PokemonController {
	return controllerImpl.NewPokemonController(r.NewPokemonInteractor())
}

func (r *registry) NewPokemonInteractor() interactor.PokemonUseCase {
	return interactorImpl.NewPokemonUseCase(r.NewPokemonRepository())
}

func (r *registry) NewPokemonRepository() repository.PokemonRepository {
	return repositoryImpl.NewPokemonRepository(r.NewPokemonFileReader())
}

func (r *registry) NewPokemonFileReader() datastore.ReadWriteCloser {
	return datastoreImpl.NewPokemonFileReader(r.csvPath)
}
