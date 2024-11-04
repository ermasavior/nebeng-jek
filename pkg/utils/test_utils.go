package utils

import (
	"errors"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type FailingType struct{}

func (f FailingType) MarshalJSON() ([]byte, error) {
	return nil, errors.New("forced marshal failure")
}

func MockHTTPServer(t *testing.T, baseURL string, handlerMock http.Handler) *httptest.Server {
	l, err := net.Listen("tcp", baseURL)
	if err != nil {
		assert.Fail(t, "error create test server", err.Error())
	}
	server := httptest.NewUnstartedServer(handlerMock)

	server.Listener.Close()
	server.Listener = l
	server.Start()

	return server
}
