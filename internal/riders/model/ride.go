package model

const (
	EventMatchedRide       = "matched_ride"
	EventRideReadyToPickup = "ride_ready_to_pickup"
	EventRideStarted       = "ride_started"
	EventRideEnded         = "ride_ended"
	EventRidePaid          = "ride_paid"
)

type Coordinate struct {
	Longitude float64 `json:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
}

type RideMatchedDriverMessage struct {
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

type RideStartedMessage struct {
	RideID      int64  `json:"ride_id"`
	RiderMSISDN string `json:"rider_msisdn"`
}

type RideEndedMessage struct {
	RideID      int64   `json:"ride_id"`
	Distance    float64 `json:"distance"`
	Fare        float64 `json:"fare"`
	RiderMSISDN string  `json:"rider_msisdn"`
}

type RidePaidMessage struct {
	RideID      int64   `json:"ride_id"`
	Distance    float64 `json:"distance"`
	FinalPrice  float64 `json:"final_price"`
	RiderMSISDN string  `json:"rider_msisdn"`
}
