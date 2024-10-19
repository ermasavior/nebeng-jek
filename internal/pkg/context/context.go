package pkg_context

import "context"

type KeyContext string

const (
	keyDriverID = KeyContext("driver_id")
	keyRiderID  = KeyContext("rider_id")
)

func SetDriverIDToContext(ctx context.Context, driverID int64) context.Context {
	return context.WithValue(ctx, keyDriverID, driverID)
}

func GetDriverIDFromContext(ctx context.Context) int64 {
	val := ctx.Value(keyDriverID)

	if id, ok := val.(int64); ok {
		return id
	}

	return 0
}

func SetRiderIDToContext(ctx context.Context, riderID int64) context.Context {
	return context.WithValue(ctx, keyRiderID, riderID)
}

func GetRiderIDFromContext(ctx context.Context) int64 {
	val := ctx.Value(keyRiderID)

	if id, ok := val.(int64); ok {
		return id
	}

	return 0
}
