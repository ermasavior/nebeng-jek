package model

type NewRideRequestMessage struct {
	RideID            int64      `json:"ride_id"`
	Rider             RiderData  `json:"rider"`
	PickupLocation    Coordinate `json:"pickup_location"`
	Destination       Coordinate `json:"destination"`
	AvailableDriverID int64      `json:"available_driver_id"`
}

type RideReadyToPickupMessage struct {
	RideID   int64 `json:"ride_id"`
	RiderID  int64 `json:"rider_id"`
	DriverID int64 `json:"driver_id"`
}
