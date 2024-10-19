package model

import "encoding/json"

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

func ToTrackUserLocationRequest(data interface{}) (req TrackUserLocationRequest, err error) {
	msgBytes, err := json.Marshal(data)
	if err != nil {
		return
	}

	err = json.Unmarshal(msgBytes, &req)
	return
}
