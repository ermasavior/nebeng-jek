package model

import "nebeng-jek/internal/pkg/location"

const (
	StatusNumRideNewRequest    = 1
	StatusNumRideMatchedDriver = 2
	StatusNumRideReadyToPickup = 3
	StatusNumRideStarted       = 4
	StatusNumRideEnded         = 5
	StatusNumRidePaid          = 6
	StatusNumRideCancelled     = 7

	StatusRideNewRequest    = "NEW_RIDE_REQUEST"
	StatusRideMatchedDriver = "MATCHED_DRIVER"
	StatusRideReadyToPickup = "READY_TO_PICKUP"
	StatusRideStarted       = "RIDE_STARTED"
	StatusRideEnded         = "RIDE_ENDED"
	StatusRidePaid          = "RIDE_PAID"
	StatusRideCancelled     = "RIDE_CANCELLED"

	RidePricePerKm  = 3000
	RideFeeDiscount = 0.3 // 30% percentage
)

var (
	mapStatusRide = map[int]string{
		StatusNumRideNewRequest:    StatusRideNewRequest,
		StatusNumRideMatchedDriver: StatusRideMatchedDriver,
		StatusNumRideReadyToPickup: StatusRideReadyToPickup,
		StatusNumRideStarted:       StatusRideStarted,
		StatusNumRideEnded:         StatusRideEnded,
		StatusNumRidePaid:          StatusRidePaid,
		StatusNumRideCancelled:     StatusRideCancelled,
	}
)

type RideData struct {
	RideID         int64               `db:"id" json:"ride_id"`
	RiderID        int64               `db:"rider_id" json:"rider_id"`
	DriverID       *int64              `db:"driver_id" json:"driver_id"`
	PickupLocation location.Coordinate `db:"pickup_location" json:"pickup_location"`
	Destination    location.Coordinate `db:"destination" json:"destination"`
	Distance       *float64            `db:"distance" json:"distance"`
	Fare           *float64            `db:"fare" json:"fare"`
	FinalPrice     *float64            `db:"final_price" json:"final_price"`
	StatusNum      int                 `db:"status" json:"-"`
	Status         string              `json:"status"`
}

func (r *RideData) SetDistance(distance float64) {
	r.Distance = &distance
}

func (r *RideData) SetFare(fare float64) {
	r.Fare = &fare
}

func (r *RideData) SetFinalPrice(price float64) {
	r.FinalPrice = &price
}

func (r *RideData) SetStatus(statusNum int) {
	r.StatusNum = statusNum
}

type StoreRideCommissionRequest struct {
	RideID     int64
	Commission float64
}

type UpdateRideDataRequest struct {
	RideID     int64
	DriverID   int64
	Distance   *float64
	Fare       *float64
	FinalPrice *float64
	Status     int
}

type NewRideRequestMessage struct {
	RideID            int64               `json:"ride_id"`
	Rider             RiderData           `json:"rider"`
	PickupLocation    location.Coordinate `json:"pickup_location"`
	Destination       location.Coordinate `json:"destination"`
	AvailableDriverID int64               `json:"available_driver_id"`
}

type RideMatchedDriverMessage struct {
	RideID  int64      `json:"ride_id"`
	Driver  DriverData `json:"driver"`
	RiderID int64      `json:"rider_id"`
}

type RideReadyToPickupMessage struct {
	RideID   int64 `json:"ride_id"`
	RiderID  int64 `json:"rider_id"`
	DriverID int64 `json:"driver_id"`
}

type RideStartedMessage struct {
	RideID  int64 `json:"ride_id"`
	RiderID int64 `json:"rider_id"`
}

type RideEndedMessage struct {
	RideID   int64   `json:"ride_id"`
	Distance float64 `json:"distance"`
	Fare     float64 `json:"fare"`
	RiderID  int64   `json:"rider_id"`
}

type RidePaidMessage struct {
	RideID     int64   `json:"ride_id"`
	Distance   float64 `json:"distance"`
	FinalPrice float64 `json:"final_price"`
	RiderID    int64   `json:"rider_id"`
}
