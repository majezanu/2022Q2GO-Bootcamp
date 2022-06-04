package controller

import (
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

func responseError(c controller.Context, err error) error {
	errorResponse := custom_error.NewErrorResponse(err)
	return c.JSON(errorResponse.Code, errorResponse)
}

func (p pokemonController) FetchByIdAndSave(c controller.Context) error {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)

	if err != nil {
		return responseError(c, custom_error.PokemonIdFormatError)
	}

	err = p.pokemonInteractor.GetFromApiAndSave(id)
	if err != nil {
		return responseError(c, err)
	}

	return c.JSON(http.StatusCreated, "ok")
}

func (p pokemonController) GetById(c controller.Context) error {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)

	if err != nil {
		return responseError(c, custom_error.PokemonIdFormatError)
	}

	var u *model.Pokemon
	u, err = p.pokemonInteractor.GetById(id)
	if err != nil {
		return responseError(c, err)
	}

	return c.JSON(http.StatusOK, u)
}

func (p pokemonController) GetByName(c controller.Context) error {
	paramName := c.Param("name")

	var u *model.Pokemon
	u, err := p.pokemonInteractor.GetByName(paramName)
	if err != nil {
		return responseError(c, err)
	}

	return c.JSON(http.StatusOK, u)
}

func (p *pokemonController) GetAll(c controller.Context) error {
	var u []model.Pokemon

	u, err := p.pokemonInteractor.GetAll()
	if err != nil {
		return responseError(c, err)
	}
	return c.JSON(http.StatusOK, u)
}

func NewPokemonController(useCase interactor.PokemonUseCase) controller.PokemonController {
	return &pokemonController{useCase}
}
