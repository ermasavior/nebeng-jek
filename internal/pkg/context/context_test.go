package pkg_context

import (
	"context"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSetDriverIDToContext(t *testing.T) {
	t.Run("return context with set driverID", func(t *testing.T) {
		var driverID = int64(123)
		ctx := context.Background()
		ctx = SetDriverIDToContext(ctx, driverID)
		assert.Equal(t, driverID, ctx.Value(keyDriverID))
	})
}

func TestGetdriverIDFromContext(t *testing.T) {
	t.Run("return context with set driverID", func(t *testing.T) {
		var driverID = int64(123)
		ctx := context.Background()
		ctx = SetDriverIDToContext(ctx, driverID)

		res := GetDriverIDFromContext(ctx)
		assert.Equal(t, driverID, res)
	})
}
