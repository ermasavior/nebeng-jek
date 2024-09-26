package model

import "fmt"

const (
	KeyAvailableDrivers = "available_drivers"
)

type Coordinate struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func SetDriverKey(msisdn string) string {
	return fmt.Sprintf("driver:%s", msisdn)
}
