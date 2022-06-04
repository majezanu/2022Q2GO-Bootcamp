package client

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"majezanu/capstone/domain/custom_error"
	"majezanu/capstone/domain/model"
	"net/http"
	"testing"
)

func setup(t *testing.T) *MockHttpClient {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	return NewMockHttpClient(mockCtl)
}

func TestPokemonClient_GetById(t *testing.T) {
	type test struct {
		name         string
		mock         func(apiResponse http.Response)
		id           int
		statusCode   int
		bodyResponse string
		res          *model.Pokemon
		err          error
	}

	t.Parallel()

	httpClient := setup(t)
	// build response JSON
	tests := []test{
		{
			name: "Pokemon api handler timeout",
			mock: func(apiResponse http.Response) {
				httpClient.EXPECT().Get(BuildPath(1)).Return(&apiResponse, http.ErrHandlerTimeout)
			},
			statusCode:   http.StatusRequestTimeout,
			bodyResponse: ``,
			id:           1,
			res:          nil,
			err:          http.ErrHandlerTimeout,
		},
		{
			name: "Pokemon api request timeout",
			mock: func(apiResponse http.Response) {
				httpClient.EXPECT().Get(BuildPath(1)).Return(&apiResponse, nil)
			},
			statusCode:   http.StatusRequestTimeout,
			bodyResponse: ``,
			id:           1,
			res:          nil,
			err:          custom_error.PokemonApiTimeoutError,
		},
		{
			name: "Pokemon not found",
			mock: func(apiResponse http.Response) {
				httpClient.EXPECT().Get(BuildPath(1)).Return(&apiResponse, nil)
			},
			statusCode:   http.StatusNotFound,
			bodyResponse: `{"error":"Pokemon not found""}`,
			id:           1,
			res:          nil,
			err:          custom_error.PokemonNotFoundError,
		},
		{
			name: "Pokemon not expected error",
			mock: func(apiResponse http.Response) {
				httpClient.EXPECT().Get(BuildPath(1)).Return(&apiResponse, nil)
			},
			statusCode:   http.StatusInternalServerError,
			bodyResponse: ``,
			id:           1,
			res:          nil,
			err:          custom_error.UnexpectedError,
		},
		{
			name: "Pokemon exist",
			mock: func(apiResponse http.Response) {
				httpClient.EXPECT().Get(BuildPath(1)).Return(&apiResponse, nil)
			},
			statusCode:   http.StatusOK,
			bodyResponse: `{"id":1, "name": "Pikachu"}`,
			id:           1,
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

			r := ioutil.NopCloser(bytes.NewReader([]byte(tc.bodyResponse)))
			response := http.Response{
				StatusCode: tc.statusCode,
				Body:       r,
			}

			tc.mock(response)
			res, err := NewPokemonClient(httpClient).GetById(tc.id)

			require.Equal(t, tc.res, res)
			require.ErrorIs(t, err, tc.err)
		})
	}

}
