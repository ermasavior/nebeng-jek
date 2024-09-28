package model

const (
	StatusNumRideWaitingForDriver = 1
)

type CreateNewRideRequest struct {
	RiderID        int64      `json:"-"`
	PickupLocation Coordinate `json:"pickup_location" binding:"required"`
	Destination    Coordinate `json:"destination" binding:"required"`
}

type RideRequestMessage struct {
	RideID           int64           `json:"ride_id"`
	RiderID          int64           `json:"rider_id"`
	PickupLocation   Coordinate      `json:"pickup_location"`
	Destination      Coordinate      `json:"destination"`
	AvailableDrivers map[string]bool `json:"available_drivers"`
}
