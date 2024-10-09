package middleware

import (
	"nebeng-jek/pkg/jwt"
)

type ridesMiddleware struct {
	jwtGen jwt.JWTInterface
}

func NewRidesMiddleware(jwtGen jwt.JWTInterface) ridesMiddleware {
	return ridesMiddleware{
		jwtGen: jwtGen,
	}
}
