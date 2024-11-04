package handler_http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	pkgLocation "nebeng-jek/internal/pkg/location"
	"nebeng-jek/internal/rides/model"
	mock_usecase "nebeng-jek/mock/usecase"
	errorPkg "nebeng-jek/pkg/error"
	httpUtils "nebeng-jek/pkg/http/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_DriverSetAvailability(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockRidesUsecase(ctrl)
	handler := httpHandler{
		usecase: mockUsecase,
	}

	url := "/"

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST(url, handler.DriverSetAvailability)

	reqBody := model.DriverSetAvailabilityRequest{
		IsAvailable: true,
		CurrentLocation: pkgLocation.Coordinate{
			Longitude: 11,
			Latitude:  11,
		},
	}
	reqBytes, _ := json.Marshal(reqBody)

	t.Run("success - returns status code 200 when successfully set driver availability", func(t *testing.T) {
		mockUsecase.EXPECT().DriverSetAvailability(gomock.Any(), reqBody).Return(nil)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(reqBytes))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := httpUtils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, nil, resBody.Data)
	})

	t.Run("failed - returns 400 status code when invalid body params", func(t *testing.T) {
		reqBody := model.DriverSetAvailabilityRequest{
			CurrentLocation: pkgLocation.Coordinate{
				Longitude: 11,
				Latitude:  11,
			},
		}
		reqBytes, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(reqBytes))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := httpUtils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("failed - returns 404 - usecase returns not found", func(t *testing.T) {
		expectedError := errorPkg.NewNotFoundError("not found")

		mockUsecase.EXPECT().DriverSetAvailability(gomock.Any(), reqBody).Return(expectedError)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(reqBytes))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := httpUtils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, expectedError.GetMessage(), resBody.Meta.Message)
	})

	t.Run("failed - returns 500 - usecase returns error", func(t *testing.T) {
		expectedError := errorPkg.NewInternalServerError("error from usecase")

		mockUsecase.EXPECT().DriverSetAvailability(gomock.Any(), reqBody).Return(expectedError)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(reqBytes))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := httpUtils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, expectedError.GetMessage(), resBody.Meta.Message)
	})
}
