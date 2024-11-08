package usecase

import (
	"context"
	"nebeng-jek/internal/pkg/constants"
	pkg_context "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/pkg/location"
	"nebeng-jek/internal/riders/model"
	"nebeng-jek/internal/riders/repository"
	"nebeng-jek/pkg/logger"
)

type riderUsecase struct {
	repo repository.RidesPubsubRepository
}

func NewRiderUsecase(repo repository.RidesPubsubRepository) RiderUsecase {
	return &riderUsecase{
		repo: repo,
	}
}

func (uc *riderUsecase) TrackUserLocation(ctx context.Context, req model.TrackUserLocationRequest) error {
	msg := location.TrackUserLocationMessage{
		RideID:    req.RideID,
		Timestamp: req.Timestamp,
		Location:  req.Location,
		UserID:    pkg_context.GetRiderIDFromContext(ctx),
		IsDriver:  false,
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
