package model

const (
	EventNewRideRequest    = "new_ride_request"
	EventRideReadyToPickup = "ride_ready_to_pickup"
	EventRideStarted       = "ride_started"
	EventRideEnded         = "ride_ended"
)

type Coordinate struct {
	Longitude float64 `json:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
}

type NewRideRequestBroadcast struct {
	RideID         int64      `json:"ride_id"`
	Rider          RiderData  `json:"rider"`
	PickupLocation Coordinate `json:"pickup_location"`
	Destination    Coordinate `json:"destination"`
}

type NewRideRequestMessage struct {
	RideID           int64           `json:"ride_id"`
	Rider            RiderData       `json:"rider"`
	PickupLocation   Coordinate      `json:"pickup_location"`
	Destination      Coordinate      `json:"destination"`
	AvailableDrivers map[string]bool `json:"available_drivers"`
}

type RideReadyToPickupMessage struct {
	RideID         int64      `json:"ride_id"`
	PickupLocation Coordinate `json:"pickup_location"`
	Destination    Coordinate `json:"destination"`
	RiderMSISDN    string     `json:"rider_msisdn"`
	DriverMSISDN   string     `json:"driver_msisdn"`
}

type RideEndedMessage struct {
	RideID       int64   `json:"ride_id"`
	Distance     float64 `json:"distance"`
	Fare         float64 `json:"fare"`
	RiderMSISDN  string  `json:"rider_msisdn"`
	DriverMSISDN string  `json:"driver_msisdn"`
}
