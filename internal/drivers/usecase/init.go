package usecase

import (
	"context"
	"nebeng-jek/internal/drivers/model"
	"nebeng-jek/internal/drivers/repository"
	"nebeng-jek/internal/pkg/constants"
	pkg_context "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/pkg/location"
	"nebeng-jek/pkg/logger"
)

type driverUsecase struct {
	repo repository.RidesPubsubRepository
}

func NewDriverUsecase(repo repository.RidesPubsubRepository) DriverUsecase {
	return &driverUsecase{
		repo: repo,
	}
}

func (uc *driverUsecase) TrackUserLocation(ctx context.Context, req model.TrackUserLocationRequest) error {
	msg := location.TrackUserLocationMessage{
		RideID:    req.RideID,
		Timestamp: req.Timestamp,
		Location:  req.Location,
		UserID:    pkg_context.GetDriverIDFromContext(ctx),
		IsDriver:  true,
	}
	err := uc.repo.BroadcastMessage(ctx, constants.TopicUserLiveLocation, msg)
	if err != nil {
		logger.Error(ctx, "error broadcasting message", map[string]interface{}{
			logger.ErrorKey: err,
		})
		return err
	}
	return nil
}
