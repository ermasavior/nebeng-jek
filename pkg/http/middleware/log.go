package middleware

import (
	"net/http"
	"time"

	httpUtils "nebeng-jek/pkg/http/utils"
	"nebeng-jek/pkg/logger"

	"github.com/gin-gonic/gin"
)

func LogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now() // Start timer
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Fill the params
		param := gin.LogFormatterParams{}

		param.TimeStamp = time.Now() // Stop timer
		param.Latency = param.TimeStamp.Sub(start)
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.BodySize = c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}

		if param.StatusCode >= http.StatusInternalServerError {
			logger.Error(
				c.Request.Context(),
				httpUtils.MessageInternalServerError,
				map[string]interface{}{
					"client_id": param.ClientIP,
					"method":    param.Method,
					"body_size": param.BodySize,
					"path":      path,
					"latency":   param.Latency.String(),
					"error":     param.ErrorMessage,
				},
			)
		}
	}
}
