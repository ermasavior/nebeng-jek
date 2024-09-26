package repository_redis

import (
	"context"
	"nebeng-jek/internal/rides/model"

	"github.com/go-redis/redis/v8"
)

func (r *ridesRepo) AddAvailableDriver(ctx context.Context, msisdn string, location model.Coordinate) error {
	return r.cache.GeoAdd(ctx, model.KeyAvailableDrivers, &redis.GeoLocation{
		Name:      model.SetDriverKey(msisdn),
		Longitude: location.Longitude,
		Latitude:  location.Latitude,
	}).Err()
}
