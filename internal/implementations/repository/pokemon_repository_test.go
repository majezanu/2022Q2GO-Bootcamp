package repository

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"majezanu/capstone/domain/custom_error"
	"majezanu/capstone/domain/model"
	"majezanu/capstone/internal/implementations/datastore"
	"os"
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
				var buff *bytes.Buffer
				buff = bytes.NewBufferString(mockedCsv)
				fileReader.EXPECT().OpenToRead().Return(buff, os.ErrNotExist)
				fileReader.EXPECT().Close().Return(nil)
			},
			field: "id",
			value: 3,
			res:   nil,
			err:   custom_error.PokemonFileCantBeOpen,
		},
		{
			name: "Pokemon not found - Id",
			mock: func() {
				var buff *bytes.Buffer
				buff = bytes.NewBufferString(mockedCsv)
				fileReader.EXPECT().OpenToRead().Return(buff, nil)
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
				var buff *bytes.Buffer
				buff = bytes.NewBufferString(mockedCsv)
				fileReader.EXPECT().OpenToRead().Return(buff, nil)
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
				var buff *bytes.Buffer
				buff = bytes.NewBufferString("")
				fileReader.EXPECT().OpenToRead().Return(buff, nil)
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
				var buff *bytes.Buffer
				buff = bytes.NewBufferString(mockedCsv)
				fileReader.EXPECT().OpenToRead().Return(buff, nil)
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
				var buff *bytes.Buffer
				buff = bytes.NewBufferString(mockedCsv)
				fileReader.EXPECT().OpenToRead().Return(buff, nil)
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
				var buff *bytes.Buffer
				buff = bytes.NewBufferString(mockedCsv)
				fileReader.EXPECT().OpenToRead().Return(buff, nil)
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
				var buff *bytes.Buffer
				buff = bytes.NewBufferString("b,Picachu")
				fileReader.EXPECT().OpenToRead().Return(buff, nil)
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
				var buff *bytes.Buffer
				buff = bytes.NewBufferString(mockedCsv)
				fileReader.EXPECT().OpenToRead().Return(buff, os.ErrNotExist)
				fileReader.EXPECT().Close().Return(nil)
			},
			res: nil,
			err: custom_error.PokemonFileCantBeOpen,
		},
		{
			name: "Pokemon list",
			mock: func() {
				var buff *bytes.Buffer
				buff = bytes.NewBufferString(mockedCsv)
				fileReader.EXPECT().OpenToRead().Return(buff, nil)
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
				var buff *bytes.Buffer
				buff = bytes.NewBufferString("")
				fileReader.EXPECT().OpenToRead().Return(buff, nil)
				fileReader.EXPECT().Close().Return(nil)
			},
			res: []model.Pokemon(nil),
			err: nil,
		},
		{
			name: "Pokemon bad formatted id",
			mock: func() {
				var buff *bytes.Buffer
				buff = bytes.NewBufferString("b,Picachu\n1,Charmander")
				fileReader.EXPECT().OpenToRead().Return(buff, nil)
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

func TestNewPokemonRepository_Save(t *testing.T) {
	type test struct {
		name     string
		mock     func(buff *bytes.Buffer)
		finalCsv string
		input    *model.Pokemon
		err      error
	}
	fileReader := setup(t)
	t.Parallel()
	tests := []test{
		{
			name: "Pokemon save success",
			mock: func(buff *bytes.Buffer) {
				fileReader.EXPECT().OpenToWrite().Return(buff, nil)
				fileReader.EXPECT().Close().Return(nil)
			},
			finalCsv: "1,Picachu\n2,Charmander\n3,Charizard\n",
			input:    &model.Pokemon{Id: 3, Name: "Charizard"},
			err:      nil,
		},
		{
			name: "Pokemon alrady exists",
			mock: func(buff *bytes.Buffer) {
				fileReader.EXPECT().OpenToWrite().Return(buff, nil)
				fileReader.EXPECT().Close().Return(nil)
			},
			finalCsv: "1,Picachu\n2,Charmander\n1,Picachu\n",
			input:    &model.Pokemon{Id: 1, Name: "Picachu"},
			err:      nil,
		},
		{
			name: "Pokemon save error",
			mock: func(buff *bytes.Buffer) {
				fileReader.EXPECT().OpenToWrite().Return(nil, os.ErrNotExist)
				fileReader.EXPECT().Close().Return(nil)
			},
			finalCsv: "1,Picachu\n2,Charmander\n",
			input:    &model.Pokemon{Id: 3, Name: "Charizard"},
			err:      custom_error.PokemonFileCantBeOpen,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			var buff *bytes.Buffer
			buff = bytes.NewBufferString("1,Picachu\n2,Charmander\n")
			tc.mock(buff)
			err := NewPokemonRepository(fileReader).Save(tc.input)
			expectedError := tc.err
			finalString := buff.String()
			require.Equal(t, tc.finalCsv, finalString)
			require.ErrorIs(t, expectedError, err)
		})
	}
}
