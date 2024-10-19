package model

const (
	StatusNumRideWaitingForDriver = 1
	StatusNumRideWaitingForPickup = 2
	StatusNumRideStarted          = 3
	StatusNumRideEnded            = 4
	StatusNumRidePaid             = 5

	StatusRideWaitingForDriver = "WAITING_FOR_DRIVER"
	StatusRideWaitingForPickup = "WAITING_FOR_PICKUP"
	StatusRideStarted          = "RIDE_STARTED"
	StatusRideEnded            = "RIDE_ENDED"
	StatusRidePaid             = "RIDE_PAID"

	RidePricePerKm  = 3000
	RideFeeDiscount = 0.3 // 30% percentage
)

var (
	mapStatusRide = map[int]string{
		StatusNumRideWaitingForDriver: StatusRideWaitingForDriver,
		StatusNumRideWaitingForPickup: StatusRideWaitingForPickup,
		StatusNumRideStarted:          StatusRideStarted,
		StatusNumRideEnded:            StatusRideEnded,
		StatusNumRidePaid:             StatusRidePaid,
	}
)

type RideData struct {
	RideID         int64      `db:"id" json:"ride_id"`
	RiderID        int64      `db:"rider_id" json:"rider_id"`
	DriverID       int64      `db:"driver_id" json:"driver_id"`
	PickupLocation Coordinate `db:"pickup_location" json:"pickup_location"`
	Destination    Coordinate `db:"destination" json:"destination"`
	Distance       *float64   `db:"distance" json:"distance"`
	Fare           *float64   `db:"fare" json:"fare"`
	StatusNum      int        `db:"status" json:"-"`
	Status         string     `json:"status"`
}

func (r *RideData) SetDistance(distance float64) {
	r.Distance = &distance
}

func (r *RideData) SetFare(fare float64) {
	r.Fare = &fare
}

func (r *RideData) SetStatus(statusNum int) {
	r.StatusNum = statusNum
	r.MapStatus()
}

func (r *RideData) MapStatus() {
	r.Status = mapStatusRide[r.StatusNum]
}

type RideRequestMessage struct {
	RideID           int64          `json:"ride_id"`
	Rider            RiderData      `json:"rider"`
	PickupLocation   Coordinate     `json:"pickup_location"`
	Destination      Coordinate     `json:"destination"`
	AvailableDrivers map[int64]bool `json:"available_drivers"`
}

type MatchedRideMessage struct {
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
