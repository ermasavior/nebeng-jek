package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetDriverKey(t *testing.T) {
	t.Run("return expected driver key", func(t *testing.T) {
		assert.Equal(t, "driver:0123", SetDriverKey("0123"))
	})
}
