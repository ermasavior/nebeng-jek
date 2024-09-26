package pkg_context

import (
	"context"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSetMSISDNToContext(t *testing.T) {
	t.Run("return context with set msisdn", func(t *testing.T) {
		var msisdn = "0123"
		ctx := context.Background()
		ctx = SetMSISDNToContext(ctx, msisdn)
		assert.Equal(t, msisdn, ctx.Value(keyMSISDN))
	})
}

func TestGetMSISDNFromContext(t *testing.T) {
	t.Run("return context with set msisdn", func(t *testing.T) {
		var msisdn = "0123"
		ctx := context.Background()
		ctx = SetMSISDNToContext(ctx, msisdn)

		res := GetMSISDNFromContext(ctx)
		assert.Equal(t, msisdn, res)
	})
}
