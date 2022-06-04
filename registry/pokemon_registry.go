package registry

import (
	"majezanu/capstone/internal/contracts/client"
	"majezanu/capstone/internal/contracts/controller"
	"majezanu/capstone/internal/contracts/datastore"
	"majezanu/capstone/internal/contracts/interactor"
	"majezanu/capstone/internal/contracts/repository"
	clientImpl "majezanu/capstone/internal/implementations/client"
	controllerImpl "majezanu/capstone/internal/implementations/controller"
	datastoreImpl "majezanu/capstone/internal/implementations/datastore"
	interactorImpl "majezanu/capstone/internal/implementations/interactor"
	repositoryImpl "majezanu/capstone/internal/implementations/repository"
	"net/http"
)

func (r *registry) NewPokemonController() controller.PokemonController {
	return controllerImpl.NewPokemonController(r.NewPokemonInteractor())
}

func (r *registry) NewPokemonInteractor() interactor.PokemonUseCase {
	return interactorImpl.NewPokemonUseCase(r.NewPokemonRepository(), r.NewPokemonClient())
}

func (r *registry) NewPokemonRepository() repository.PokemonRepository {
	return repositoryImpl.NewPokemonRepository(r.NewPokemonFileReader())
}

func (r *registry) NewPokemonClient() client.PokemonClient {
	return clientImpl.NewPokemonClient(&http.Client{})
}

func (r *registry) NewPokemonFileReader() datastore.OpenerCloser {
	return datastoreImpl.NewPokemonFileReader(r.csvPath)
}
