package usecase

import (
	"context"
	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
)

func (u *ridesUsecase) ConfirmRideRider(ctx context.Context, req model.ConfirmRideRiderRequest) *pkgError.AppError {
	if !req.IsAccept {
		return nil
	}
	riderMSISDN := pkgContext.GetMSISDNFromContext(ctx)

	rider, err := u.ridesRepo.GetRiderDataByMSISDN(ctx, riderMSISDN)
	if err == constants.ErrorDataNotFound {
		return pkgError.NewNotFound(err, "rider is not found")
	}
	if err != nil {
		logger.Error(ctx, "error get rider data", map[string]interface{}{
			"msisdn": riderMSISDN,
			"error":  err,
		})
		return pkgError.NewInternalServerError(err, "error get rider data")
	}

	req.RiderID = rider.ID
	rideData, err := u.ridesRepo.ConfirmRideRider(ctx, req)
	if err == constants.ErrorDataNotFound {
		return pkgError.NewNotFound(err, "ride data is not found")
	}
	if err != nil {
		logger.Error(ctx, "error confirm ride by driver", map[string]interface{}{
			"msisdn": riderMSISDN,
			"error":  err,
		})
		return pkgError.NewInternalServerError(err, "error confirm ride by driver")
	}

	driverMSISDN, err := u.ridesRepo.GetDriverMSISDNByID(ctx, rideData.DriverID)
	if err != nil {
		logger.Error(ctx, "error get driver msisdn", map[string]interface{}{
			"msisdn": riderMSISDN,
			"error":  err,
		})
		return pkgError.NewInternalServerError(err, "error get driver msisdn")
	}

	err = u.ridesPubSub.BroadcastRideReadyToPickup(ctx, model.RideReadyToPickupMessage{
		RideID:         rideData.RideID,
		PickupLocation: rideData.PickupLocation,
		Destination:    rideData.Destination,
		RiderMSISDN:    riderMSISDN,
		DriverMSISDN:   driverMSISDN,
	})
	if err != nil {
		logger.Error(ctx, "error broadcasting ride ready to pickup", map[string]interface{}{
			"msisdn": riderMSISDN,
			"error":  err,
		})
		return pkgError.NewInternalServerError(err, "error broadcasting ride ready to pickup")
	}

	return nil
}
