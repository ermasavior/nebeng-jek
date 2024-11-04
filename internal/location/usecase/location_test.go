package usecase

import (
	"context"
	"testing"

	"nebeng-jek/internal/location/model"
	pkgContext "nebeng-jek/internal/pkg/context"
	pkgLocation "nebeng-jek/internal/pkg/location"
	mockRedis "nebeng-jek/mock/pkg/redis"

	"github.com/go-redis/redis/v8"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestLocationUsecase_AddAvailableDriver(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	redisMock := mockRedis.NewMockCollections(ctrl)
	locationUCMock := NewLocationUsecase(redisMock)

	var (
		driverID    = int64(1111)
		driverIDKey = "1111"
		location    = pkgLocation.Coordinate{
			Longitude: 11,
			Latitude:  11,
		}
	)

	ctx := context.Background()
	ctx = pkgContext.SetDriverIDToContext(ctx, driverID)

	t.Run("success - should execute redis GEOADD", func(t *testing.T) {
		res := &redis.IntCmd{}
		redisMock.EXPECT().GeoAdd(ctx, pkgLocation.KeyAvailableDrivers, &redis.GeoLocation{
			Name:      driverIDKey,
			Longitude: location.Longitude,
			Latitude:  location.Latitude,
		}).Return(res)

		err := locationUCMock.AddAvailableDriver(ctx, driverID, location)
		assert.Nil(t, err)
	})

	t.Run("failed - should return error when GEOADD returns error", func(t *testing.T) {
		res := &redis.IntCmd{}
		res.SetErr(redis.ErrClosed)
		redisMock.EXPECT().GeoAdd(ctx, pkgLocation.KeyAvailableDrivers, &redis.GeoLocation{
			Name:      driverIDKey,
			Longitude: location.Longitude,
			Latitude:  location.Latitude,
		}).Return(res)

		err := locationUCMock.AddAvailableDriver(ctx, driverID, location)
		assert.EqualError(t, err, redis.ErrClosed.Error())
	})
}

func TestLocationUsecase_GetNearestAvailableDrivers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	redisMock := mockRedis.NewMockCollections(ctrl)
	locationUCMock := NewLocationUsecase(redisMock)

	var (
		location = pkgLocation.Coordinate{
			Longitude: 11,
			Latitude:  11,
		}
	)

	ctx := context.Background()

	t.Run("success - should execute redis GeoRadius", func(t *testing.T) {
		res := &redis.GeoLocationCmd{}
		res.SetVal([]redis.GeoLocation{
			{Name: "123"}, {Name: "456"}, {Name: "789"},
		})
		redisMock.EXPECT().GeoRadius(ctx, pkgLocation.KeyAvailableDrivers,
			location.Longitude, location.Latitude, &redis.GeoRadiusQuery{
				Radius:   pkgLocation.NearestRadius,
				Unit:     pkgLocation.NearestRadiusUnit,
				WithDist: true,
			}).Return(res)

		actual, err := locationUCMock.GetNearestAvailableDrivers(ctx, location)
		assert.Nil(t, err)
		assert.Equal(t, []int64{123, 456, 789}, actual)
	})

	t.Run("failed - should return error when GeoRadius returns error", func(t *testing.T) {
		res := &redis.GeoLocationCmd{}
		res.SetErr(redis.ErrClosed)
		redisMock.EXPECT().GeoRadius(ctx, pkgLocation.KeyAvailableDrivers,
			location.Longitude, location.Latitude, &redis.GeoRadiusQuery{
				Radius:   pkgLocation.NearestRadius,
				Unit:     pkgLocation.NearestRadiusUnit,
				WithDist: true,
			}).Return(res)

		_, err := locationUCMock.GetNearestAvailableDrivers(ctx, location)
		assert.EqualError(t, err, redis.ErrClosed.Error())
	})
}

func TestLocationUsecase_RemoveAvailableDriver(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	redisMock := mockRedis.NewMockCollections(ctrl)
	locationUCMock := locationUC{
		cache: redisMock,
	}

	var (
		driverID = int64(1111)
	)

	ctx := context.Background()
	ctx = pkgContext.SetDriverIDToContext(ctx, driverID)

	t.Run("success - should execute redis ZREM", func(t *testing.T) {
		res := &redis.IntCmd{}
		redisMock.EXPECT().ZRem(ctx, pkgLocation.KeyAvailableDrivers, driverID).Return(res)

		err := locationUCMock.RemoveAvailableDriver(ctx, driverID)
		assert.Nil(t, err)
	})

	t.Run("failed - should return error when ZREM returns error", func(t *testing.T) {
		res := &redis.IntCmd{}
		res.SetErr(redis.ErrClosed)
		redisMock.EXPECT().ZRem(ctx, pkgLocation.KeyAvailableDrivers, driverID).Return(res)

		err := locationUCMock.RemoveAvailableDriver(ctx, driverID)
		assert.EqualError(t, err, redis.ErrClosed.Error())
	})
}

func TestLocationUsecase_GetRidePath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	redisMock := mockRedis.NewMockCollections(ctrl)
	locationUCMock := NewLocationUsecase(redisMock)

	var (
		rideID   = int64(666)
		driverID = int64(1111)

		keyRedis    = model.GetDriverPathKey(rideID, driverID)
		start, stop = int64(0), int64(-1)

		expectedPath = []pkgLocation.Coordinate{
			{Longitude: 0.00000001, Latitude: -1}, {Longitude: 0.1, Latitude: -1.1}, {Longitude: 0.2, Latitude: -1.2},
		}
	)

	ctx := context.Background()

	t.Run("success - should execute redis ZRange", func(t *testing.T) {
		res := &redis.StringSliceCmd{}
		res.SetVal([]string{
			"0.00000001:-1.00000000:12345", "0.10000000:-1.10000000:12346", "0.20000000:-1.20000000:12347",
		})
		redisMock.EXPECT().ZRange(ctx, keyRedis, start, stop).Return(res)

		actual, err := locationUCMock.GetRidePath(ctx, rideID, driverID)

		assert.Nil(t, err)
		assert.Equal(t, expectedPath, actual)
	})

	t.Run("success - invalid coordinate - skip invalid coordinate", func(t *testing.T) {
		res := &redis.StringSliceCmd{}
		res.SetVal([]string{
			"INVALID-COORDINATE", "0.29999999:-1.2111111:12347",
		})
		redisMock.EXPECT().ZRange(ctx, keyRedis, start, stop).Return(res)

		actual, err := locationUCMock.GetRidePath(ctx, rideID, driverID)

		assert.Nil(t, err)
		assert.Equal(t, []pkgLocation.Coordinate{{Longitude: 0.29999999, Latitude: -1.2111111}}, actual)
	})

	t.Run("failed - should return error when ZRange returns error", func(t *testing.T) {
		res := &redis.StringSliceCmd{}
		res.SetErr(redis.ErrClosed)
		redisMock.EXPECT().ZRange(ctx, keyRedis, start, stop).Return(res)

		_, err := locationUCMock.GetRidePath(ctx, rideID, driverID)
		assert.EqualError(t, err, redis.ErrClosed.Error())
	})
}

func TestLocationUsecase_TrackUserLocation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	redisMock := mockRedis.NewMockCollections(ctrl)
	locationUCMock := locationUC{
		cache: redisMock,
	}

	var (
		driverID = int64(1111)
		req      = model.TrackUserLocationRequest{
			RideID:    666,
			UserID:    driverID,
			Timestamp: 123456789,
			Location: pkgLocation.Coordinate{
				Longitude: 1,
				Latitude:  2.3,
			},
		}

		redisKey = model.GetDriverPathKey(req.RideID, req.UserID)
	)

	ctx := context.Background()
	ctx = pkgContext.SetDriverIDToContext(ctx, driverID)

	t.Run("success - should execute redis ZADD", func(t *testing.T) {
		res := &redis.IntCmd{}
		redisMock.EXPECT().ZAdd(ctx, redisKey, &redis.Z{
			Score:  123456789,
			Member: "1.00000000:2.30000000:123456789",
		}).Return(res)

		err := locationUCMock.TrackUserLocation(ctx, req)
		assert.Nil(t, err)
	})
}