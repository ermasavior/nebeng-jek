package middleware

import (
	pkg_context "nebeng-jek/internal/pkg/context"
	httpUtils "nebeng-jek/pkg/http/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (r *ridesMiddleware) DriverAuthMiddleware(c *gin.Context) {
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

	driverID, _ := claims[DriverID].(float64)
	ctx := pkg_context.SetDriverIDToContext(c.Request.Context(), int64(driverID))
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}

func (r *ridesMiddleware) RiderAuthMiddleware(c *gin.Context) {
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

	riderID, _ := claims[RiderID].(float64)
	ctx := pkg_context.SetRiderIDToContext(c.Request.Context(), int64(riderID))
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}
