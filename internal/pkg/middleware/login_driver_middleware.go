package middleware

import (
	pkg_context "nebeng-jek/internal/pkg/context"
	httpUtils "nebeng-jek/pkg/http/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (r *ridesMiddleware) AuthJWTMiddleware(c *gin.Context) {
	token := strings.ReplaceAll(c.GetHeader("Authorization"), "Bearer ", "")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, httpUtils.NewFailedResponse(http.StatusUnauthorized, "no token provided"))
		return
	}

	claims, err := r.jwtGen.ValidateJWT(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, httpUtils.NewFailedResponse(http.StatusUnauthorized, "invalid token"))
		return
	}

	msisdn := claims["msisdn"].(string)
	ctx := pkg_context.SetMSISDNToContext(c.Request.Context(), msisdn)
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}
