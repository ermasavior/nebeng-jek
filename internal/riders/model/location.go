package model

import "nebeng-jek/internal/pkg/location"

type TrackUserLocationRequest struct {
	RideID    int64               `json:"ride_id"`
	Timestamp int64               `json:"timestamp"`
	Location  location.Coordinate `json:"location"`
}
