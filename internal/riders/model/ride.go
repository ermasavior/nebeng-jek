package model

const (
	EventMatchedRide       = "matched_ride"
	EventRideReadyToPickup = "ride_ready_to_pickup"
)

type Coordinate struct {
	Longitude float64 `json:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
}

type MatchedRideMessage struct {
	RideID      int64      `json:"ride_id"`
	Driver      DriverData `json:"driver"`
	RiderMSISDN string     `json:"rider_msisdn"`
}

type RideReadyToPickupMessage struct {
	RideID         int64      `json:"ride_id"`
	PickupLocation Coordinate `json:"pickup_location"`
	Destination    Coordinate `json:"destination"`
	RiderMSISDN    string     `json:"rider_msisdn"`
	DriverMSISDN   string     `json:"driver_msisdn"`
}
