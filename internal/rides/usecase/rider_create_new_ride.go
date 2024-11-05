package usecase

import (
	"context"
	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
)

func (u *ridesUsecase) RiderCreateNewRide(ctx context.Context, req model.CreateNewRideRequest) (int64, pkgError.AppError) {
	riderID := pkgContext.GetRiderIDFromContext(ctx)

	riderData, err := u.ridesRepo.GetRiderDataByID(ctx, riderID)
	if err == constants.ErrorDataNotFound {
		return 0, pkgError.NewUnauthorizedError("invalid rider id")
	}
	if err != nil {
		logger.Error(ctx, "error get rider data", map[string]interface{}{
			"rider_id": riderID,
			"error":    err,
		})
		return 0, pkgError.NewInternalServerError("error get rider data")
	}

	drivers, err := u.locationRepo.GetNearestAvailableDrivers(ctx, req.PickupLocation)
	if err != nil {
		logger.Error(ctx, "error get nearest available drivers", map[string]interface{}{
			"rider_id": riderID,
			"error":    err,
		})
		return 0, pkgError.NewInternalServerError("error get nearest available drivers")
	}
	if len(drivers) == 0 {
		return 0, pkgError.NewNotFoundError("no nearest driver available, try again later")
	}

	req.RiderID = riderData.ID
	rideID, err := u.ridesRepo.CreateNewRide(ctx, req)
	if err != nil {
		logger.Error(ctx, "error create new ride", map[string]interface{}{
			"rider_id": riderID,
			"req":      req,
			"error":    err,
		})
		return 0, pkgError.NewInternalServerError("error create new ride")
	}

	go func() {
		msg := model.NewRideRequestMessage{
			RideID:         rideID,
			Rider:          riderData,
			PickupLocation: req.PickupLocation,
			Destination:    req.Destination,
		}
		for _, id := range drivers {
			msg.AvailableDriverID = id
			err = u.ridesPubSub.BroadcastMessage(ctx, constants.TopicRideNewRequest, msg)
			if err != nil {
				logger.Error(ctx, "error broadcasting ride to drivers", map[string]interface{}{
					"rider_id": riderID,
					"msg":      msg,
					"error":    err,
				})
			}
		}
	}()

	return rideID, nil
}
