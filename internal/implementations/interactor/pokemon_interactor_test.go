package interactor

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"majezanu/capstone/domain/custom_error"
	"majezanu/capstone/domain/model"
	"majezanu/capstone/internal/implementations/repository"
	"testing"
)

func setup(t *testing.T) *repository.MockPokemonRepository {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	return repository.NewMockPokemonRepository(mockCtl)
}

var fakePokemonData = []model.Pokemon{
	{
		Id:   1,
		Name: "Picachu",
	},
}

// test function
func TestPokemonUseCase_GetById(t *testing.T) {

	type test struct {
		name string
		mock func()
		res  *model.Pokemon
		err  error
	}

	t.Parallel()

	repo := setup(t)

	pokemon := fakePokemonData[0]

	tests := []test{
		{
			name: "Not found result",
			mock: func() {
				repo.EXPECT().FindByField("id", 1).Return(nil, custom_error.PokemonNotFoundError)
			},
			res: nil,
			err: custom_error.PokemonNotFoundError,
		},
		{
			name: "Found",
			mock: func() {
				repo.EXPECT().FindByField("id", 1).Return(&pokemon, nil)
			},
			res: &pokemon,
			err: nil,
		},
		{
			name: "Unexpected error",
			mock: func() {
				repo.EXPECT().FindByField("id", 1).Return(nil, nil)
			},
			res: nil,
			err: custom_error.UnexpectedError,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()
			res, err := NewPokemonUseCase(repo).GetById(1)

			require.Equal(t, tc.res, res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestPokemonUseCase_GetByName(t *testing.T) {

	type test struct {
		name string
		mock func()
		res  *model.Pokemon
		err  error
	}

	t.Parallel()

	repo := setup(t)

	pokemon := fakePokemonData[0]

	tests := []test{
		{
			name: "Not found result",
			mock: func() {
				repo.EXPECT().FindByField("name", "Picachu").Return(nil, custom_error.PokemonNotFoundError)
			},
			res: nil,
			err: custom_error.PokemonNotFoundError,
		},
		{
			name: "Found",
			mock: func() {
				repo.EXPECT().FindByField("name", "Picachu").Return(&pokemon, nil)
			},
			res: &pokemon,
			err: nil,
		},
		{
			name: "Unexpected error",
			mock: func() {
				repo.EXPECT().FindByField("name", "Picachu").Return(nil, nil)
			},
			res: nil,
			err: custom_error.UnexpectedError,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()
			res, err := NewPokemonUseCase(repo).GetByName("Picachu")

			require.Equal(t, tc.res, res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestPokemonUseCase_GetAll(t *testing.T) {
	t.Parallel()

	repo := setup(t)

	type test struct {
		name string
		mock func()
		res  []model.Pokemon
		err  error
	}

	tests := []test{
		{
			name: "Empty result",
			mock: func() {
				repo.EXPECT().FindAll().Return(nil, nil)
			},
			res: []model.Pokemon(nil),
			err: nil,
		},
		{
			name: "Empty result",
			mock: func() {
				repo.EXPECT().FindAll().Return(nil, nil)
			},
			res: []model.Pokemon(nil),
			err: nil,
		},
		{
			name: "Correct result",
			mock: func() {
				repo.EXPECT().FindAll().Return(fakePokemonData, nil)
			},
			res: fakePokemonData,
			err: nil,
		},
		{
			name: "Cant open file",
			mock: func() {
				repo.EXPECT().FindAll().Return(nil, custom_error.PokemonFileCantBeRead)
			},
			res: nil,
			err: custom_error.PokemonFileCantBeRead,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()
			res, err := NewPokemonUseCase(repo).GetAll()
			expectedResult := tc.res
			expectedError := tc.err
			require.Equal(t, expectedResult, res)
			require.ErrorIs(t, expectedError, err)
		})
	}
}
