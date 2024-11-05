package model

const (
	VehicleTypeIntCar        = 1
	VehicleTypeIntMotorcycle = 2

	VehicleTypeCar        = "CAR"
	VehicleTypeMotorcycle = "MOTORCYCLE"

	StatusDriverOff       = 0
	StatusDriverAvailable = 1
)

var (
	mapVehicleType = map[int]string{
		VehicleTypeIntCar:        VehicleTypeCar,
		VehicleTypeIntMotorcycle: VehicleTypeMotorcycle,
	}
)

type DriverData struct {
	ID             int64  `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	Status         int    `json:"status" db:"status"`
	MSISDN         string `json:"phone_number" db:"phone_number"`
	VehicleType    string `json:"vehicle_type"`
	VehiclePlate   string `json:"vehicle_plate" db:"vehicle_plate"`
	VehicleTypeInt int    `json:"-" db:"vehicle_type"`
}

func (d *DriverData) MapVehicleType() {
	d.VehicleType = mapVehicleType[d.VehicleTypeInt]
}

type UpdateDriverStatusRequest struct {
	DriverID int64
	Status   int
}
