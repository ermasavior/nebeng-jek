package usecase

import (
	"context"
	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
)

func (u *ridesUsecase) ConfirmRideDriver(ctx context.Context, req model.ConfirmRideDriverRequest) *pkgError.AppError {
	if !req.IsAccept {
		return nil
	}
	msisdn := pkgContext.GetMSISDNFromContext(ctx)

	driver, err := u.ridesRepo.GetDriverDataByMSISDN(ctx, msisdn)
	if err == constants.ErrorDataNotFound {
		return pkgError.NewNotFound(err, "driver is not found")
	}
	if err != nil {
		logger.Error(ctx, "error get driver data", map[string]interface{}{
			"msisdn": msisdn,
			"error":  err,
		})
		return pkgError.NewInternalServerError(err, "error get driver data")
	}

	req.DriverID = driver.ID
	rideData, err := u.ridesRepo.ConfirmRideDriver(ctx, req)
	if err == constants.ErrorDataNotFound {
		return pkgError.NewNotFound(err, "ride data is not found or has been allocated to another driver")
	}
	if err != nil {
		logger.Error(ctx, "error confirm ride by driver", map[string]interface{}{
			"msisdn": msisdn,
			"error":  err,
		})
		return pkgError.NewInternalServerError(err, "error confirm ride by driver")
	}

	riderMSISDN, err := u.ridesRepo.GetRiderMSISDNByID(ctx, rideData.RiderID)
	if err != nil {
		logger.Error(ctx, "error get rider msisdn", map[string]interface{}{
			"msisdn": msisdn,
			"error":  err,
		})
		return pkgError.NewInternalServerError(err, "error get rider msisdn")
	}

	err = u.ridesPubSub.BroadcastMatchedRideToRider(ctx, model.MatchedRideMessage{
		RideID:         rideData.RideID,
		Driver:         driver,
		PickupLocation: rideData.PickupLocation,
		Destination:    rideData.Destination,
		RiderMSISDN:    riderMSISDN,
	})
	if err != nil {
		logger.Error(ctx, "error broadcasting matched ride to rider", map[string]interface{}{
			"msisdn": msisdn,
			"error":  err,
		})
		return pkgError.NewInternalServerError(err, "error broadcasting matched ride to rider")
	}

	return nil
}
