package usecase

import (
	"context"
	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/pkg/haversine"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
)

func (u *ridesUsecase) DriverEndRide(ctx context.Context, req model.DriverEndRideRequest) (model.RideData, pkgError.AppError) {
	driverID := pkgContext.GetDriverIDFromContext(ctx)

	ridePath, err := u.locationRepo.GetRidePath(ctx, req.RideID, driverID)
	if err != nil {
		logger.Error(ctx, "error get distance traversed", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError("error get distance traversed")
	}

	distance := calculateTotalDistance(ridePath)
	fare := calculateRideFare(distance)

	rideData, err := u.ridesRepo.GetRideData(ctx, req.RideID)
	if err != nil {
		logger.Error(ctx, "error get ride data", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError("error get ride data")
	}
	if rideData.Status != model.StatusRideStarted {
		return model.RideData{}, pkgError.NewBadRequestError("invalid ride status")
	}
	if rideData.DriverID != driverID {
		return model.RideData{}, pkgError.NewForbiddenError(pkgError.ErrForbiddenMsg)
	}

	err = u.ridesRepo.UpdateRideData(ctx, model.UpdateRideDataRequest{
		RideID: req.RideID,
		Status: model.StatusNumRideEnded,
	})
	if err != nil {
		logger.Error(ctx, "error update ride by driver", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError("error update ride by driver")
	}

	err = u.ridesPubSub.BroadcastMessage(ctx, constants.TopicRideEnded, model.RideEndedMessage{
		RideID:   req.RideID,
		Distance: distance,
		Fare:     fare,
		RiderID:  rideData.RiderID,
	})
	if err != nil {
		logger.Error(ctx, "error broadcasting ride ended", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError("error broadcasting ride ended")
	}

	rideData.SetDistance(distance)
	rideData.SetFare(fare)
	return rideData, nil
}

func calculateTotalDistance(path []model.Coordinate) float64 {
	distance := float64(0)
	if len(path) >= 2 {
		i := 0
		for _, posB := range path[1:] {
			posA := path[i]
			distance += haversine.CalculateDistance(posA.Latitude, posA.Longitude, posB.Latitude, posB.Longitude)
			i += 1
		}
	}

	return distance
}

func calculateRideFare(distance float64) float64 {
	return distance * model.RidePricePerKm
}