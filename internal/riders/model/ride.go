package model

const (
	EventMatchedRide = "matched_ride"
)

type Coordinate struct {
	Longitude float64 `json:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
}

type MatchedRideMessage struct {
	RideID         int64      `json:"ride_id"`
	Driver         DriverData `json:"driver"`
	PickupLocation Coordinate `json:"pickup_location"`
	Destination    Coordinate `json:"destination"`
	RiderMSISDN    string     `json:"rider_msisdn"`
}

type MatchedRideBroadcast struct {
	RideID         int64      `json:"ride_id"`
	Driver         DriverData `json:"driver"`
	PickupLocation Coordinate `json:"pickup_location"`
	Destination    Coordinate `json:"destination"`
}
