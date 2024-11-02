package repository_redis

import (
	"context"
	"nebeng-jek/internal/rides/model"
	"nebeng-jek/internal/rides/repository"
	"nebeng-jek/pkg/logger"
	pkgRedis "nebeng-jek/pkg/redis"
	"strconv"

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

func (r *ridesRepo) AddAvailableDriver(ctx context.Context, driverID int64, location model.Coordinate) error {
	return r.cache.GeoAdd(ctx, model.KeyAvailableDrivers, &redis.GeoLocation{
		Name:      strconv.FormatInt(driverID, 10),
		Longitude: location.Longitude,
		Latitude:  location.Latitude,
	}).Err()
}

func (r *ridesRepo) RemoveAvailableDriver(ctx context.Context, driverID int64) error {
	return r.cache.ZRem(ctx, model.KeyAvailableDrivers, driverID).Err()
}

func (r *ridesRepo) GetNearestAvailableDrivers(ctx context.Context, location model.Coordinate) ([]int64, error) {
	res := r.cache.GeoRadius(ctx, model.KeyAvailableDrivers, location.Longitude, location.Latitude, &redis.GeoRadiusQuery{
		Radius:   model.NearestRadius,
		Unit:     model.NearestRadiusUnit,
		WithDist: true,
	})

	drivers, err := res.Result()
	if err != nil {
		logger.Error(ctx, "error get result", map[string]interface{}{logger.ErrorKey: err})
		return nil, err
	}

	driverIDs := make([]int64, 0, len(drivers))
	for _, d := range drivers {
		id, _ := strconv.ParseInt(d.Name, 10, 64)
		driverIDs = append(driverIDs, id)
	}

	return driverIDs, nil
}

func (r *ridesRepo) GetRidePath(ctx context.Context, rideID int64, driverID int64) ([]model.Coordinate, error) {
	key := model.GetDriverPathKey(rideID, driverID)
	res := r.cache.ZRange(ctx, key, 0, -1)

	coordinates, err := res.Result()
	if err != nil {
		logger.Error(ctx, "error get result", map[string]interface{}{logger.ErrorKey: err})
		return nil, err
	}

	result := make([]model.Coordinate, 0, len(coordinates))

	for _, coorString := range coordinates {
		coor, err := model.ParseCoordinate(coorString)
		if err != nil {
			logger.Info(ctx, "failed parsing coordinate", map[string]interface{}{
				"ride_id":       rideID,
				"coordinate":    coorString,
				logger.ErrorKey: err,
			})
			continue
		}
		result = append(result, coor)
	}
	return result, nil
}

func (r *ridesRepo) TrackUserLocation(ctx context.Context, req model.TrackUserLocationRequest) error {
	key := model.GetDriverPathKey(req.RideID, req.UserID)
	res := r.cache.ZAdd(ctx, key, &redis.Z{
		Score:  float64(req.Timestamp),
		Member: req.Location.ToStringValue(req.Timestamp),
	})
	return res.Err()
}
