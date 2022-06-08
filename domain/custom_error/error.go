package custom_error

import (
	"errors"
	"fmt"
	"majezanu/capstone/domain/model"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error" example:"message"`
	Code  int    `json:"code" example:"400"`
}

type PokemonWithError struct {
	Error   string        `json:"error" example:"message"`
	Pokemon model.Pokemon `json:"pokemon"`
}

func NewPokemonWithError(pokemon *model.Pokemon, err error) PokemonWithError {
	return PokemonWithError{fmt.Sprint(err), *pokemon}
}

func NewErrorResponse(err error) ErrorResponse {
	httpStatusCode := http.StatusInternalServerError
	switch {
	case errors.Is(err, PokemonIdFormatError) || errors.Is(err, BadPokemonFieldError) || errors.Is(err, PokemonAlreadyExistError):
		httpStatusCode = http.StatusUnprocessableEntity
	case errors.Is(err, PokemonNotFoundError):
		httpStatusCode = http.StatusNotFound
	case errors.Is(err, PokemonFileCantBeOpen):
		httpStatusCode = http.StatusBadRequest
	default:
		httpStatusCode = http.StatusInternalServerError
	}
	return ErrorResponse{
		Error: fmt.Sprint(err),
		Code:  httpStatusCode,
	}
}
