package model

import pkgError "nebeng-jek/pkg/error"

func ValidateDriverConfirmRide(r RideData, driverID int64) pkgError.AppError {
	if r.DriverID == nil || *r.DriverID != driverID {
		return pkgError.NewForbiddenError((pkgError.ErrForbiddenMsg))
	}
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

func ValidateEndRide(r RideData, driverID int64) pkgError.AppError {
	if r.DriverID == nil || *r.DriverID != driverID {
		return pkgError.NewForbiddenError(pkgError.ErrForbiddenMsg)
	}
	if r.StatusNum != StatusNumRideStarted {
		return pkgError.NewForbiddenError(ErrMsgInvalidRideStatus)
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
		return pkgError.NewForbiddenError("invalid fare, must not be empty")
	}
	if r.Distance == nil {
		return pkgError.NewForbiddenError("invalid distance, must not be empty")
	}
	if customPrice > *r.Fare {
		return pkgError.NewBadRequestError("custom price must be lower than fare price")
	}
	return nil
}
