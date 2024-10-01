package handler

import (
	"bytes"
	"encoding/json"
	"errors"
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

func TestHandler_CreateNewRide(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockRidesUsecase(ctrl)
	handler := ridesHandler{
		usecase: mockUsecase,
	}

	url := "/"
	rideID := int64(111)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST(url, handler.CreateNewRide)

	reqBody := model.CreateNewRideRequest{
		PickupLocation: model.Coordinate{
			Longitude: 11,
			Latitude:  11,
		},
		Destination: model.Coordinate{
			Longitude: 12,
			Latitude:  12,
		},
	}
	reqBytes, _ := json.Marshal(reqBody)

	t.Run("success - returns status code 200 when successfully create new ride", func(t *testing.T) {
		mockUsecase.EXPECT().CreateNewRide(gomock.Any(), reqBody).Return(rideID, nil)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(reqBytes))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := httpUtils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("failed - returns 400 status code when invalid body params", func(t *testing.T) {
		reqBody := model.CreateNewRideRequest{}
		reqBytes, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(reqBytes))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := httpUtils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("failed - usecase returns error", func(t *testing.T) {
		expectedError := errorPkg.NewInternalServerError(errors.New("error"), "error from usecase")

		mockUsecase.EXPECT().CreateNewRide(gomock.Any(), reqBody).Return(int64(0), expectedError)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(reqBytes))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := httpUtils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, expectedError.Message, resBody.Meta.Message)
	})
}