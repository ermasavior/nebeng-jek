package usecase

import (
	"context"
	"nebeng-jek/internal/location/model"
	pkgLocation "nebeng-jek/internal/pkg/location"
	"nebeng-jek/pkg/logger"
	pkgRedis "nebeng-jek/pkg/redis"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type locationUC struct {
	cache pkgRedis.Collections
}

func NewLocationUsecase(cache pkgRedis.Collections) LocationUsecase {
	return &locationUC{
		cache: cache,
	}
}

func (r *locationUC) AddAvailableDriver(ctx context.Context, driverID int64, location pkgLocation.Coordinate) error {
	return r.cache.GeoAdd(ctx, pkgLocation.KeyAvailableDrivers, &redis.GeoLocation{
		Name:      strconv.FormatInt(driverID, 10),
		Longitude: location.Longitude,
		Latitude:  location.Latitude,
	}).Err()
}

func (r *locationUC) RemoveAvailableDriver(ctx context.Context, driverID int64) error {
	return r.cache.ZRem(ctx, pkgLocation.KeyAvailableDrivers, driverID).Err()
}

func (r *locationUC) GetNearestAvailableDrivers(ctx context.Context, location pkgLocation.Coordinate) ([]int64, error) {
	res := r.cache.GeoRadius(ctx, pkgLocation.KeyAvailableDrivers, location.Longitude, location.Latitude, &redis.GeoRadiusQuery{
		Radius:   pkgLocation.NearestRadius,
		Unit:     pkgLocation.NearestRadiusUnit,
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

func (r *locationUC) GetRidePath(ctx context.Context, rideID int64, driverID int64) ([]pkgLocation.Coordinate, error) {
	key := model.GetDriverPathKey(rideID, driverID)
	res := r.cache.ZRange(ctx, key, 0, -1)

	coordinates, err := res.Result()
	if err != nil {
		logger.Error(ctx, "error get result", map[string]interface{}{logger.ErrorKey: err})
		return nil, err
	}

	result := make([]pkgLocation.Coordinate, 0, len(coordinates))

	for _, coorString := range coordinates {
		coor, err := pkgLocation.ParseCoordinate(coorString)
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

func (r *locationUC) TrackUserLocation(ctx context.Context, req model.TrackUserLocationRequest) error {
	key := model.GetDriverPathKey(req.RideID, req.UserID)
	res := r.cache.ZAdd(ctx, key, &redis.Z{
		Score:  float64(req.Timestamp),
		Member: req.Location.ToStringValue(req.Timestamp),
	})
	return res.Err()
}
