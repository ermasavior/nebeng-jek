package model

import pkgLocation "nebeng-jek/internal/pkg/location"

type CreateNewRideRequest struct {
	RiderID        int64                  `json:"-"`
	PickupLocation pkgLocation.Coordinate `json:"pickup_location" binding:"required"`
	Destination    pkgLocation.Coordinate `json:"destination" binding:"required"`
}

type DriverSetAvailabilityRequest struct {
	IsAvailable     bool                   `json:"is_available"`
	CurrentLocation pkgLocation.Coordinate `json:"current_location" binding:"required"`
}

type RiderConfirmRideRequest struct {
	RideID   int64 `json:"ride_id" binding:"required"`
	IsAccept bool  `json:"is_accept" binding:"required"`
}

type DriverConfirmRideRequest struct {
	RideID   int64 `json:"ride_id" binding:"required"`
	IsAccept bool  `json:"is_accept" binding:"required"`
}

type DriverStartRideRequest struct {
	RideID int64 `json:"ride_id" binding:"required"`
}

type DriverEndRideRequest struct {
	RideID int64 `json:"ride_id" binding:"required"`
}

type DriverConfirmPaymentRequest struct {
	RideID      int64   `json:"ride_id" binding:"required"`
	CustomPrice float64 `json:"custom_price" binding:"gte=0"`
}
