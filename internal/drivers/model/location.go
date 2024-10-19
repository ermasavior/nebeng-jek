package model

const (
	EventRealTimeLocation = "real_time_location"
)

type TrackUserLocationRequest struct {
	RideID    int64      `json:"ride_id"`
	MSISDN    string     `json:"msisdn"`
	Timestamp int64      `json:"timestamp"`
	Location  Coordinate `json:"location"`
	IsDriver  bool       `json:"is_driver"`
}
