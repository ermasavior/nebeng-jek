package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

//go:generate mockgen -package=mock_redis -source=type.go -destination=../../mock/pkg/redis/mock_redis.go
type Collections interface {
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd
	EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Conn(ctx context.Context) *redis.Conn
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd

	GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd
	GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd
	ZRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd

	Close() error
}
