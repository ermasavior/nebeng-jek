package usecase

import (
	"context"
	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
	"time"
)

func (u *ridesUsecase) DriverStartRide(ctx context.Context, req model.DriverStartRideRequest) (model.RideData, pkgError.AppError) {
	driverID := pkgContext.GetDriverIDFromContext(ctx)

	rideData, err := u.ridesRepo.GetRideData(ctx, req.RideID)
	if err == constants.ErrorDataNotFound {
		return model.RideData{}, pkgError.NewNotFoundError(pkgError.ErrResourceNotFoundMsg)
	}
	if err != nil {
		logger.Error(ctx, model.ErrMsgFailGetRideData, map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(model.ErrMsgFailGetRideData)
	}

	if err := model.ValidateStartRide(rideData, driverID); err != nil {
		return model.RideData{}, err
	}

	now := time.Now()
	err = u.ridesRepo.UpdateRideData(ctx, model.UpdateRideDataRequest{
		RideID:    req.RideID,
		Status:    model.StatusNumRideStarted,
		StartTime: &now,
	})
	if err != nil {
		logger.Error(ctx, model.ErrMsgFailUpdateRideData, map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(model.ErrMsgFailUpdateRideData)
	}
	rideData.SetStatus(model.StatusNumRideStarted)

	err = u.ridesRepo.UpdateDriverStatus(ctx, model.UpdateDriverStatusRequest{
		DriverID: driverID,
		Status:   model.StatusDriverOff,
	})
	if err != nil {
		logger.Error(ctx, model.ErrMsgFailUpdateStatusDriver, map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(model.ErrMsgFailUpdateStatusDriver)
	}

	err = u.locationRepo.RemoveAvailableDriver(ctx, driverID)
	if err != nil {
		logger.Error(ctx, model.ErrMsgFailRemoveAvailableDriver, map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(model.ErrMsgFailRemoveAvailableDriver)
	}

	go func() {
		err = u.ridesPubSub.BroadcastMessage(ctx, constants.TopicRideStarted, model.RideStartedMessage{
			RideID:  rideData.RideID,
			RiderID: rideData.RiderID,
		})
		if err != nil {
			logger.Error(ctx, model.ErrMsgFailBroadcastMessage, map[string]interface{}{
				"driver_id": driverID,
				"error":     err,
			})
		}
	}()

	return rideData, nil
}
