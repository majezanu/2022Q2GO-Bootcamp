package repository

import (
	"context"
	"encoding/csv"
	"fmt"
	"majezanu/capstone/domain/custom_error"
	"majezanu/capstone/domain/model"
	"majezanu/capstone/internal/contracts/datastore"
	"majezanu/capstone/internal/contracts/repository"
	"strconv"
	"strings"
	"sync"
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
	reader, err := p.OpenerCloser.OpenToRead()
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
	reader, err := fileReader.OpenToRead()
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

const EVEN = "even"
const ODD = "odd"

func (p pokemonRepository) FindAllByIdType(idType string, items int, itemsPerWorker int) (result []model.Pokemon, err error) {
	reader, err := p.OpenerCloser.OpenToRead()
	if err != nil {
		return nil, custom_error.PokemonFileCantBeOpen
	}
	defer p.OpenerCloser.Close()
	csvReader := csv.NewReader(reader)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	src := make(chan model.Pokemon, items)
	out := make(chan model.Pokemon, items)

	// use a waitgroup to manage synchronization
	var wg sync.WaitGroup

	// declare the workers

	finder := func() {
		for {
			select {
			case <-ctx.Done(): // if the context is cancelled, quit.
				return
			default: // you must check for readable state of the channel.
				line, readingError := csvReader.Read()
				if readingError != nil {
					cancel()
					return
				}

				id, numError := strconv.Atoi(line[0])
				if numError != nil {
					return
				}
				if (idType == EVEN && id%2 == 0) || (idType == ODD && id%2 == 1) {
					pokemon := model.Pokemon{
						Id:   id,
						Name: line[1],
					}
					src <- pokemon
				}
			}
		}
	}

	collector := func() {
		for {
			select {
			case <-ctx.Done(): // if the context is cancelled, quit.
				return
			case pokemon, ok := <-src: // you must check for readable state of the channel.
				if !ok {
					return
				}
				if len(result) == items {
					cancel()
					return
				}
				out <- pokemon // do somethingg useful.
			}
		}
	}

	wg.Add(1)
	go func() {
		collector()
		wg.Done()
	}()

	go finder()

	// wait for worker group to finish and close out
	go func() {
		wg.Wait()  // wait for writers to quit.
		close(out) // when you close(out) it breaks the below loop.
	}()

	// drain the output
	for res := range out {
		if len(result) == items {
			cancel()
			break
		}
		result = append(result, res)
	}

	return
}

func (p pokemonRepository) Save(pokemon *model.Pokemon) (err error) {
	file, err := p.OpenerCloser.OpenToWrite()
	if err != nil {
		return
	}
	writer := csv.NewWriter(file)

	id := strconv.Itoa(pokemon.Id)
	newData := []string{id, pokemon.Name}

	err = writer.Write(newData)
	if err != nil {
		return
	}

	writer.Flush()
	err = writer.Error()
	if err != nil {
		return err
	}

	return p.OpenerCloser.Close()
}

func NewPokemonRepository(reader datastore.OpenerCloser) repository.PokemonRepository {

	return &pokemonRepository{
		reader,
	}
}
