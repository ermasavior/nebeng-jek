package usecase

import (
	"context"
	"nebeng-jek/internal/riders/model"
)

type RiderUsecase interface {
	TrackUserLocation(context.Context, model.TrackUserLocationRequest) error
}
