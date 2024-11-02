package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRideData_SetDistance(t *testing.T) {
	r := RideData{}
	t.Run("set ride distance", func(t *testing.T) {
		r.SetDistance(10)
		assert.Equal(t, float64(10), *r.Distance)
	})
}

func TestRideData_SetFare(t *testing.T) {
	r := RideData{}
	t.Run("set ride fare", func(t *testing.T) {
		r.SetFare(10000)
		assert.Equal(t, float64(10000), *r.Fare)
	})
}

func TestRideData_SetFinalPrice(t *testing.T) {
	r := RideData{}
	t.Run("set ride final price", func(t *testing.T) {
		r.SetFinalPrice(10000)
		assert.Equal(t, float64(10000), *r.FinalPrice)
	})
}

func TestRideData_SetStatus(t *testing.T) {
	r := RideData{}
	t.Run("set ride status", func(t *testing.T) {
		r.SetStatus(StatusNumRideMatchedDriver)
		assert.Equal(t, StatusNumRideMatchedDriver, r.StatusNum)
	})
}
