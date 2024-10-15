package repository_redis

import (
	"context"
	"fmt"
	"nebeng-jek/internal/rides/model"
	"nebeng-jek/internal/rides/repository"
	"nebeng-jek/pkg/logger"
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

func (r *ridesRepo) GetRidePath(ctx context.Context, rideID int64, msisdn string) ([]model.Coordinate, error) {
	key := model.GetDriverPathKey(rideID, msisdn)
	res := r.cache.ZRange(ctx, key, 0, -1)

	if res.Err() != nil {
		return nil, res.Err()
	}

	coordinates, err := res.Result()
	if err != nil {
		return nil, err
	}

	result := make([]model.Coordinate, 0, len(coordinates))

	for _, coorString := range coordinates {
		coor, err := model.ParseCoordinate(coorString)
		if err != nil {
			logger.Info(ctx, "failed parsing coordinate", map[string]interface{}{
				"rideID":     rideID,
				"msisdn":     msisdn,
				"coordinate": coorString,
			})
			continue
		}
		result = append(result, coor)
	}
	return result, nil
}

func (r *ridesRepo) TrackUserLocation(ctx context.Context, req model.TrackUserLocationRequest) error {
	key := model.GetDriverPathKey(req.RideID, req.MSISDN)
	res := r.cache.ZAdd(ctx, key, &redis.Z{
		Score:  float64(req.Timestamp),
		Member: fmt.Sprintf(model.CoordinateFormat, req.Location.Latitude, req.Location.Longitude, req.Timestamp),
	})
	return res.Err()
}
