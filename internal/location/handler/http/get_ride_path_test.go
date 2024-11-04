package handler_http

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"nebeng-jek/internal/location/model"
	pkgLocation "nebeng-jek/internal/pkg/location"
	mock_usecase "nebeng-jek/mock/usecase"
	httpUtils "nebeng-jek/pkg/http/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetRidePath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockLocationUsecase(ctrl)
	h := NewHandler(mockUsecase)

	url := "/"

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST(url, h.GetRidePath)

	reqBody := model.GetRidePathRequest{
		RideID:   666,
		DriverID: 2222,
	}
	reqBytes, _ := json.Marshal(reqBody)

	ridePath := []pkgLocation.Coordinate{
		{Longitude: 1, Latitude: 2},
		{Longitude: 2, Latitude: 3},
		{Longitude: 3, Latitude: 4},
	}

	t.Run("success - returns status code 200 when successfully confirm new ride", func(t *testing.T) {
		mockUsecase.EXPECT().GetRidePath(gomock.Any(), reqBody.RideID, reqBody.DriverID).Return(ridePath, nil)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(reqBytes))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := httpUtils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("failed - returns 400 status code when invalid body params", func(t *testing.T) {
		reqBody := model.GetRidePathRequest{}
		reqBytes, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(reqBytes))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := httpUtils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("failed - returns 500 when usecase returns error", func(t *testing.T) {
		expectedError := errors.New("error from usecase")

		mockUsecase.EXPECT().GetRidePath(gomock.Any(), reqBody.RideID, reqBody.DriverID).Return(nil, expectedError)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(reqBytes))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := httpUtils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, expectedError.Error(), resBody.Meta.Message)
	})
}
