package location

import (
	"context"
	"encoding/json"
	"nebeng-jek/internal/rides/model"
	"nebeng-jek/pkg/configs"
	http_utils "nebeng-jek/pkg/http/utils"
	"nebeng-jek/pkg/http_client"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func mockHTTPServer(t *testing.T, baseURL string, handlerMock http.Handler) *httptest.Server {
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

func TestLocationRepository_AddAvailableDriver(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	baseURL := "127.0.0.1:7171"

	mockConfig := &configs.Config{
		LocationServiceURL:    "http://" + baseURL,
		LocationServiceAPIKey: "mock-api-key",
	}
	serviceMock := NewLocationRepository(mockConfig, http_client.HttpClient())

	driverID := int64(1111)
	location := model.Coordinate{
		Longitude: 1, Latitude: 2,
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
		server := mockHTTPServer(t, baseURL, handlerMock)
		defer server.Close()

		err := serviceMock.AddAvailableDriver(context.TODO(), driverID, location)
		assert.NoError(t, err)
	})

	t.Run("return error - error from server", func(t *testing.T) {
		handlerMock := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(failedJson)
		})
		server := mockHTTPServer(t, baseURL, handlerMock)
		defer server.Close()

		err := serviceMock.AddAvailableDriver(context.TODO(), driverID, location)
		assert.Error(t, err)
	})

	t.Run("return error - connection refused", func(t *testing.T) {
		// no server running
		err := serviceMock.AddAvailableDriver(context.TODO(), driverID, location)
		assert.Error(t, err)
	})
}

func TestLocationRepository_RemoveAvailableDriver(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	baseURL := "127.0.0.1:7171"

	mockConfig := &configs.Config{
		LocationServiceURL:    "http://" + baseURL,
		LocationServiceAPIKey: "mock-api-key",
	}
	serviceMock := NewLocationRepository(mockConfig, http_client.HttpClient())

	driverID := int64(2222)

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
		server := mockHTTPServer(t, baseURL, handlerMock)
		defer server.Close()

		err := serviceMock.RemoveAvailableDriver(context.TODO(), driverID)
		assert.NoError(t, err)
	})

	t.Run("return error - error from server", func(t *testing.T) {
		handlerMock := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(failedJson)
		})
		server := mockHTTPServer(t, baseURL, handlerMock)
		defer server.Close()

		err := serviceMock.RemoveAvailableDriver(context.TODO(), driverID)
		assert.Error(t, err)
	})

	t.Run("return error - connection refused", func(t *testing.T) {
		// no server running
		err := serviceMock.RemoveAvailableDriver(context.TODO(), driverID)
		assert.Error(t, err)
	})
}

func TestLocationRepository_GetNearestAvailableDrivers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	baseURL := "127.0.0.1:7171"

	mockConfig := &configs.Config{
		LocationServiceURL:    "http://" + baseURL,
		LocationServiceAPIKey: "mock-api-key",
	}
	serviceMock := NewLocationRepository(mockConfig, http_client.HttpClient())

	location := model.Coordinate{Longitude: 1, Latitude: 2}
	driverIDs := []int64{2222, 4444, 7777}
	data, _ := json.Marshal(model.GetNearestAvailableDriversResponse{
		DriverIDs: driverIDs,
	})

	responseMock := http_utils.ClientResponse{
		Meta: http_utils.MetaResponse{
			Code:    0,
			Message: "success",
		},
		Data: data,
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
		server := mockHTTPServer(t, baseURL, handlerMock)
		defer server.Close()

		actual, err := serviceMock.GetNearestAvailableDrivers(context.TODO(), location)
		assert.NoError(t, err)
		assert.Equal(t, driverIDs, actual)
	})

	t.Run("return error - error from server", func(t *testing.T) {
		handlerMock := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadGateway)
			_, _ = w.Write(failedJson)
		})
		server := mockHTTPServer(t, baseURL, handlerMock)
		defer server.Close()

		actual, err := serviceMock.GetNearestAvailableDrivers(context.TODO(), location)
		assert.Error(t, err)
		assert.Nil(t, actual)
	})

	t.Run("return error - json data response is broken", func(t *testing.T) {
		failedJson = []byte(`{`)
		handlerMock := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(failedJson)
		})
		server := mockHTTPServer(t, baseURL, handlerMock)
		defer server.Close()

		actual, err := serviceMock.GetNearestAvailableDrivers(context.TODO(), location)
		assert.Error(t, err)
		assert.Nil(t, actual)
	})

	t.Run("return error - connection refused", func(t *testing.T) {
		// no server running
		actual, err := serviceMock.GetNearestAvailableDrivers(context.TODO(), location)
		assert.Error(t, err)
		assert.Nil(t, actual)
	})
}

func TestLocationRepository_GetRidePath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	baseURL := "127.0.0.1:7171"

	mockConfig := &configs.Config{
		LocationServiceURL:    "http://" + baseURL,
		LocationServiceAPIKey: "mock-api-key",
	}
	serviceMock := NewLocationRepository(mockConfig, http_client.HttpClient())

	rideID := int64(666)
	driverID := int64(1111)
	ridePath := []model.Coordinate{
		{Longitude: 1, Latitude: 2}, {Longitude: 2, Latitude: 3},
	}

	data, _ := json.Marshal(model.GetRidePathResponse{
		Path: ridePath,
	})

	responseMock := http_utils.ClientResponse{
		Meta: http_utils.MetaResponse{
			Code:    0,
			Message: "success",
		},
		Data: data,
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
		server := mockHTTPServer(t, baseURL, handlerMock)
		defer server.Close()

		actual, err := serviceMock.GetRidePath(context.TODO(), rideID, driverID)
		assert.NoError(t, err)
		assert.Equal(t, ridePath, actual)
	})

	t.Run("return error - error from server", func(t *testing.T) {
		handlerMock := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(failedJson)
		})
		server := mockHTTPServer(t, baseURL, handlerMock)
		defer server.Close()

		actual, err := serviceMock.GetRidePath(context.TODO(), rideID, driverID)
		assert.Error(t, err)
		assert.Nil(t, actual)
	})

	t.Run("return error - json response is broken", func(t *testing.T) {
		failedJson = []byte(`{`)
		handlerMock := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(failedJson)
		})
		server := mockHTTPServer(t, baseURL, handlerMock)
		defer server.Close()

		actual, err := serviceMock.GetRidePath(context.TODO(), rideID, driverID)
		assert.Error(t, err)
		assert.Nil(t, actual)
	})

	t.Run("return error - json data response is broken", func(t *testing.T) {
		failedJson, _ = json.Marshal(http_utils.ClientResponse{
			Meta: http_utils.MetaResponse{Code: 1, Message: "success"},
			Data: []byte(""),
		})
		handlerMock := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(failedJson)
		})
		server := mockHTTPServer(t, baseURL, handlerMock)
		defer server.Close()

		actual, err := serviceMock.GetRidePath(context.TODO(), rideID, driverID)
		assert.Error(t, err)
		assert.Nil(t, actual)
	})

	t.Run("return error - connection refused", func(t *testing.T) {
		// no server running
		actual, err := serviceMock.GetRidePath(context.TODO(), rideID, driverID)
		assert.Error(t, err)
		assert.Nil(t, actual)
	})
}
