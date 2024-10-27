package usecase

import (
	"context"
	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
)

func (u *ridesUsecase) RiderConfirmRide(ctx context.Context, req model.RiderConfirmRideRequest) (model.RideData, pkgError.AppError) {
	riderID := pkgContext.GetRiderIDFromContext(ctx)

	rideData, err := u.ridesRepo.GetRideData(ctx, req.RideID)
	if err != nil {
		logger.Error(ctx, "error get ride data", map[string]interface{}{
			"rider_id": riderID,
			"error":    err,
		})
		return model.RideData{}, pkgError.NewInternalServerError("error get ride data")
	}

	if rideData.RiderID != riderID || rideData.DriverID == nil {
		return model.RideData{}, pkgError.NewForbiddenError(pkgError.ErrForbiddenMsg)
	}
	if rideData.StatusNum != model.StatusNumRideMatchedDriver {
		return model.RideData{}, pkgError.NewBadRequestError(model.ErrMsgInvalidRideStatus)
	}

	var status = model.StatusNumRideReadyToPickup
	if !req.IsAccept {
		status = model.StatusNumRideCancelled
	}

	err = u.ridesRepo.UpdateRideData(ctx, model.UpdateRideDataRequest{
		RideID: req.RideID,
		Status: status,
	})
	if err != nil {
		logger.Error(ctx, "error update ride data", map[string]interface{}{
			"rider_id": riderID,
			"error":    err,
		})
		return model.RideData{}, pkgError.NewInternalServerError("error update ride data")
	}

	rideData.SetStatus(status)

	if !req.IsAccept {
		return rideData, nil
	}

	err = u.ridesPubSub.BroadcastMessage(ctx, constants.TopicRideReadyToPickup, model.RideReadyToPickupMessage{
		RideID:   rideData.RideID,
		RiderID:  riderID,
		DriverID: *rideData.DriverID,
	})
	if err != nil {
		logger.Error(ctx, "error broadcasting ride ready to pickup", map[string]interface{}{
			"rider_id": riderID,
			"error":    err,
		})
		return model.RideData{}, pkgError.NewInternalServerError("error broadcasting ride ready to pickup")
	}

	return rideData, nil
}
