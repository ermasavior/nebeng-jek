package usecase

import (
	"context"
	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
)

func (u *ridesUsecase) DriverConfirmRide(ctx context.Context, req model.DriverConfirmRideRequest) *pkgError.AppError {
	if !req.IsAccept {
		return nil
	}
	driverID := pkgContext.GetDriverIDFromContext(ctx)

	driver, err := u.ridesRepo.GetDriverDataByID(ctx, driverID)
	if err == constants.ErrorDataNotFound {
		return pkgError.NewUnauthorized(err, "invalid driver id")
	}
	if err != nil {
		logger.Error(ctx, "error get driver data", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return pkgError.NewInternalServerError(err, "error get driver data")
	}

	req.DriverID = driver.ID
	rideData, err := u.ridesRepo.DriverConfirmRide(ctx, req)
	if err == constants.ErrorDataNotFound {
		return pkgError.NewNotFound(err, "ride data is not found or has been allocated to another driver")
	}
	if err != nil {
		logger.Error(ctx, "error confirm ride by driver", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return pkgError.NewInternalServerError(err, "error confirm ride by driver")
	}

	err = u.ridesPubSub.BroadcastMessage(ctx, constants.TopicRideMatchedDriver, model.MatchedRideMessage{
		RideID:  rideData.RideID,
		Driver:  driver,
		RiderID: rideData.RiderID,
	})
	if err != nil {
		logger.Error(ctx, "error broadcasting matched ride to rider", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return pkgError.NewInternalServerError(err, "error broadcasting matched ride to rider")
	}

	return nil
}
