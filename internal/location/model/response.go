package model

type GetNearestAvailableDriversResponse struct {
	DriverIDs []int64 `json:"driver_ids"`
}

type GetRidePathResponse struct {
	Path []Coordinate `json:"path"`
}
