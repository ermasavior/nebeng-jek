package model

import "fmt"

const (
	StatusNumRideWaitingForDriver = 1
	StatusNumRideWaitingForPickup = 2
	StatusNumRideInProgress       = 3
	StatusNumRideDone             = 4

	RidePricePerKm = 3000
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
	RideID      int64  `json:"ride_id"`
	RiderMSISDN string `json:"rider_msisdn"`
}

type RideEndedMessage struct {
	RideID      int64   `json:"ride_id"`
	Distance    float64 `json:"distance"`
	Fare        float64 `json:"fare"`
	RiderMSISDN string  `json:"rider_msisdn"`
}

type RideData struct {
	RideID         int64      `db:"id" json:"ride_id"`
	RiderID        int64      `db:"rider_id" json:"rider_id"`
	DriverID       int64      `db:"driver_id" json:"driver_id"`
	PickupLocation Coordinate `db:"pickup_location" json:"pickup_location"`
	Destination    Coordinate `db:"destination" json:"destination"`
	Distance       float64    `json:"distance"`
	Fare           float64    `json:"fare"`
}

func (r *RideData) SetDistance(distance float64) {
	r.Distance = distance
}

func (r *RideData) CalculateRideFare(distance float64) {
	r.Fare = distance * RidePricePerKm
}

func GetDriverPathKey(rideID int64, msisdn string) string {
	return fmt.Sprintf("path_ride:%d_driver:%s", rideID, msisdn)
}
