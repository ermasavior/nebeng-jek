package middleware

import (
	"nebeng-jek/pkg/jwt"
)

const (
	DriverID = "driver_id"
	RiderID  = "rider_id"
)

type ridesMiddleware struct {
	jwtGen jwt.JWTInterface
}

func NewRidesMiddleware(jwtGen jwt.JWTInterface) ridesMiddleware {
	return ridesMiddleware{
		jwtGen: jwtGen,
	}
}
