package model

import (
	"nebeng-jek/internal/pkg/location"
	pkgError "nebeng-jek/pkg/error"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDriverConfirmRide(t *testing.T) {
	var (
		riderID = int64(1111)

		r = RideData{
			RiderID:   riderID,
			StatusNum: StatusNumRideNewRequest,
		}
	)

	t.Run("success - ride data is valid", func(t *testing.T) {
		err := ValidateDriverConfirmRide(r)
		assert.Nil(t, err)
	})
	t.Run("error - status is invalid", func(t *testing.T) {
		r := RideData{
			RiderID:   riderID,
			StatusNum: StatusNumRideStarted,
		}
		err := ValidateDriverConfirmRide(r)
		assert.Equal(t, pkgError.ErrForbiddenCode, err.GetCode())
	})
}

func TestValidateRiderConfirmRide(t *testing.T) {
	var (
		riderID  = int64(1111)
		driverID = int64(2222)

		r = RideData{
			RiderID:   riderID,
			DriverID:  &driverID,
			StatusNum: StatusNumRideMatchedDriver,
		}
	)

	t.Run("success - ride data is valid", func(t *testing.T) {
		err := ValidateRiderConfirmRide(r, riderID)
		assert.Nil(t, err)
	})
	t.Run("error - status is invalid", func(t *testing.T) {
		r := RideData{
			RiderID:   riderID,
			DriverID:  &driverID,
			StatusNum: StatusNumRideStarted,
		}
		err := ValidateRiderConfirmRide(r, riderID)
		assert.Equal(t, pkgError.ErrForbiddenCode, err.GetCode())
	})
	t.Run("error - rider id is invalid", func(t *testing.T) {
		err := ValidateRiderConfirmRide(r, 999999)
		assert.Equal(t, pkgError.ErrForbiddenCode, err.GetCode())
	})
}

func TestValidateStartRide(t *testing.T) {
	var (
		riderID  = int64(1111)
		driverID = int64(2222)

		r = RideData{
			RiderID:   riderID,
			DriverID:  &driverID,
			StatusNum: StatusNumRideReadyToPickup,
		}
	)

	t.Run("success - ride data is valid", func(t *testing.T) {
		err := ValidateStartRide(r, driverID)
		assert.Nil(t, err)
	})
	t.Run("error - status is invalid", func(t *testing.T) {
		r := RideData{
			RiderID:   riderID,
			DriverID:  &driverID,
			StatusNum: StatusNumRideStarted,
		}
		err := ValidateStartRide(r, driverID)
		assert.Equal(t, pkgError.ErrForbiddenCode, err.GetCode())
	})
	t.Run("error - driver id is invalid", func(t *testing.T) {
		err := ValidateStartRide(r, 999999)
		assert.Equal(t, pkgError.ErrForbiddenCode, err.GetCode())
	})
}

func TestValidateEndRide(t *testing.T) {
	var (
		riderID  = int64(1111)
		driverID = int64(2222)

		r = RideData{
			RiderID:   riderID,
			DriverID:  &driverID,
			StatusNum: StatusNumRideStarted,
		}

		path = []location.Coordinate{
			{Longitude: 1, Latitude: 1.5}, {Longitude: 1, Latitude: 1.56},
		}
		ridePath = GetRidePathResponse{
			DriverPath: path,
			RiderPath:  path,
		}
	)

	t.Run("success - ride data is valid", func(t *testing.T) {
		err := ValidateEndRide(r, driverID, ridePath)
		assert.Nil(t, err)
	})
	t.Run("error - status is invalid", func(t *testing.T) {
		r := RideData{
			RiderID:   riderID,
			DriverID:  &driverID,
			StatusNum: StatusNumRideCancelled,
		}
		err := ValidateEndRide(r, driverID, ridePath)
		assert.Equal(t, pkgError.ErrForbiddenCode, err.GetCode())
	})
	t.Run("error - ride path is invalid", func(t *testing.T) {
		ridePath = GetRidePathResponse{
			DriverPath: []location.Coordinate{},
			RiderPath:  []location.Coordinate{},
		}
		err := ValidateEndRide(r, driverID, ridePath)
		assert.Equal(t, pkgError.ErrResourceUnprocessableCode, err.GetCode())
	})
	t.Run("error - driver and rider initial location does not match", func(t *testing.T) {
		ridePath = GetRidePathResponse{
			DriverPath: []location.Coordinate{{Longitude: 1, Latitude: 1}, {Longitude: 2, Latitude: 2}},
			RiderPath:  []location.Coordinate{{Longitude: 1, Latitude: 2}, {Longitude: 2, Latitude: 2}},
		}
		err := ValidateEndRide(r, driverID, ridePath)
		assert.Equal(t, pkgError.ErrResourceUnprocessableCode, err.GetCode())
	})
	t.Run("error - driver and rider last location does not match", func(t *testing.T) {
		ridePath = GetRidePathResponse{
			DriverPath: []location.Coordinate{{Longitude: 1, Latitude: 1}, {Longitude: 2, Latitude: 2}},
			RiderPath:  []location.Coordinate{{Longitude: 1, Latitude: 1}, {Longitude: 2, Latitude: 1}},
		}
		err := ValidateEndRide(r, driverID, ridePath)
		assert.Equal(t, pkgError.ErrResourceUnprocessableCode, err.GetCode())
	})
	t.Run("error - driver id is invalid", func(t *testing.T) {
		err := ValidateEndRide(r, 999999, ridePath)
		assert.Equal(t, pkgError.ErrForbiddenCode, err.GetCode())
	})
}

func TestValidateConfirmPayment(t *testing.T) {
	var (
		riderID     = int64(1111)
		driverID    = int64(2222)
		customPrice = float64(20000)
		fare        = float64(25000)
		distance    = float64(10)

		r = RideData{
			RiderID:   riderID,
			DriverID:  &driverID,
			StatusNum: StatusNumRideEnded,
			Fare:      &fare,
			Distance:  &distance,
		}
	)

	t.Run("success - ride data is valid", func(t *testing.T) {
		err := ValidateConfirmPayment(r, driverID, customPrice)
		assert.Nil(t, err)
	})
	t.Run("error - status is invalid", func(t *testing.T) {
		r := RideData{
			RiderID:   riderID,
			DriverID:  &driverID,
			Distance:  &distance,
			Fare:      &fare,
			StatusNum: StatusNumRideCancelled,
		}
		err := ValidateConfirmPayment(r, driverID, customPrice)
		assert.Equal(t, pkgError.ErrForbiddenCode, err.GetCode())
	})
	t.Run("error - driver id is invalid", func(t *testing.T) {
		err := ValidateConfirmPayment(r, 99999, customPrice)
		assert.Equal(t, pkgError.ErrForbiddenCode, err.GetCode())
	})
	t.Run("error - invalid fare", func(t *testing.T) {
		r := RideData{
			RiderID:   riderID,
			DriverID:  &driverID,
			Distance:  &distance,
			Fare:      nil,
			StatusNum: StatusNumRideEnded,
		}
		err := ValidateConfirmPayment(r, driverID, customPrice)
		assert.Equal(t, pkgError.ErrForbiddenCode, err.GetCode())
	})
	t.Run("error - invalid distance", func(t *testing.T) {
		r := RideData{
			RiderID:   riderID,
			DriverID:  &driverID,
			Distance:  nil,
			Fare:      &fare,
			StatusNum: StatusNumRideEnded,
		}
		err := ValidateConfirmPayment(r, driverID, customPrice)
		assert.Equal(t, pkgError.ErrForbiddenCode, err.GetCode())
	})
	t.Run("error - custom price is invalid", func(t *testing.T) {
		err := ValidateConfirmPayment(r, driverID, 5000000)
		assert.Equal(t, pkgError.ErrBadRequestCode, err.GetCode())
	})
}
