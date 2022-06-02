package controller

import (
	"errors"
	"fmt"
	"majezanu/capstone/domain/custom_error"
	"majezanu/capstone/domain/model"
	"majezanu/capstone/internal/contracts/controller"
	"majezanu/capstone/internal/contracts/interactor"
	"net/http"
	"strconv"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonUseCase
}

func (p pokemonController) GetById(c controller.Context) error {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)

	if err != nil {
		errorResponse := custom_error.ErrorResponse{Error: fmt.Sprint(err), Code: http.StatusUnprocessableEntity}
		return c.JSON(errorResponse.Code, errorResponse)
	}

	var u *model.Pokemon
	u, err = p.pokemonInteractor.GetById(id)
	if err != nil {
		httpStatusCode := http.StatusInternalServerError
		switch {
		case errors.Is(err, custom_error.PokemonIdFormatError) || errors.Is(err, custom_error.BadPokemonFieldError):
			httpStatusCode = http.StatusUnprocessableEntity
		case errors.Is(err, custom_error.PokemonNotFoundError):
			httpStatusCode = http.StatusNotFound
		default:
			httpStatusCode = http.StatusInternalServerError
		}
		errorResponse := custom_error.ErrorResponse{Error: fmt.Sprint(err), Code: httpStatusCode}
		return c.JSON(errorResponse.Code, errorResponse)
	}

	return c.JSON(http.StatusOK, u)
}

func (p pokemonController) GetByName(c controller.Context) error {
	//TODO implement me
	panic("implement me")
}

func (p *pokemonController) GetAll(c controller.Context) error {
	var u []model.Pokemon

	u, err := p.pokemonInteractor.GetAll()
	if err != nil {
		httpStatusCode := http.StatusInternalServerError
		switch {
		case errors.Is(err, custom_error.PokemonIdFormatError):
			httpStatusCode = http.StatusUnprocessableEntity
		default:
			httpStatusCode = http.StatusInternalServerError
		}
		errorResponse := custom_error.ErrorResponse{Error: fmt.Sprint(err), Code: httpStatusCode}
		return c.JSON(errorResponse.Code, errorResponse)
	}
	return c.JSON(http.StatusOK, u)
}

func NewPokemonController(useCase interactor.PokemonUseCase) controller.PokemonController {
	return &pokemonController{useCase}
}
