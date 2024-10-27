package usecase

import (
	"context"
	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
)

func (u *ridesUsecase) DriverStartRide(ctx context.Context, req model.DriverStartRideRequest) (model.RideData, pkgError.AppError) {
	driverID := pkgContext.GetDriverIDFromContext(ctx)

	rideData, err := u.ridesRepo.GetRideData(ctx, req.RideID)
	if err != nil {
		logger.Error(ctx, "error get ride data", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError("error get ride data")
	}

	if rideData.DriverID == nil || *rideData.DriverID != driverID {
		return model.RideData{}, pkgError.NewForbiddenError(pkgError.ErrForbiddenMsg)
	}
	if rideData.StatusNum != model.StatusNumRideReadyToPickup {
		return model.RideData{}, pkgError.NewBadRequestError(model.ErrMsgInvalidRideStatus)
	}

	err = u.ridesRepo.UpdateRideData(ctx, model.UpdateRideDataRequest{
		RideID: req.RideID,
		Status: model.StatusNumRideStarted,
	})
	if err != nil {
		logger.Error(ctx, "error update ride by driver", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError("error update ride by driver")
	}
	rideData.SetStatus(model.StatusNumRideStarted)

	err = u.locationRepo.RemoveAvailableDriver(ctx, driverID)
	if err != nil {
		logger.Error(ctx, "error removing available driver", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError("error removing available driver")
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
		return model.RideData{}, pkgError.NewInternalServerError("error broadcasting message")
	}

	return rideData, nil
}
