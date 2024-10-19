package model

type NewRideRequestMessage struct {
	RideID           int64          `json:"ride_id"`
	Rider            RiderData      `json:"rider"`
	PickupLocation   Coordinate     `json:"pickup_location"`
	Destination      Coordinate     `json:"destination"`
	AvailableDrivers map[int64]bool `json:"available_drivers"`
}

type RideReadyToPickupMessage struct {
	RideID         int64      `json:"ride_id"`
	PickupLocation Coordinate `json:"pickup_location"`
	Destination    Coordinate `json:"destination"`
	RiderID        int64      `json:"rider_id"`
	DriverID       int64      `json:"driver_id"`
}
