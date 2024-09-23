package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"nebeng-jek/internal/modules/products/model"
	mock_usecase "nebeng-jek/mock/usecase"
	errorPkg "nebeng-jek/pkg/error"
	httpUtils "nebeng-jek/pkg/http/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_CreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockProductUsecase(ctrl)
	handler := productHandler{
		usecase: mockUsecase,
	}

	url := "/products"

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST(url, handler.CreateProduct)

	reqBody := model.CreateProduct{
		Name:  "Laptop Lenovo XYZ",
		Price: 30000000,
	}
	reqBytes, _ := json.Marshal(reqBody)

	productID := "c5f08a52-cc46-47d1-879b-15e120885366"
	expectedRes := map[string]interface{}{
		"id": productID,
	}

	t.Run("success - returns 201 status code when successfully created new product", func(t *testing.T) {
		mockUsecase.EXPECT().CreateProduct(gomock.Any(), reqBody).Return(productID, nil)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(reqBytes))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := httpUtils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, expectedRes, resBody.Data)
	})

	t.Run("failed - returns 400 status code when invalid body params", func(t *testing.T) {
		reqBody := model.CreateProduct{
			Price: 30000000,
		}
		reqBytes, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(reqBytes))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := httpUtils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("failed - usecase returns error", func(t *testing.T) {
		expectedError := errorPkg.NewInternalServerError(errors.New("error from db"), "error while creating product")

		mockUsecase.EXPECT().CreateProduct(gomock.Any(), reqBody).Return(productID, expectedError)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(reqBytes))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := httpUtils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, expectedError.Message, resBody.Meta.Message)
	})
}
