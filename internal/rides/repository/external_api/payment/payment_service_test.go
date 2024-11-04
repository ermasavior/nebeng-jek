package payment

import (
	"context"
	"encoding/json"
	"nebeng-jek/internal/rides/model"
	"nebeng-jek/pkg/configs"
	http_utils "nebeng-jek/pkg/http/utils"
	"nebeng-jek/pkg/http_client"
	"nebeng-jek/pkg/utils"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPaymentRepository_DeductCredit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	baseURL := "127.0.0.1:7171"

	mockConfig := &configs.Config{
		PaymentServiceURL:    "http://" + baseURL,
		PaymentServiceAPIKey: "mock-api-key",
	}
	serviceMock := NewPaymentRepository(mockConfig, http_client.HttpClient())

	param := model.DeductCreditRequest{
		MSISDN: "08123456",
		Value:  20000,
	}

	responseMock := http_utils.ClientResponse{
		Meta: http_utils.MetaResponse{
			Code:    0,
			Message: "success",
		},
	}

	failedResponseMock := http_utils.ClientResponse{
		Meta: http_utils.MetaResponse{
			Code:    -1,
			Message: "failed",
		},
	}

	successRes, _ := json.Marshal(responseMock)
	failedJson, _ := json.Marshal(failedResponseMock)

	t.Run("success - request deduct credit return success", func(t *testing.T) {
		handlerMock := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(successRes)
		})
		server := utils.MockHTTPServer(t, baseURL, handlerMock)
		defer server.Close()

		err := serviceMock.DeductCredit(context.TODO(), param)
		assert.NoError(t, err)
	})

	t.Run("return error - error from server", func(t *testing.T) {
		handlerMock := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(failedJson)
		})
		server := utils.MockHTTPServer(t, baseURL, handlerMock)
		defer server.Close()

		err := serviceMock.DeductCredit(context.TODO(), param)
		assert.Error(t, err)
	})

	t.Run("return error - server connection refused", func(t *testing.T) {
		// no server running
		err := serviceMock.DeductCredit(context.TODO(), param)
		assert.Error(t, err)
	})
}

func TestPaymentRepository_AddCredit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	baseURL := "127.0.0.1:7171"

	mockConfig := &configs.Config{
		PaymentServiceURL:    "http://" + baseURL,
		PaymentServiceAPIKey: "mock-api-key",
	}
	serviceMock := NewPaymentRepository(mockConfig, http_client.HttpClient())

	param := model.AddCreditRequest{
		MSISDN: "08123456",
		Value:  20000,
	}

	responseMock := http_utils.ClientResponse{
		Meta: http_utils.MetaResponse{
			Code:    0,
			Message: "success",
		},
	}

	failedResponseMock := http_utils.ClientResponse{
		Meta: http_utils.MetaResponse{
			Code:    -1,
			Message: "failed",
		},
	}

	successRes, _ := json.Marshal(responseMock)
	failedJson, _ := json.Marshal(failedResponseMock)

	t.Run("success - request add credit return success", func(t *testing.T) {
		handlerMock := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(successRes)
		})
		server := utils.MockHTTPServer(t, baseURL, handlerMock)
		defer server.Close()

		err := serviceMock.AddCredit(context.TODO(), param)
		assert.NoError(t, err)
	})

	t.Run("return error - error from server", func(t *testing.T) {
		handlerMock := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(failedJson)
		})
		server := utils.MockHTTPServer(t, baseURL, handlerMock)
		defer server.Close()

		err := serviceMock.AddCredit(context.TODO(), param)
		assert.Error(t, err)
	})

	t.Run("return error - server connection refused", func(t *testing.T) {
		// no server running
		err := serviceMock.AddCredit(context.TODO(), param)
		assert.Error(t, err)
	})
}
