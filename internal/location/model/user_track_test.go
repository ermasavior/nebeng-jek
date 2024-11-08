package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDriverPathKey(t *testing.T) {
	t.Run("get driver path key", func(t *testing.T) {
		assert.Equal(t, "path_ride:1_driver:2222", GetDriverPathKey(1, 2222))
	})
}

func TestGetRiderPathKey(t *testing.T) {
	t.Run("get Rider path key", func(t *testing.T) {
		assert.Equal(t, "path_ride:1_rider:1111", GetRiderPathKey(1, 1111))
	})
}
