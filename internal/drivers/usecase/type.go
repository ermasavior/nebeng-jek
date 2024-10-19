package usecase

import (
	"context"
	"nebeng-jek/internal/drivers/model"
)

type DriverUsecase interface {
	TrackUserLocation(context.Context, model.TrackUserLocationRequest) error
}
