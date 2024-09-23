package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsHandler() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Accept", "Origin", "Content-Length", "Content-Type", "Authorization",
			"Accept-Encoding", "X-CSRF-Token", "X-Trace-Id",
		},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           1 * time.Minute,
	})
}
