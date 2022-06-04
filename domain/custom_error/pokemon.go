package custom_error

import "errors"

var PokemonNotFoundError = errors.New("pokemon not found")

var BadPokemonFieldError = errors.New("bad pokemon field")

var PokemonFieldNotMappedError = errors.New("pokemon field is not mapped")

var PokemonIdFormatError = errors.New("pokemon id cannot be formatted")

var PokemonFileCantBeOpen = errors.New("pokemon file cannot be open")

var PokemonApiTimeoutError = errors.New("pokemon api timeout")

var PokemonSaveError = errors.New("pokemon can't be saved")

var PokemonAlreadyExistError = errors.New("pokemon already exist")
