package repository_postgres

const (
	queryGetRiderByMSISDN = `
		SELECT id, name, phone_number FROM riders
		WHERE phone_number = $1
	`
	queryGetDriverByMSISDN = `
		SELECT id, name, phone_number, vehicle_type, vehicle_plate FROM drivers
		WHERE phone_number = $1
	`
	queryGetRiderMSISDNByID = `
		SELECT phone_number FROM riders
		WHERE id = $1
	`
	queryGetDriverMSISDNByID = `
		SELECT phone_number FROM drivers
		WHERE id = $1
	`
	queryInsertNewRide = `
		INSERT INTO rides(rider_id, status, pickup_location, destination)
		VALUES ($1, $2, POINT($3, $4), POINT($5, $6))
		RETURNING id
	`
	queryConfirmRideDriver = `
		UPDATE rides
		SET driver_id = $1, updated_at = NOW()
		WHERE id = $2 AND driver_id IS NOT NULL
		RETURNING id, rider_id, driver_id,
				  pickup_location[0] AS "pickup_location.latitude", pickup_location[1] AS "pickup_location.longitude",
				  destination[0] AS "destination.latitude", destination[1] AS "destination.longitude"
	`
	queryConfirmRideRider = `
		UPDATE rides
		SET status = $1, updated_at = NOW()
		WHERE id = $2 AND rider_id = $3 AND driver_id IS NOT NULL
		RETURNING id, rider_id, driver_id,
				  pickup_location[0] AS "pickup_location.latitude", pickup_location[1] AS "pickup_location.longitude",
				  destination[0] AS "destination.latitude", destination[1] AS "destination.longitude"
	`
)
