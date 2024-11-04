package model

import "nebeng-jek/internal/pkg/location"

type GetNearestAvailableDriversResponse struct {
	DriverIDs []int64 `json:"driver_ids"`
}

type GetRidePathResponse struct {
	Path []location.Coordinate `json:"path"`
}
