package usecase

import (
	"context"
	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
)

func (u *ridesUsecase) StartRideDriver(ctx context.Context, req model.StartRideDriverRequest) *pkgError.AppError {
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

	rideData, err := u.ridesRepo.UpdateRideByDriver(ctx, model.UpdateRideByDriverRequest{
		DriverID: driver.ID,
		RideID:   req.RideID,
		Status:   model.StatusNumRideInProgress,
	})
	if err == constants.ErrorDataNotFound {
		return pkgError.NewNotFound(err, "ride data is not found or has been allocated to another driver")
	}
	if err != nil {
		logger.Error(ctx, "error update ride by driver", map[string]interface{}{
			"msisdn": msisdn,
			"error":  err,
		})
		return pkgError.NewInternalServerError(err, "error update ride by driver")
	}

	err = u.locationRepo.RemoveAvailableDriver(ctx, msisdn)
	if err != nil {
		logger.Error(ctx, "error removing available driver", map[string]interface{}{
			"msisdn": msisdn,
			"error":  err,
		})
		return pkgError.NewInternalServerError(err, "error removing available driver")
	}

	riderMSISDN, err := u.ridesRepo.GetRiderMSISDNByID(ctx, rideData.RiderID)
	if err != nil {
		logger.Error(ctx, "error get rider msisdn", map[string]interface{}{
			"msisdn": msisdn,
			"error":  err,
		})
		return pkgError.NewInternalServerError(err, "error get rider msisdn")
	}

	err = u.ridesPubSub.BroadcastRideStarted(ctx, model.RideStartedMessage{
		RideID:       rideData.RideID,
		DriverMSISDN: msisdn,
		RiderMSISDN:  riderMSISDN,
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
