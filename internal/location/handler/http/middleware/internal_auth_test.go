package middleware

import (
	"encoding/json"
	http_utils "nebeng-jek/pkg/http/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestMiddleware_InternalAuthorization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		path   = "/"
		apiKey = "secret"
	)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Use(InternalAuthorization(apiKey))
	router.GET(path)

	t.Run("success - return status 200", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		req.Header.Set(http_utils.HeaderApiKey, apiKey)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := http_utils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("failed - return status unauthorized 401 - no auth header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := http_utils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("failed - return status unauthorized 401 - invalid api key", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		req.Header.Set(http_utils.HeaderApiKey, "invalid")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		resBody := http_utils.Response{}
		_ = json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
