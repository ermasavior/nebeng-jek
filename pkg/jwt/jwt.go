package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var mySecretKey = []byte("your_secret_key")

// GenerateJWT generates a new JWT token
func GenerateJWT(key, value string) (string, error) {
	claims := jwt.MapClaims{
		key:   value,
		"exp": time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(mySecretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// ValidateJWT validates the token and returns the claims
func ValidateJWT(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return mySecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, jwt.ErrInvalidKey
	}
}
