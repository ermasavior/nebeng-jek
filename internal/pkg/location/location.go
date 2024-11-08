package location

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

	ProximityThreshold = 0.1

	// longitude:latitude:timestamp
	coordinateFormat = "%.8f:%.8f:%d"

	EventRealTimeLocation = "real_time_location"
)

type Coordinate struct {
	Longitude float64 `json:"longitude" db:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" db:"latitude" binding:"required"`
}

func (c *Coordinate) ToStringValue(timestamp int64) string {
	return fmt.Sprintf(coordinateFormat, c.Longitude, c.Latitude, timestamp)
}

func ParseCoordinate(coordinateStr string) (Coordinate, error) {
	latlon := strings.Split(coordinateStr, ":")
	if len(latlon) < 2 {
		return Coordinate{}, errors.New("invalid coordinate format input")
	}

	var (
		coor Coordinate
		err  error
	)

	coor.Longitude, err = strconv.ParseFloat(latlon[0], 64)
	if err != nil {
		return Coordinate{}, err
	}

	coor.Latitude, err = strconv.ParseFloat(latlon[1], 64)
	if err != nil {
		return Coordinate{}, err
	}

	return coor, nil
}

type TrackUserLocationMessage struct {
	RideID    int64      `json:"ride_id"`
	UserID    int64      `json:"user_id"`
	Timestamp int64      `json:"timestamp"`
	Location  Coordinate `json:"location"`
	IsDriver  bool       `json:"is_driver"`
}
