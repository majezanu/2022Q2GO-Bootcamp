package client

import (
	"github.com/golang/mock/gomock"
	"net/http"
	"reflect"
)

type MockHttpClient struct {
	ctrl     *gomock.Controller
	recorder *MockHttpClientRecorder
}

type MockHttpClientRecorder struct {
	mock *MockHttpClient
}

func NewMockHttpClient(ctrl *gomock.Controller) *MockHttpClient {
	mock := &MockHttpClient{ctrl: ctrl}
	mock.recorder = &MockHttpClientRecorder{mock}
	return mock
}

func (m *MockHttpClient) EXPECT() *MockHttpClientRecorder {
	return m.recorder
}

func (m *MockHttpClient) Get(url string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", url)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockHttpClientRecorder) Get(url string) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHttpClient)(nil).Get), url)
}
