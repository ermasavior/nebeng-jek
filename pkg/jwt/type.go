package jwt

import "github.com/golang-jwt/jwt/v5"

type JWTInterface interface {
	GenerateToken(keyValues map[string]interface{}) (string, error)
	ValidateJWT(tokenStr string) (jwt.MapClaims, error)
}
