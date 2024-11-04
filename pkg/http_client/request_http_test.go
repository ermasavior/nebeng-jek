package http_client

import (
	"context"
	"encoding/json"
	http_utils "nebeng-jek/pkg/http/utils"
	"nebeng-jek/pkg/utils"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRequestHTTPAndParseResponse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	baseURL := "127.0.0.1:7171"

	transport := RestTransport{
		Url:    "http://" + baseURL,
		Method: http.MethodGet,
	}

	responseMock := http_utils.ClientResponse{
		Meta: http_utils.MetaResponse{
			Code:    0,
			Message: "success",
		},
	}
	httpRes, _ := json.Marshal(responseMock)

	t.Run("success - request add credit return success", func(t *testing.T) {
		handlerMock := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(httpRes)
		})
		server := utils.MockHTTPServer(t, baseURL, handlerMock)
		defer server.Close()

		res, err := RequestHTTPAndParseResponse(context.TODO(), HttpClient(), transport)
		assert.NoError(t, err)
		assert.Equal(t, responseMock, res)
	})

	t.Run("return error - error from server", func(t *testing.T) {
		handlerMock := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(httpRes)
		})
		server := utils.MockHTTPServer(t, baseURL, handlerMock)
		defer server.Close()

		res, err := RequestHTTPAndParseResponse(context.TODO(), HttpClient(), transport)
		assert.Error(t, err)
		assert.Equal(t, http_utils.ClientResponse{}, res)
	})

	t.Run("return error - connection refused", func(t *testing.T) {
		// no server running
		_, err := RequestHTTPAndParseResponse(context.TODO(), HttpClient(), transport)
		assert.Error(t, err)
	})
}
