package model

type AddAvailableDriverRequest struct {
	DriverID int64      `json:"driver_id" binding:"required"`
	Location Coordinate `json:"location" binding:"required"`
}

type GetNearestAvailableDriversRequest struct {
	Location Coordinate `json:"location" binding:"required"`
}

type GetRidePathRequest struct {
	RideID   int64 `json:"ride_id" binding:"required"`
	DriverID int64 `json:"driver_id" binding:"required"`
}

type TrackUserLocationRequest struct {
	RideID    int64      `json:"ride_id" binding:"required"`
	UserID    int64      `json:"user_id" binding:"required"`
	Timestamp int64      `json:"timestamp" binding:"required"`
	Location  Coordinate `json:"location" binding:"required"`
	IsDriver  bool       `json:"is_driver"`
}
