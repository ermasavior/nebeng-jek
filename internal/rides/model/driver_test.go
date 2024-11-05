package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDriverData_MapVehicleType(t *testing.T) {
	d := DriverData{
		VehicleTypeInt: VehicleTypeIntCar,
	}
	t.Run("map vehicle type", func(t *testing.T) {
		d.MapVehicleType()
		assert.Equal(t, VehicleTypeCar, d.VehicleType)
	})
}
