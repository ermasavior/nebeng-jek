package model

type DriverData struct {
	Name         string `json:"name"`
	MSISDN       string `json:"phone_number"`
	VehicleType  string `json:"vehicle_type"`
	VehiclePlate string `json:"vehicle_plate"`
}
