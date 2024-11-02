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

func TestDriverData_MapVehicleType(t *testing.T) {
	d := DriverData{
		VehicleTypeInt: VehicleTypeIntCar,
	}
	t.Run("map vehicle type", func(t *testing.T) {
		d.MapVehicleType()
		assert.Equal(t, VehicleTypeCar, d.VehicleType)
	})
}
