package model

type DriverData struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	MSISDN       string `json:"phone_number"`
	VehicleType  string `json:"vehicle_type"`
	VehiclePlate string `json:"vehicle_plate"`
}
