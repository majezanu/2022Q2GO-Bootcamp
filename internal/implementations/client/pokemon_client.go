package client

import (
	"encoding/json"
	"fmt"
	"majezanu/capstone/domain/custom_error"
	"majezanu/capstone/domain/model"
	"majezanu/capstone/internal/contracts/client"
	"net/http"
)

type pokemonClient struct {
	Client client.HttpClient
}

func BuildPath(id int) string {
	return fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", id)
}

func (p pokemonClient) GetById(id int) (pokemon *model.Pokemon, err error) {
	path := BuildPath(id)
	r, err := p.Client.Get(path)
	if err != nil {
		return
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		switch r.StatusCode {
		case http.StatusNotFound:
			err = custom_error.PokemonNotFoundError
		case http.StatusRequestTimeout:
			err = custom_error.PokemonApiTimeoutError
		default:
			err = custom_error.UnexpectedError
		}
		return
	}

	err = json.NewDecoder(r.Body).Decode(&pokemon)
	if err != nil {
		return
	}
	return
}

func NewPokemonClient(httpClient client.HttpClient) client.PokemonClient {
	return &pokemonClient{httpClient}
}
