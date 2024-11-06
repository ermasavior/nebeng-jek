package model

import (
	"nebeng-jek/internal/pkg/location"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRideData_ToResponse(t *testing.T) {
	driverID := int64(7777)
	distance := float64(10)
	fare := float64(10000)
	finalPrice := float64(9000)
	timeMock := time.Now()

	r := RideData{
		RideID:   111,
		RiderID:  6666,
		DriverID: &driverID,
		PickupLocation: location.Coordinate{
			Longitude: 1,
			Latitude:  1,
		},
		Destination: location.Coordinate{
			Longitude: 2,
			Latitude:  2,
		},
		Distance:   &distance,
		Fare:       &fare,
		FinalPrice: &finalPrice,
		StatusNum:  StatusNumRideCancelled,
		Status:     StatusRideCancelled,
		StartTime:  &timeMock,
		EndTime:    &timeMock,
	}
	expected := RideDataResponse{
		RideID:   111,
		RiderID:  6666,
		DriverID: 7777,
		PickupLocation: location.Coordinate{
			Longitude: 1,
			Latitude:  1,
		},
		Destination: location.Coordinate{
			Longitude: 2,
			Latitude:  2,
		},
		Distance:   "10.00",
		Fare:       "10000.00",
		FinalPrice: "9000.00",
		Status:     StatusRideCancelled,
		StartTime:  timeMock.Format(time.RFC3339),
		EndTime:    timeMock.Format(time.RFC3339),
	}

	t.Run("get driver path key", func(t *testing.T) {
		assert.Equal(t, expected, r.ToResponse())
	})
}
