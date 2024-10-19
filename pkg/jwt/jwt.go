package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtGenerator struct {
	expiryTime time.Duration
	secretKey  []byte
}

func NewJWTGenerator(exp time.Duration, secret string) JWTInterface {
	return &jwtGenerator{
		expiryTime: 24 * time.Hour, // exp,
		secretKey:  []byte(secret),
	}
}

// GenerateToken generates a new JWT token
func (j *jwtGenerator) GenerateToken(keyValues map[string]interface{}) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(j.expiryTime).Unix(),
	}

	for k, v := range keyValues {
		claims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// ValidateJWT validates the token and returns the claims
func (j *jwtGenerator) ValidateJWT(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return j.secretKey, nil
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
