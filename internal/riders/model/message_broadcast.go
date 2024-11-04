package model

import "encoding/json"

const (
	EventMatchedRide       = "matched_ride"
	EventRideReadyToPickup = "ride_ready_to_pickup"
	EventRideStarted       = "ride_started"
	EventRideEnded         = "ride_ended"
	EventRidePaid          = "ride_paid"
)

type RiderMessage struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}
