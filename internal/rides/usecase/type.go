package usecase

import (
	"context"
	"nebeng-jek/internal/rides/model"
)

type RidesUsecase interface {
	SetDriverAvailability(context.Context, model.SetDriverAvailabilityRequest) error
}
