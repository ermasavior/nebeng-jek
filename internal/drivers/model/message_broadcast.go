package model

const (
	EventNewRideRequest    = "new_ride_request"
	EventRideReadyToPickup = "ride_ready_to_pickup"
	EventRideStarted       = "ride_started"
	EventRideEnded         = "ride_ended"
	EventRealTimeLocation  = "real_time_location"
)

type DriverMessage struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

type NewRideRequestBroadcast struct {
	RideID         int64      `json:"ride_id"`
	Rider          RiderData  `json:"rider"`
	PickupLocation Coordinate `json:"pickup_location"`
	Destination    Coordinate `json:"destination"`
}
