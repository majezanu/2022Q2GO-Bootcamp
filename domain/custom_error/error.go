package custom_error

import (
	"errors"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error" example:"message"`
	Code  int    `json:"code" example:"400"`
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
