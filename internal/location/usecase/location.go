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

func (r *locationUC) GetRidePath(ctx context.Context, req model.GetRidePathRequest) (model.GetRidePathResponse, error) {
	res := model.GetRidePathResponse{}

	driverKey := model.GetDriverPathKey(req.RideID, req.DriverID)
	driverRes := r.cache.ZRange(ctx, driverKey, 0, -1)
	driverPath, err := driverRes.Result()
	if err != nil {
		logger.Error(ctx, "error get result", map[string]interface{}{
			logger.ErrorKey: err, "key": driverKey,
		})
		return res, err
	}

	riderKey := model.GetRiderPathKey(req.RideID, req.RiderID)
	riderRes := r.cache.ZRange(ctx, riderKey, 0, -1)
	riderPath, err := riderRes.Result()
	if err != nil {
		logger.Error(ctx, "error get result", map[string]interface{}{
			logger.ErrorKey: err, "key": riderKey,
		})
		return res, err
	}

	res.DriverPath = parsePathCacheResult(ctx, driverKey, driverPath)
	res.RiderPath = parsePathCacheResult(ctx, driverKey, riderPath)

	return res, nil
}

func (r *locationUC) TrackUserLocation(ctx context.Context, req model.TrackUserLocationRequest) error {
	var key string
	if req.IsDriver {
		key = model.GetDriverPathKey(req.RideID, req.UserID)
	} else {
		key = model.GetRiderPathKey(req.RideID, req.UserID)
	}

	res := r.cache.ZAdd(ctx, key, &redis.Z{
		Score:  float64(req.Timestamp),
		Member: req.Location.ToStringValue(req.Timestamp),
	})
	return res.Err()
}

func parsePathCacheResult(ctx context.Context, key string, pathResult []string) []pkgLocation.Coordinate {
	path := make([]pkgLocation.Coordinate, 0, len(pathResult))
	for _, coorString := range pathResult {
		coor, err := pkgLocation.ParseCoordinate(coorString)
		if err != nil {
			logger.Info(ctx, "failed parsing coordinate", map[string]interface{}{
				"cache_key":     key,
				"coordinate":    coorString,
				logger.ErrorKey: err,
			})
			continue
		}
		path = append(path, coor)
	}
	return path
}
