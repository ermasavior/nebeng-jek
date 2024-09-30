package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mockJWT "nebeng-jek/mock/pkg/jwt"
	httpUtils "nebeng-jek/pkg/http/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware_LoginDriverMiddleware(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockJWT := mockJWT.NewMockJWTInterface(ctrl)

	mid := NewRidesMiddleware(mockJWT)

	var (
		path  = "/"
		token = "token-test"
	)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Use(mid.LoginDriverMiddleware)
	router.GET(path)

	t.Run("failed - return status unauthorized 401 - no auth header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := httpUtils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("failed - return status unauthorized 401 - invalid token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		req.Header.Set("Authorization", "Bearer "+token)

		mockJWT.EXPECT().ValidateJWT(token).Return(jwt.MapClaims{}, jwt.ErrInvalidKey)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := httpUtils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("success - return status 200", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		req.Header.Set("Authorization", "Bearer "+token)

		mockJWT.EXPECT().ValidateJWT(token).Return(jwt.MapClaims{"msisdn": "0123"}, nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := httpUtils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
