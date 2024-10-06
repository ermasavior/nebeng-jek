package model

const (
	StatusNumRideWaitingForDriver = 1
	StatusNumRideWaitingForPickup = 2
)

type CreateNewRideRequest struct {
	RiderID        int64      `json:"-"`
	PickupLocation Coordinate `json:"pickup_location" binding:"required"`
	Destination    Coordinate `json:"destination" binding:"required"`
}

type RideRequestMessage struct {
	RideID           int64           `json:"ride_id"`
	Rider            RiderData       `json:"rider"`
	PickupLocation   Coordinate      `json:"pickup_location"`
	Destination      Coordinate      `json:"destination"`
	AvailableDrivers map[string]bool `json:"available_drivers"`
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

type UpdateRideByIDRequest struct {
	RideID   int64
	DriverID int64
	RiderID  int64
	Status   int
}

type RideData struct {
	RideID         int64      `db:"id"`
	RiderID        int64      `db:"rider_id"`
	DriverID       int64      `db:"driver_id"`
	PickupLocation Coordinate `db:"pickup_location"`
	Destination    Coordinate `db:"destination"`
}
