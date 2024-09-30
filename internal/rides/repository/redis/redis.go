package repository_redis

import (
	"context"
	"nebeng-jek/internal/rides/model"
	"nebeng-jek/internal/rides/repository"
	pkgRedis "nebeng-jek/pkg/redis"

	"github.com/go-redis/redis/v8"
)

type ridesRepo struct {
	cache pkgRedis.Collections
}

func NewRepository(cache pkgRedis.Collections) repository.RidesLocationRepository {
	return &ridesRepo{
		cache: cache,
	}
}

func (r *ridesRepo) AddAvailableDriver(ctx context.Context, msisdn string, location model.Coordinate) error {
	return r.cache.GeoAdd(ctx, model.KeyAvailableDrivers, &redis.GeoLocation{
		Name:      msisdn,
		Longitude: location.Longitude,
		Latitude:  location.Latitude,
	}).Err()
}

func (r *ridesRepo) RemoveAvailableDriver(ctx context.Context, msisdn string) error {
	return r.cache.ZRem(ctx, model.KeyAvailableDrivers, msisdn).Err()
}

func (r *ridesRepo) GetNearestAvailableDrivers(ctx context.Context, location model.Coordinate) ([]string, error) {
	res := r.cache.GeoRadius(ctx, model.KeyAvailableDrivers, location.Longitude, location.Latitude, &redis.GeoRadiusQuery{
		Radius:   model.NearestRadius,
		Unit:     model.NearestRadiusUnit,
		WithDist: true,
	})

	drivers, err := res.Result()
	if err != nil {
		return nil, err
	}

	driverMSISDN := make([]string, 0, len(drivers))
	for _, d := range drivers {
		driverMSISDN = append(driverMSISDN, d.Name)
	}

	return driverMSISDN, nil
}
