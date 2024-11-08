package model

import (
	"nebeng-jek/internal/pkg/location"
	pkgError "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/haversine"
)

func ValidateDriverConfirmRide(r RideData) pkgError.AppError {
	if r.StatusNum != StatusNumRideNewRequest {
		return pkgError.NewForbiddenError(ErrMsgInvalidRideStatus)
	}
	return nil
}

func ValidateRiderConfirmRide(r RideData, riderID int64) pkgError.AppError {
	if r.RiderID != riderID || r.DriverID == nil {
		return pkgError.NewForbiddenError(pkgError.ErrForbiddenMsg)
	}
	if r.StatusNum != StatusNumRideMatchedDriver {
		return pkgError.NewForbiddenError(ErrMsgInvalidRideStatus)
	}
	return nil
}

func ValidateStartRide(r RideData, driverID int64) pkgError.AppError {
	if r.DriverID == nil || *r.DriverID != driverID {
		return pkgError.NewForbiddenError(pkgError.ErrForbiddenMsg)
	}
	if r.StatusNum != StatusNumRideReadyToPickup {
		return pkgError.NewForbiddenError(ErrMsgInvalidRideStatus)
	}
	return nil
}

func ValidateEndRide(r RideData, driverID int64, ridePath GetRidePathResponse) pkgError.AppError {
	if r.DriverID == nil || *r.DriverID != driverID {
		return pkgError.NewForbiddenError(pkgError.ErrForbiddenMsg)
	}
	if r.StatusNum != StatusNumRideStarted {
		return pkgError.NewForbiddenError(ErrMsgInvalidRideStatus)
	}
	if len(ridePath.DriverPath) == 0 || len(ridePath.RiderPath) == 0 {
		return pkgError.NewUnprocessableError(ErrMsgRideEmptyPath)
	}

	// compare initial positions
	dp := ridePath.DriverPath[0]
	rp := ridePath.RiderPath[0]
	if haversine.CalculateDistance(dp.Latitude, dp.Longitude, rp.Latitude, rp.Longitude) > location.ProximityThreshold {
		return pkgError.NewUnprocessableError(ErrMsgUnmatchedDriverRiderInitPosition)
	}

	// compare last position
	dp = ridePath.DriverPath[len(ridePath.DriverPath)-1]
	rp = ridePath.RiderPath[len(ridePath.RiderPath)-1]
	if haversine.CalculateDistance(dp.Latitude, dp.Longitude, rp.Latitude, rp.Longitude) > location.ProximityThreshold {
		return pkgError.NewUnprocessableError(ErrMsgUnmatchedDriverRiderLastPosition)
	}

	return nil
}

func ValidateConfirmPayment(r RideData, driverID int64, customPrice float64) pkgError.AppError {
	if r.DriverID == nil || *r.DriverID != driverID {
		return pkgError.NewForbiddenError((pkgError.ErrForbiddenMsg))
	}
	if r.StatusNum != StatusNumRideEnded {
		return pkgError.NewForbiddenError(ErrMsgInvalidRideStatus)
	}
	if r.Fare == nil {
		return pkgError.NewForbiddenError(ErrMsgInvalidFare)
	}
	if r.Distance == nil {
		return pkgError.NewForbiddenError(ErrMsgInvalidDistance)
	}
	if customPrice > *r.Fare {
		return pkgError.NewBadRequestError(ErrMsgInvalidCustomPrice)
	}
	return nil
}
