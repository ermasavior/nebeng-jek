package middleware

import (
	"nebeng-jek/pkg/utils"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
)

func TracerIDHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		spanCtx := trace.SpanContextFromContext(c.Request.Context())

		traceID := spanCtx.TraceID().String()
		if !spanCtx.HasTraceID() {
			traceID = uuid.NewString()
		}

		ctx := context.WithValue(c.Request.Context(), utils.TraceID, traceID)
		c.Request = c.Request.WithContext(ctx)

		c.Writer.Header().Add("X-Trace-Id", traceID)
		c.Next()
	}
}
