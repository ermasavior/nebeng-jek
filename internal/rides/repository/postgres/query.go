package repository_postgres

const (
	queryGetRiderByMSISDN = `
		SELECT id, name, phone_number FROM riders
		WHERE phone_number = $1
	`
	queryGetRiderMSISDNByID = `
		SELECT phone_number FROM riders
		WHERE id = $1
	`
	queryGetDriverByMSISDN = `
		SELECT id, name, phone_number, vehicle_type, vehicle_plate FROM drivers
		WHERE phone_number = $1
	`
	queryInsertNewRide = `
		INSERT INTO rides(rider_id, status, pickup_location, destination)
		VALUES ($1, $2, POINT($3, $4), POINT($5, $6))
		RETURNING id
	`
	queryConfirmRideDriver = `
		UPDATE rides
		SET driver_id = $1
		WHERE id = $2 AND driver_id IS NULL
		RETURNING id, rider_id, driver_id,
				  pickup_location[0] AS "pickup_location.latitude", pickup_location[1] AS "pickup_location.longitude",
				  destination[0] AS "destination.latitude", destination[1] AS "destination.longitude"
	`
)
