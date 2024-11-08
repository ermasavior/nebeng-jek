package model

import "nebeng-jek/internal/pkg/location"

type GetNearestAvailableDriversResponse struct {
	DriverIDs []int64 `json:"driver_ids"`
}

type GetRidePathResponse struct {
	DriverPath []location.Coordinate `json:"driver_path"`
	RiderPath  []location.Coordinate `json:"rider_path"`
}
