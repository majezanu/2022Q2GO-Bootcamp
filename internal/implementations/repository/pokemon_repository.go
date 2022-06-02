package repository

import (
	"encoding/csv"
	"fmt"
	"majezanu/capstone/domain/custom_error"
	"majezanu/capstone/domain/model"
	"majezanu/capstone/internal/contracts/datastore"
	"majezanu/capstone/internal/contracts/repository"
	"os"
	"strconv"
	"strings"
)

type pokemonRepository struct {
	OpenerCloser datastore.OpenerCloser
}

const IdColumnName = "id"
const IdColumn = 0
const NameColumnName = "name"
const NameColumn = 1

func getValidFields() []string {
	return []string{IdColumnName, NameColumnName}
}

func isValidField(field string) bool {
	validFields := getValidFields()
	isValid := false
	for _, validField := range validFields {
		if strings.ToLower(validField) == strings.ToLower(field) {
			isValid = true
			break
		}
	}
	return isValid
}

func getColumnByField(field string) (*int, error) {
	if !isValidField(field) {
		return nil, custom_error.BadPokemonFieldError
	}
	fieldsAndColumns := map[string]int{IdColumnName: IdColumn, NameColumnName: NameColumn}
	if val, ok := fieldsAndColumns[field]; ok {
		if ok {
			return &val, nil
		}
	}
	return nil, custom_error.PokemonFieldNotMappedError
}

func (p pokemonRepository) FindByField(field string, value interface{}) (pokemon *model.Pokemon, err error) {
	reader, err := p.OpenerCloser.Open()
	if err != nil {
		return nil, custom_error.PokemonFileCantBeOpen
	}
	defer p.OpenerCloser.Close()
	csvReader := csv.NewReader(reader)
	found := false
	column, err := getColumnByField(field)
	if err != nil {
		return nil, err
	}
	value = fmt.Sprint(value)
	for {
		line, readingError := csvReader.Read()
		if readingError != nil {
			break
		}
		if value == line[*column] {
			found = true
			id, numError := strconv.Atoi(line[0])
			if numError != nil {
				err = custom_error.PokemonIdFormatError
				break
			}
			pokemon = &model.Pokemon{
				Id:   id,
				Name: line[1],
			}
		}
	}
	if !found {
		err = custom_error.PokemonNotFoundError
	}

	return
}

func (p pokemonRepository) FindAll() (pokemonList []model.Pokemon, err error) {
	fileReader := p.OpenerCloser
	reader, err := fileReader.Open()
	if err != nil {
		return nil, custom_error.PokemonFileCantBeOpen
	}
	defer fileReader.Close()
	csvLines, err := csv.NewReader(reader).ReadAll()
	if err != nil {
		return
	}
	for _, line := range csvLines {
		id, errNum := strconv.Atoi(line[0])
		if errNum != nil {
			err = custom_error.PokemonIdFormatError
			break
		}
		item := model.Pokemon{
			Id:   id,
			Name: line[1],
		}
		pokemonList = append(pokemonList, item)
	}
	return
}

func (p pokemonRepository) Save(pokemon *model.Pokemon) (err error) {
	pokemonExistingList, err := p.FindAll()
	if err != nil {
		return
	}
	pokemonExistingList = append(pokemonExistingList, *pokemon)

	file, err := os.Create("infrastructure/datastore/data.csv")
	writer := csv.NewWriter(file)
	var data [][]string
	for _, item := range pokemonExistingList {
		id := strconv.Itoa(item.Id)
		i := []string{id, item.Name}
		data = append(data, i)
	}
	writer.WriteAll(data)
	defer writer.Flush()
	return p.OpenerCloser.Close()
}

func NewPokemonRepository(reader datastore.OpenerCloser) repository.PokemonRepository {

	return &pokemonRepository{
		reader,
	}
}
