package repository

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"majezanu/capstone/domain/custom_error"
	"majezanu/capstone/domain/model"
	"majezanu/capstone/internal/implementations/datastore"
	"os"
	"strings"
	"testing"
)

func setup(t *testing.T) *datastore.MockPokemonFileReader {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	return datastore.NewMockPokemonFileReader(mockCtl)
}

func TestPokemonRepository_FindByField(t *testing.T) {
	t.Parallel()

	fileReader := setup(t)
	type test struct {
		name  string
		mock  func()
		field string
		value interface{}
		res   *model.Pokemon
		err   error
	}
	const mockedCsv = "1,Picachu\n2,Charmander"
	tests := []test{
		{
			name: "Pokemon can't read data",
			mock: func() {
				reader := strings.NewReader(mockedCsv)
				fileReader.EXPECT().Read().Return(reader, os.ErrNotExist)
				fileReader.EXPECT().Close().Return(nil)
			},
			field: "id",
			value: 3,
			res:   nil,
			err:   custom_error.PokemonFileCantBeRead,
		},
		{
			name: "Pokemon not found - Id",
			mock: func() {
				reader := strings.NewReader(mockedCsv)
				fileReader.EXPECT().Read().Return(reader, nil)
				fileReader.EXPECT().Close().Return(nil)
			},
			field: "id",
			value: 3,
			res:   nil,
			err:   custom_error.PokemonNotFoundError,
		},
		{
			name: "Pokemon not found - Name",
			mock: func() {
				reader := strings.NewReader(mockedCsv)
				fileReader.EXPECT().Read().Return(reader, nil)
				fileReader.EXPECT().Close().Return(nil)
			},
			field: "name",
			value: "Karmander",
			res:   nil,
			err:   custom_error.PokemonNotFoundError,
		},
		{
			name: "Pokemon not found - EOF",
			mock: func() {
				reader := strings.NewReader("")
				fileReader.EXPECT().Read().Return(reader, nil)
				fileReader.EXPECT().Close().Return(nil)
			},
			field: "id",
			value: 3,
			res:   nil,
			err:   custom_error.PokemonNotFoundError,
		},
		{
			name: "Pokemon bad field",
			mock: func() {
				reader := strings.NewReader(mockedCsv)
				fileReader.EXPECT().Read().Return(reader, nil)
				fileReader.EXPECT().Close().Return(nil)
			},
			field: "gender",
			value: "male",
			res:   nil,
			err:   custom_error.BadPokemonFieldError,
		},
		{
			name: "Pokemon exist - Id",
			mock: func() {
				reader := strings.NewReader(mockedCsv)
				fileReader.EXPECT().Read().Return(reader, nil)
				fileReader.EXPECT().Close().Return(nil)
			},
			field: "id",
			value: 1,
			res:   &model.Pokemon{Id: 1, Name: "Picachu"},
			err:   nil,
		},
		{
			name: "Pokemon exist - Name",
			mock: func() {
				reader := strings.NewReader(mockedCsv)
				fileReader.EXPECT().Read().Return(reader, nil)
				fileReader.EXPECT().Close().Return(nil)
			},
			field: "name",
			value: "Charmander",
			res:   &model.Pokemon{Id: 2, Name: "Charmander"},
			err:   nil,
		},
		{
			name: "Pokemon exist - Bad formatted id",
			mock: func() {
				reader := strings.NewReader("b,Picachu")
				fileReader.EXPECT().Read().Return(reader, nil)
				fileReader.EXPECT().Close().Return(nil)
			},
			field: "name",
			value: "Picachu",
			res:   nil,
			err:   custom_error.PokemonIdFormatError,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			res, err := NewPokemonRepository(fileReader).FindByField(tc.field, tc.value)
			expectedResult := tc.res
			expectedError := tc.err
			require.Equal(t, expectedResult, res)
			require.ErrorIs(t, expectedError, err)
		})
	}
}

func TestNewPokemonRepository_FindAll(t *testing.T) {
	type test struct {
		name string
		mock func()
		res  []model.Pokemon
		err  error
	}
	fileReader := setup(t)
	const mockedCsv = "1,Picachu\n2,Charmander"
	tests := []test{
		{
			name: "Pokemon can't read data",
			mock: func() {
				reader := strings.NewReader(mockedCsv)
				fileReader.EXPECT().Read().Return(reader, os.ErrNotExist)
				fileReader.EXPECT().Close().Return(nil)
			},
			res: nil,
			err: custom_error.PokemonFileCantBeRead,
		},
		{
			name: "Pokemon list",
			mock: func() {
				reader := strings.NewReader(mockedCsv)
				fileReader.EXPECT().Read().Return(reader, nil)
				fileReader.EXPECT().Close().Return(nil)
			},
			res: []model.Pokemon{
				{1, "Picachu"},
				{2, "Charmander"},
			},
			err: nil,
		},
		{
			name: "Pokemon empty result",
			mock: func() {
				reader := strings.NewReader("")
				fileReader.EXPECT().Read().Return(reader, nil)
				fileReader.EXPECT().Close().Return(nil)
			},
			res: []model.Pokemon(nil),
			err: nil,
		},
		{
			name: "Pokemon bad formatted id",
			mock: func() {
				reader := strings.NewReader("b,Picachu\n1,Charmander")
				fileReader.EXPECT().Read().Return(reader, nil)
				fileReader.EXPECT().Close().Return(nil)
			},
			res: []model.Pokemon(nil),
			err: custom_error.PokemonIdFormatError,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			res, err := NewPokemonRepository(fileReader).FindAll()
			expectedResult := tc.res
			expectedError := tc.err
			require.Equal(t, expectedResult, res)
			require.ErrorIs(t, expectedError, err)
		})
	}
}
