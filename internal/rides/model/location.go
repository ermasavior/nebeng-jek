package model

const (
	KeyAvailableDrivers = "available_drivers"

	NearestRadius     = 1
	NearestRadiusUnit = "km"
)

type Coordinate struct {
	Longitude float64 `json:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
}
