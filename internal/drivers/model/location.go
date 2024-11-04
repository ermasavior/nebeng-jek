package model

type Coordinate struct {
	Longitude float64 `json:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
}

type TrackUserLocationRequest struct {
	RideID    int64      `json:"ride_id"`
	UserID    int64      `json:"user_id"`
	Timestamp int64      `json:"timestamp"`
	Location  Coordinate `json:"location"`
	IsDriver  bool       `json:"is_driver"`
}
