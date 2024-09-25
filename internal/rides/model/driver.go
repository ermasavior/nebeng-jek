package model

type SetDriverAvailabilityRequest struct {
	IsAvailable     bool       `json:"is_available" binding:"required"`
	CurrentLocation Coordinate `json:"current_location"`
}
