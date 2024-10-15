package model

import (
	"errors"
	"strconv"
	"strings"
)

const (
	KeyAvailableDrivers = "available_drivers"

	NearestRadius     = 1
	NearestRadiusUnit = "km"

	// latitude:longitude:timestamp
	CoordinateFormat = "%.2f:%.2f:%d"
)

type Coordinate struct {
	Longitude float64 `json:"longitude" db:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" db:"latitude" binding:"required"`
}

func ParseCoordinate(coordinateStr string) (coor Coordinate, err error) {
	latlon := strings.Split(coordinateStr, ":")
	if len(latlon) < 2 {
		err = errors.New("invalid coordinate format input")
		return
	}

	coor.Latitude, err = strconv.ParseFloat(latlon[0], 64)
	if err != nil {
		return
	}

	coor.Longitude, err = strconv.ParseFloat(latlon[1], 64)
	if err != nil {
		return
	}

	return
}
