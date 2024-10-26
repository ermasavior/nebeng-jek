package usecase

import (
	"context"
	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
)

func (u *ridesUsecase) RiderConfirmRide(ctx context.Context, req model.RiderConfirmRideRequest) pkgError.AppError {
	if !req.IsAccept {
		return nil
	}
	riderID := pkgContext.GetRiderIDFromContext(ctx)

	rideData, err := u.ridesRepo.GetRideData(ctx, req.RideID)
	if err != nil {
		logger.Error(ctx, "error get ride data", map[string]interface{}{
			"rider_id": riderID,
			"error":    err,
		})
		return pkgError.NewInternalServerError("error get ride data")
	}

	if rideData.RiderID != riderID {
		return pkgError.NewForbiddenError(pkgError.ErrForbiddenMsg)
	}
	if rideData.Status != model.StatusRideWaitingForDriver {
		return pkgError.NewBadRequestError("invalid ride status")
	}

	err = u.ridesRepo.UpdateRideData(ctx, model.UpdateRideDataRequest{
		RideID: req.RideID,
		Status: model.StatusNumRideWaitingForPickup,
	})
	if err != nil {
		logger.Error(ctx, "error update ride data", map[string]interface{}{
			"rider_id": riderID,
			"error":    err,
		})
		return pkgError.NewInternalServerError("error update ride data")
	}

	err = u.ridesPubSub.BroadcastMessage(ctx, constants.TopicRideReadyToPickup, model.RideReadyToPickupMessage{
		RideID:   rideData.RideID,
		RiderID:  riderID,
		DriverID: rideData.DriverID,
	})
	if err != nil {
		logger.Error(ctx, "error broadcasting ride ready to pickup", map[string]interface{}{
			"rider_id": riderID,
			"error":    err,
		})
		return pkgError.NewInternalServerError("error broadcasting ride ready to pickup")
	}

	return nil
}