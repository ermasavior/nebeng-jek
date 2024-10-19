package model

type RideMatchedDriverMessage struct {
	RideID  int64      `json:"ride_id"`
	Driver  DriverData `json:"driver"`
	RiderID int64      `json:"rider_id"`
}

type RideReadyToPickupMessage struct {
	RideID   int64 `json:"ride_id"`
	RiderID  int64 `json:"rider_id"`
	DriverID int64 `json:"driver_id"`
}

type RideStartedMessage struct {
	RideID  int64 `json:"ride_id"`
	RiderID int64 `json:"rider_id"`
}

type RideEndedMessage struct {
	RideID   int64   `json:"ride_id"`
	Distance float64 `json:"distance"`
	Fare     float64 `json:"fare"`
	RiderID  int64   `json:"rider_id"`
}

type RidePaidMessage struct {
	RideID     int64   `json:"ride_id"`
	Distance   float64 `json:"distance"`
	FinalPrice float64 `json:"final_price"`
	RiderID    int64   `json:"rider_id"`
}

type Coordinate struct {
	Longitude float64 `json:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
}
