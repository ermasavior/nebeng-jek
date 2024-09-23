package http_utils

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFailedResponse(t *testing.T) {
	t.Run("returns failed response", func(t *testing.T) {
		expectedRes := Response{
			Meta: MetaResponse{
				Code:    http.StatusInternalServerError,
				Message: "Something went wrong.",
			},
		}
		res := NewFailedResponse(http.StatusInternalServerError, "Something went wrong.")
		assert.Equal(t, expectedRes, res)
	})
}

func TestNewSuccessResponse(t *testing.T) {
	t.Run("returns success response", func(t *testing.T) {
		data := map[string]string{
			"id": "1",
		}
		expectedRes := Response{
			Meta: MetaResponse{
				Code:    http.StatusOK,
				Message: "success",
			},
			Data: data,
		}

		res := NewSuccessResponse(data)
		assert.Equal(t, expectedRes, res)
	})
}
