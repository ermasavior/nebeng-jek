package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJWTGenerator_GenerateJWT(t *testing.T) {
	var (
		secretKey  = "test"
		expiryTime = 1 * time.Hour
		keyValues  = map[string]interface{}{
			"id":    123,
			"phone": "08111",
		}
	)

	t.Run("return generated jwt token", func(t *testing.T) {
		j := NewJWTGenerator(expiryTime, secretKey)
		actual, err := j.GenerateJWT(keyValues)

		assert.NotEmpty(t, actual)
		assert.NoError(t, err)
	})
}

func TestJWTGenerator_ValidateJWT(t *testing.T) {
	var (
		secretKey  = "test"
		expiryTime = 1 * time.Hour
		keyValues  = map[string]interface{}{
			"id":    123,
			"phone": "08111",
		}

		j        = NewJWTGenerator(expiryTime, secretKey)
		token, _ = j.GenerateJWT(keyValues)
	)

	t.Run("return no error if jwt token is valid", func(t *testing.T) {
		res, err := j.ValidateJWT(token)
		assert.NotEmpty(t, res)
		assert.NoError(t, err)
	})

	t.Run("return error if jwt token is random string", func(t *testing.T) {
		res, err := j.ValidateJWT("invalid-token")
		assert.Empty(t, res)
		assert.Error(t, err)
	})

	t.Run("return error if jwt token with different secret key", func(t *testing.T) {
		jInvalid := NewJWTGenerator(expiryTime, "other-secret")
		invalidToken, _ := jInvalid.GenerateJWT(keyValues)

		res, err := j.ValidateJWT(invalidToken)
		assert.Empty(t, res)
		assert.Error(t, err)
	})
}
