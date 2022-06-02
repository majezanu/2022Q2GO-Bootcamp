package custom_error

import "errors"

var PokemonNotFoundError = errors.New("pokemon not found")

var BadPokemonFieldError = errors.New("bad pokemon field")

var PokemonFieldNotMappedError = errors.New("pokemon field is not mapped")

var PokemonIdFormatError = errors.New("pokemon id cannot be formatted")

var PokemonFileCantBeRead = errors.New("pokemon file cannot be read")
