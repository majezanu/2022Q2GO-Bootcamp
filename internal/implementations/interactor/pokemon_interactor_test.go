package interactor

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"majezanu/capstone/domain/custom_error"
	"majezanu/capstone/domain/model"
	externalClient "majezanu/capstone/internal/implementations/client"
	"majezanu/capstone/internal/implementations/repository"
	"testing"
)

func setup(t *testing.T) (*repository.MockPokemonRepository, *externalClient.MockPokemonClient) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	return repository.NewMockPokemonRepository(mockCtl), externalClient.NewMockPokemonClient(mockCtl)
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

	repo, client := setup(t)

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
			res, err := NewPokemonUseCase(repo, client).GetById(1)

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

	repo, client := setup(t)

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
			res, err := NewPokemonUseCase(repo, client).GetByName("Picachu")

			require.Equal(t, tc.res, res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestPokemonUseCase_GetAll(t *testing.T) {
	t.Parallel()

	repo, client := setup(t)

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
				repo.EXPECT().FindAll().Return(nil, custom_error.PokemonFileCantBeOpen)
			},
			res: nil,
			err: custom_error.PokemonFileCantBeOpen,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()
			res, err := NewPokemonUseCase(repo, client).GetAll()
			expectedResult := tc.res
			expectedError := tc.err
			require.Equal(t, expectedResult, res)
			require.ErrorIs(t, expectedError, err)
		})
	}
}

func TestPokemonUseCase_GetFromApiAndSave(t *testing.T) {
	t.Parallel()

	repo, client := setup(t)

	type test struct {
		name string
		mock func()
		res  *model.Pokemon
		err  error
	}

	tests := []test{
		{
			name: "Pokemon not found",
			mock: func() {
				repo.EXPECT().FindByField("id", 1).Return(nil, nil)
				client.EXPECT().GetById(1).Return(nil, custom_error.PokemonNotFoundError)
				repo.EXPECT().Save(nil).Times(1)
			},
			res: nil,
			err: custom_error.PokemonNotFoundError,
		},
		{
			name: "Pokemon save error",
			mock: func() {
				pokemon := model.Pokemon{
					Id:   1,
					Name: "Pikachu",
				}
				repo.EXPECT().FindByField("id", 1).Return(nil, nil)
				client.EXPECT().GetById(1).Return(&pokemon, nil)
				repo.EXPECT().Save(&pokemon).Times(1).Return(custom_error.PokemonSaveError)
			},
			res: &model.Pokemon{
				Id:   1,
				Name: "Pikachu",
			},
			err: custom_error.PokemonSaveError,
		},
		{
			name: "Pokemon already exist",
			mock: func() {
				pokemon := model.Pokemon{
					Id:   1,
					Name: "Pikachu",
				}
				repo.EXPECT().FindByField("id", 1).Return(&pokemon, nil)
				client.EXPECT().GetById(1).Times(0).Return(&pokemon, nil)
				repo.EXPECT().Save(&pokemon).Times(0).Return(custom_error.PokemonSaveError)
			},
			res: &model.Pokemon{
				Id:   1,
				Name: "Pikachu",
			},
			err: custom_error.PokemonAlreadyExistError,
		},
		{
			name: "Pokemon save success",
			mock: func() {
				pokemon := model.Pokemon{
					Id:   1,
					Name: "Pikachu",
				}
				repo.EXPECT().FindByField("id", 1).Return(nil, nil)
				client.EXPECT().GetById(1).Return(&pokemon, nil)
				repo.EXPECT().Save(&pokemon).Times(1).Return(nil)
			},
			res: &model.Pokemon{
				Id:   1,
				Name: "Pikachu",
			},
			err: nil,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()
			pokemon, err := NewPokemonUseCase(repo, client).GetFromApiAndSave(1)
			expectedError := tc.err
			require.Equal(t, tc.res, pokemon)
			require.ErrorIs(t, expectedError, err)
		})
	}
}
