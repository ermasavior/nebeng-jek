package model

import "fmt"

func GetDriverPathKey(rideID int64, driverID int64) string {
	return fmt.Sprintf("path_ride:%d_driver:%d", rideID, driverID)
}
