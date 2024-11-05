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

func TestGetDriverIDFromContext(t *testing.T) {
	t.Run("return context with set driverID", func(t *testing.T) {
		var driverID = int64(123)
		ctx := context.Background()
		ctx = SetDriverIDToContext(ctx, driverID)

		res := GetDriverIDFromContext(ctx)
		assert.Equal(t, driverID, res)
	})
}

func TestSetRiderIDToContext(t *testing.T) {
	t.Run("return context with set riderID", func(t *testing.T) {
		var riderID = int64(123)
		ctx := context.Background()
		ctx = SetRiderIDToContext(ctx, riderID)
		assert.Equal(t, riderID, ctx.Value(keyRiderID))
	})
}

func TestGetRiderIDFromContext(t *testing.T) {
	t.Run("return context with set riderID", func(t *testing.T) {
		var riderID = int64(123)
		ctx := context.Background()
		ctx = SetRiderIDToContext(ctx, riderID)

		res := GetRiderIDFromContext(ctx)
		assert.Equal(t, riderID, res)
	})
}
