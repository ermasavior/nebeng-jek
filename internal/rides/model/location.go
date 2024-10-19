package model

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	KeyAvailableDrivers = "available_drivers"

	NearestRadius     = 1
	NearestRadiusUnit = "km"

	// longitude:latitude:timestamp
	coordinateFormat = "%.2f:%.2f:%d"
)

type Coordinate struct {
	Longitude float64 `json:"longitude" db:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" db:"latitude" binding:"required"`
}

func (c *Coordinate) ToStringValue(timestamp int64) string {
	return fmt.Sprintf(coordinateFormat, c.Longitude, c.Latitude, timestamp)
}

func ParseCoordinate(coordinateStr string) (coor Coordinate, err error) {
	latlon := strings.Split(coordinateStr, ":")
	if len(latlon) < 2 {
		err = errors.New("invalid coordinate format input")
		return
	}

	coor.Longitude, err = strconv.ParseFloat(latlon[0], 64)
	if err != nil {
		return
	}

	coor.Latitude, err = strconv.ParseFloat(latlon[1], 64)
	if err != nil {
		return
	}

	return
}
