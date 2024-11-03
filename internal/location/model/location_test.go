package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoordinate_ToStringValue(t *testing.T) {
	coor := Coordinate{
		Longitude: -10.220000014,
		Latitude:  109.02000016,
	}
	timestamp := int64(123456789)
	expected := "-10.22000001:109.02000016:123456789"

	t.Run("return coordinate in string value", func(t *testing.T) {
		assert.Equal(t, expected, coor.ToStringValue(timestamp))
	})
}

func TestParseCoordinate(t *testing.T) {
	coordinateStr := "-10.22000001:109.02000016:123456789"
	expected := Coordinate{
		Longitude: -10.22000001,
		Latitude:  109.02000016,
	}

	t.Run("success - return coordinate in string value", func(t *testing.T) {
		actual, err := ParseCoordinate(coordinateStr)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("failed - return error", func(t *testing.T) {
		actual, err := ParseCoordinate("INVALID")
		assert.Error(t, err)
		assert.Equal(t, Coordinate{}, actual)
	})

	t.Run("failed - return error #2", func(t *testing.T) {
		_, err := ParseCoordinate("1:INVALID")
		assert.Error(t, err)
	})

	t.Run("failed - return error #3", func(t *testing.T) {
		_, err := ParseCoordinate("INVALID:1")
		assert.Error(t, err)
	})
}
