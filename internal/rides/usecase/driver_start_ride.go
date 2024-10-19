package usecase

import (
	"context"
	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
)

func (u *ridesUsecase) DriverStartRide(ctx context.Context, req model.DriverStartRideRequest) (model.RideData, *pkgError.AppError) {
	driverID := pkgContext.GetDriverIDFromContext(ctx)

	driver, err := u.ridesRepo.GetDriverDataByID(ctx, driverID)
	if err == constants.ErrorDataNotFound {
		return model.RideData{}, pkgError.NewUnauthorized(err, "invalid driver id")
	}
	if err != nil {
		logger.Error(ctx, "error get driver data", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(err, "error get driver data")
	}

	rideData, err := u.ridesRepo.UpdateRideByDriver(ctx, model.UpdateRideByDriverRequest{
		DriverID: driver.ID,
		RideID:   req.RideID,
		Status:   model.StatusNumRideStarted,
	})
	if err == constants.ErrorDataNotFound {
		return model.RideData{}, pkgError.NewNotFound(err, "ride data is not found or has been allocated to another driver")
	}
	if err != nil {
		logger.Error(ctx, "error update ride by driver", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(err, "error update ride by driver")
	}

	err = u.locationRepo.RemoveAvailableDriver(ctx, driverID)
	if err != nil {
		logger.Error(ctx, "error removing available driver", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(err, "error removing available driver")
	}

	err = u.ridesPubSub.BroadcastMessage(ctx, constants.TopicRideStarted, model.RideStartedMessage{
		RideID:  rideData.RideID,
		RiderID: rideData.RiderID,
	})
	if err != nil {
		logger.Error(ctx, "error broadcasting message", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(err, "error broadcasting message")
	}

	return rideData, nil
}
