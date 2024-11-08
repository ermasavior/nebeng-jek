package model

import (
	"encoding/json"
	"nebeng-jek/internal/pkg/location"
)

const (
	EventNewRideRequest    = "new_ride_request"
	EventRideReadyToPickup = "ride_ready_to_pickup"
	EventRideStarted       = "ride_started"
	EventRideEnded         = "ride_ended"
)

type DriverMessage struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}

type NewRideRequestBroadcast struct {
	RideID         int64               `json:"ride_id"`
	Rider          RiderData           `json:"rider"`
	PickupLocation location.Coordinate `json:"pickup_location"`
	Destination    location.Coordinate `json:"destination"`
}
