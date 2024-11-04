package handler_http

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_usecase "nebeng-jek/mock/usecase"
	httpUtils "nebeng-jek/pkg/http/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_RemoveAvailableDriver(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockLocationUsecase(ctrl)
	h := NewHandler(mockUsecase)

	url := "/:driver_id"

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST(url, h.RemoveAvailableDriver)

	driverID := int64(2222)
	requestURL := "/2222"

	t.Run("success - returns status code 200 when successfully confirm new ride", func(t *testing.T) {
		mockUsecase.EXPECT().RemoveAvailableDriver(gomock.Any(), driverID).Return(nil)

		req := httptest.NewRequest(http.MethodPost, requestURL, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := httpUtils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("failed - returns 400 status code when invalid param", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/invalid", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := httpUtils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("failed - returns 500 when usecase returns error", func(t *testing.T) {
		expectedError := errors.New("error from usecase")
		mockUsecase.EXPECT().RemoveAvailableDriver(gomock.Any(), driverID).Return(expectedError)

		req := httptest.NewRequest(http.MethodPost, requestURL, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := httpUtils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, expectedError.Error(), resBody.Meta.Message)
	})
}
