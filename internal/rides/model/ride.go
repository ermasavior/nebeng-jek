package model

const (
	StatusNumRideWaitingForDriver = 1
	StatusNumRideWaitingForPickup = 2
	StatusNumRideInProgress       = 3
)

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

type RideStartedMessage struct {
	RideID       int64  `json:"ride_id"`
	RiderMSISDN  string `json:"rider_msisdn"`
	DriverMSISDN string `json:"driver_msisdn"`
}

type RideData struct {
	RideID         int64      `db:"id"`
	RiderID        int64      `db:"rider_id"`
	DriverID       int64      `db:"driver_id"`
	PickupLocation Coordinate `db:"pickup_location"`
	Destination    Coordinate `db:"destination"`
}
