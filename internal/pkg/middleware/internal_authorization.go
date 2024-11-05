package middleware

import (
	"net/http"

	pkgError "nebeng-jek/pkg/error"
	http_utils "nebeng-jek/pkg/http/utils"

	"github.com/gin-gonic/gin"
)

func InternalAuthorization(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader(http_utils.HeaderApiKey)
		if apiKey == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				http_utils.NewFailedResponse(pkgError.ErrUnauthorizedCode, pkgError.ErrUnauthorizedMsg),
			)
			return
		}

		if apiKey != key {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				http_utils.NewFailedResponse(pkgError.ErrUnauthorizedCode, pkgError.ErrUnauthorizedMsg),
			)
			return
		}

		c.Next()
	}
}
