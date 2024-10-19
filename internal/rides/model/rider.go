package model

import "fmt"

type RiderData struct {
	ID     int64  `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	MSISDN string `json:"phone_number" db:"phone_number"`
}

func GetRiderPathKey(rideID int64, riderID int64) string {
	return fmt.Sprintf("path_ride:%d_rider:%d", rideID, riderID)
}
