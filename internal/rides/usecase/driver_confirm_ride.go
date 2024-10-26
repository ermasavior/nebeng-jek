package usecase

import (
	"context"
	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
)

func (u *ridesUsecase) DriverConfirmRide(ctx context.Context, req model.DriverConfirmRideRequest) pkgError.AppError {
	if !req.IsAccept {
		return nil
	}
	driverID := pkgContext.GetDriverIDFromContext(ctx)

	driver, err := u.ridesRepo.GetDriverDataByID(ctx, driverID)
	if err == constants.ErrorDataNotFound {
		return pkgError.NewUnauthorizedError("invalid driver id")
	}
	if err != nil {
		logger.Error(ctx, "error get driver data", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return pkgError.NewInternalServerError("error get driver data")
	}

	rideData, err := u.ridesRepo.GetRideData(ctx, req.RideID)
	if err != nil {
		logger.Error(ctx, "error get ride data", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return pkgError.NewInternalServerError("error get ride data")
	}

	if rideData.Status != model.StatusRideWaitingForDriver {
		return pkgError.NewForbiddenError("invalid ride status")
	}

	err = u.ridesRepo.UpdateRideData(ctx, model.UpdateRideDataRequest{
		RideID: req.RideID,
		Status: model.StatusNumRideWaitingForPickup,
	})
	if err != nil {
		logger.Error(ctx, "error update ride data", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return pkgError.NewInternalServerError("error update ride data")
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
		return pkgError.NewInternalServerError("error broadcasting matched ride to rider")
	}

	return nil
}