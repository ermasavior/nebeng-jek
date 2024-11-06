package handler_http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"nebeng-jek/internal/rides/model"
	mock_usecase "nebeng-jek/mock/usecase"
	errorPkg "nebeng-jek/pkg/error"
	httpUtils "nebeng-jek/pkg/http/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetRideData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockRidesUsecase(ctrl)
	handler := httpHandler{
		usecase: mockUsecase,
	}

	url := "/"

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/:ride_id", handler.GetRideData)

	rideID := int64(666)
	rideIDStr := "666"

	t.Run("success - returns status code 200 when successfully confirm new ride", func(t *testing.T) {
		mockUsecase.EXPECT().GetRideData(gomock.Any(), rideID).Return(model.RideData{RideID: 666}, nil)

		req := httptest.NewRequest(http.MethodGet, url+rideIDStr, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := httpUtils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("failed - returns 400 status code when invalid body params", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, url+"invalid", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := httpUtils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("failed - returns 404 when usecase returns not found", func(t *testing.T) {
		expectedError := errorPkg.NewNotFoundError("not found")

		mockUsecase.EXPECT().GetRideData(gomock.Any(), rideID).Return(model.RideData{}, expectedError)

		req := httptest.NewRequest(http.MethodGet, url+rideIDStr, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := httpUtils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, expectedError.GetMessage(), resBody.Meta.Message)
	})

	t.Run("failed - returns 500 when usecase returns error", func(t *testing.T) {
		expectedError := errorPkg.NewInternalServerError("error from usecase")

		mockUsecase.EXPECT().GetRideData(gomock.Any(), rideID).Return(model.RideData{}, expectedError)

		req := httptest.NewRequest(http.MethodGet, url+rideIDStr, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := httpUtils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, expectedError.GetMessage(), resBody.Meta.Message)
	})
}
