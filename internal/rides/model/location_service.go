package model

import "nebeng-jek/internal/pkg/location"

type AddAvailableDriverRequest struct {
	DriverID int64               `json:"driver_id" binding:"required"`
	Location location.Coordinate `json:"location" binding:"required"`
}

type GetNearestAvailableDriversRequest struct {
	Location location.Coordinate `json:"location" binding:"required"`
}

type GetRidePathRequest struct {
	RideID   int64 `json:"ride_id"`
	DriverID int64 `json:"driver_id"`
	RiderID  int64 `json:"rider_id"`
}

type TrackUserLocationRequest struct {
	RideID    int64               `json:"ride_id" binding:"required"`
	UserID    int64               `json:"user_id" binding:"required"`
	Timestamp int64               `json:"timestamp" binding:"required"`
	Location  location.Coordinate `json:"location" binding:"required"`
	IsDriver  bool                `json:"is_driver"`
}

type GetNearestAvailableDriversResponse struct {
	DriverIDs []int64 `json:"driver_ids"`
}

type GetRidePathResponse struct {
	DriverPath []location.Coordinate `json:"driver_path"`
	RiderPath  []location.Coordinate `json:"rider_path"`
}
