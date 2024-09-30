package model

const (
	EventNewRideRequest = "new_ride_request"
)

type DriverAllocationMessage struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

type NewRideRequestBroadcast struct {
	RideID         int64      `json:"ride_id"`
	RiderID        int64      `json:"rider_id"`
	PickupLocation Coordinate `json:"pickup_location"`
	Destination    Coordinate `json:"destination"`
}

type AcceptRideRequest struct {
	RideID int64 `json:"ride_id"`
}
