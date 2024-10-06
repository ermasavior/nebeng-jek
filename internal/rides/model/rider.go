package model

type RiderData struct {
	ID     int64  `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	MSISDN string `json:"phone_number" db:"phone_number"`
}

type ConfirmRideRiderRequest struct {
	RiderID  int64 `json:"-"`
	RideID   int64 `json:"ride_id" binding:"required"`
	IsAccept bool  `json:"is_accept" binding:"required"`
}
