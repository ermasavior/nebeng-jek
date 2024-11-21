package model

import (
	"fmt"
	"nebeng-jek/internal/pkg/location"
	"time"
)

type CreateNewRideResponse struct {
	ID int64 `json:"ride_id"`
}

type RideDataResponse struct {
	RideID         int64               `json:"ride_id"`
	RiderID        int64               `json:"rider_id"`
	DriverID       int64               `json:"driver_id,omitempty"`
	PickupLocation location.Coordinate `json:"pickup_location"`
	Destination    location.Coordinate `json:"destination,omitempty"`
	Distance       string              `json:"distance,omitempty"`
	Fare           string              `json:"fare,omitempty"`
	FinalPrice     string              `json:"final_price,omitempty"`
	Status         string              `json:"status"`
	StartTime      string              `json:"start_time,omitempty"`
	EndTime        string              `json:"end_time,omitempty"`
}

func (r RideData) ToResponse() RideDataResponse {
	var driverID int64
	if r.DriverID != nil {
		driverID = *r.DriverID
	}
	var distance string
	if r.Distance != nil {
		distance = fmt.Sprintf("%.6f", *r.Distance)
	}
	var fare string
	if r.Fare != nil {
		fare = fmt.Sprintf("%.2f", *r.Fare)
	}
	var finalPrice string
	if r.FinalPrice != nil {
		finalPrice = fmt.Sprintf("%.2f", *r.FinalPrice)
	}
	var startTime string
	if r.StartTime != nil {
		startTime = r.StartTime.Format(time.RFC3339)
	}
	var endTime string
	if r.EndTime != nil {
		endTime = r.EndTime.Format(time.RFC3339)
	}

	return RideDataResponse{
		RideID:         r.RideID,
		RiderID:        r.RiderID,
		DriverID:       driverID,
		PickupLocation: r.PickupLocation,
		Destination:    r.Destination,
		Distance:       distance,
		Fare:           fare,
		FinalPrice:     finalPrice,
		Status:         mapStatusRide[r.StatusNum],
		StartTime:      startTime,
		EndTime:        endTime,
	}
}
