package model

import "fmt"

type CreateNewRideResponse struct {
	ID int64 `json:"ride_id"`
}

type RideDataResponse struct {
	RideID         int64      `json:"ride_id"`
	RiderID        int64      `json:"rider_id"`
	DriverID       int64      `json:"driver_id,omitempty"`
	PickupLocation Coordinate `json:"pickup_location"`
	Destination    Coordinate `json:"destination,omitempty"`
	Distance       string     `json:"distance,omitempty"`
	Fare           string     `json:"fare,omitempty"`
	FinalPrice     string     `json:"final_price,omitempty"`
	Status         string     `json:"status"`
}

func (r RideData) ToResponse() RideDataResponse {
	var driverID int64
	if r.DriverID != nil {
		driverID = *r.DriverID
	}
	var distance string
	if r.Distance != nil {
		distance = fmt.Sprintf("%.2f", *r.Distance)
	}
	var fare string
	if r.Fare != nil {
		fare = fmt.Sprintf("%.2f", *r.Fare)
	}
	var finalPrice string
	if r.FinalPrice != nil {
		finalPrice = fmt.Sprintf("%.2f", *r.FinalPrice)
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
	}
}
