package usecase

import (
	"context"
	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
)

func (u *ridesUsecase) ConfirmRideRider(ctx context.Context, req model.ConfirmRideRiderRequest) *pkgError.AppError {
	if !req.IsAccept {
		return nil
	}
	riderID := pkgContext.GetRiderIDFromContext(ctx)

	riderData, err := u.ridesRepo.GetRiderDataByID(ctx, riderID)
	if err == constants.ErrorDataNotFound {
		return pkgError.NewUnauthorized(err, "invalid rider id")
	}
	if err != nil {
		logger.Error(ctx, "error get rider data", map[string]interface{}{
			"rider_id": riderID,
			"error":    err,
		})
		return pkgError.NewInternalServerError(err, "error get rider data")
	}

	req.RiderID = riderData.ID
	rideData, err := u.ridesRepo.ConfirmRideRider(ctx, req)
	if err == constants.ErrorDataNotFound {
		return pkgError.NewNotFound(err, "ride data is not found")
	}
	if err != nil {
		logger.Error(ctx, "error confirm ride by driver", map[string]interface{}{
			"rider_id": riderID,
			"error":    err,
		})
		return pkgError.NewInternalServerError(err, "error confirm ride by driver")
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
		return pkgError.NewInternalServerError(err, "error broadcasting ride ready to pickup")
	}

	return nil
}
