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
		logger.Error(ctx, model.ErrMsgFailGetDriverData, map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return pkgError.NewInternalServerError(model.ErrMsgFailGetDriverData)
	}

	rideData, err := u.ridesRepo.GetRideData(ctx, req.RideID)
	if err == constants.ErrorDataNotFound {
		return pkgError.NewNotFoundError(pkgError.ErrResourceNotFoundMsg)
	}
	if err != nil {
		logger.Error(ctx, model.ErrMsgFailGetRideData, map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return pkgError.NewInternalServerError(model.ErrMsgFailGetRideData)
	}

	if err := model.ValidateDriverConfirmRide(rideData, driverID); err != nil {
		return err
	}

	err = u.ridesRepo.UpdateRideData(ctx, model.UpdateRideDataRequest{
		RideID:   req.RideID,
		DriverID: driverID,
		Status:   model.StatusNumRideMatchedDriver,
	})
	if err != nil {
		logger.Error(ctx, model.ErrMsgFailUpdateRideData, map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return pkgError.NewInternalServerError(model.ErrMsgFailUpdateRideData)
	}

	err = u.ridesPubSub.BroadcastMessage(ctx, constants.TopicRideMatchedDriver, model.RideMatchedDriverMessage{
		RideID:  rideData.RideID,
		Driver:  driver,
		RiderID: rideData.RiderID,
	})
	if err != nil {
		logger.Error(ctx, model.ErrMsgFailBroadcastMessage, map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return pkgError.NewInternalServerError(model.ErrMsgFailBroadcastMessage)
	}

	return nil
}
