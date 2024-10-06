package middleware

import (
	"nebeng-jek/pkg/configs"
	"nebeng-jek/pkg/jwt"
)

type ridesMiddleware struct {
	jwtGen jwt.JWTInterface
	cfg    *configs.Config
}

func NewRidesMiddleware(jwtGen jwt.JWTInterface) ridesMiddleware {
	return ridesMiddleware{
		jwtGen: jwtGen,
	}
}
