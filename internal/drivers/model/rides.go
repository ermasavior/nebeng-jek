package model

type Coordinate struct {
	Longitude float64 `json:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
}

type RideRequestMessage struct {
	RideID           int64           `json:"ride_id"`
	RiderID          int64           `json:"rider_id"`
	PickupLocation   Coordinate      `json:"pickup_location"`
	Destination      Coordinate      `json:"destination"`
	AvailableDrivers map[string]bool `json:"available_drivers"`
}
