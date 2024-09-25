package repository

import (
	"context"
	"nebeng-jek/internal/rides/model"
)

type RidesRepository interface {
	AddAvailableDriver(context.Context, string, model.Coordinate) error
	RemoveAvailableDriver(context.Context, string) error
}
