package model

const (
	VehicleTypeIntCar        = 1
	VehicleTypeIntMotorcycle = 2

	VehicleTypeCar        = "CAR"
	VehicleTypeMotorcycle = "MOTORCYCLE"
)

var (
	MapVehicleType = map[int]string{
		VehicleTypeIntCar:        VehicleTypeCar,
		VehicleTypeIntMotorcycle: VehicleTypeMotorcycle,
	}
)

type DriverData struct {
	ID             int64  `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	MSISDN         string `json:"phone_number" db:"phone_number"`
	VehicleType    string `json:"vehicle_type"`
	VehiclePlate   string `json:"vehicle_plate" db:"vehicle_plate"`
	VehicleTypeInt int    `json:"-" db:"vehicle_type"`
}

type CreateNewRideRequest struct {
	RiderID        int64      `json:"-"`
	PickupLocation Coordinate `json:"pickup_location" binding:"required"`
	Destination    Coordinate `json:"destination" binding:"required"`
}
type SetDriverAvailabilityRequest struct {
	IsAvailable     bool       `json:"is_available" binding:"required"`
	CurrentLocation Coordinate `json:"current_location" binding:"required"`
}

type ConfirmRideDriverRequest struct {
	DriverID int64 `json:"-"`
	RideID   int64 `json:"ride_id" binding:"required"`
	IsAccept bool  `json:"is_accept" binding:"required"`
}

type StartRideDriverRequest struct {
	DriverID int64 `json:"-"`
	RideID   int64 `json:"ride_id" binding:"required"`
}

type EndRideDriverRequest struct {
	DriverID int64 `json:"-"`
	RideID   int64 `json:"ride_id" binding:"required"`
}

type UpdateRideByDriverRequest struct {
	DriverID int64
	RideID   int64
	Status   int
}
