package model

type RiderCreateNewRideRequest struct {
	RiderID        int64      `json:"-"`
	PickupLocation Coordinate `json:"pickup_location" binding:"required"`
	Destination    Coordinate `json:"destination" binding:"required"`
}

type DriverSetAvailabilityRequest struct {
	IsAvailable     bool       `json:"is_available"`
	CurrentLocation Coordinate `json:"current_location" binding:"required"`
}

type ConfirmRideRiderRequest struct {
	RiderID  int64 `json:"-"`
	RideID   int64 `json:"ride_id" binding:"required"`
	IsAccept bool  `json:"is_accept" binding:"required"`
}

type DriverConfirmRideRequest struct {
	DriverID int64 `json:"-"`
	RideID   int64 `json:"ride_id" binding:"required"`
	IsAccept bool  `json:"is_accept" binding:"required"`
}

type DriverStartRideRequest struct {
	DriverID int64 `json:"-"`
	RideID   int64 `json:"ride_id" binding:"required"`
}

type DriverEndRideRequest struct {
	DriverID int64 `json:"-"`
	RideID   int64 `json:"ride_id" binding:"required"`
}

type UpdateRideByDriverRequest struct {
	DriverID   int64
	RideID     int64
	Distance   float64
	Fare       float64
	FinalPrice float64
	Status     int
}

type UpdateRideDataRequest struct {
	RideID     int64
	DriverID   int64
	Distance   float64
	Fare       float64
	FinalPrice float64
	Status     int
}

type DriverConfirmPriceRequest struct {
	DriverID    int64   `json:"-"`
	RideID      int64   `json:"ride_id" binding:"required"`
	CustomPrice float64 `json:"custom_price"`
}

type TrackUserLocationRequest struct {
	RideID    int64      `json:"ride_id" binding:"required"`
	UserID    int64      `json:"user_id" binding:"required"`
	Timestamp int64      `json:"timestamp" binding:"required"`
	Location  Coordinate `json:"location" binding:"required"`
	IsDriver  bool       `json:"is_driver"`
}
