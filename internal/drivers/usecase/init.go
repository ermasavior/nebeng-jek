package usecase

import (
	"context"
	"nebeng-jek/internal/drivers/model"
	"nebeng-jek/internal/drivers/repository"
	"nebeng-jek/internal/pkg/constants"
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
	err := uc.repo.BroadcastMessage(ctx, constants.TopicUserLiveLocation, req)
	if err != nil {
		logger.Error(ctx, "error broadcasting message", map[string]interface{}{
			logger.ErrorKey: err,
		})
		return err
	}
	return nil
}
