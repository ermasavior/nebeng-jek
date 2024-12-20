package repository_postgres

const (
	queryGetRiderDataByID = `
		SELECT id, name, phone_number FROM riders
		WHERE id = $1
	`
	queryGetDriverDataByID = `
		SELECT id, name, phone_number, vehicle_type, vehicle_plate, status FROM drivers
		WHERE id = $1
	`
	queryUpdateDriverStatus = `
		UPDATE drivers
		SET status = $1, updated_at = NOW() WHERE id = $2 
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
	queryGetRideData = `
		SELECT id, rider_id, driver_id, status, distance, fare, final_price,
			   pickup_location[0] AS "pickup_location.latitude", pickup_location[1] AS "pickup_location.longitude",
			   destination[0] AS "destination.latitude", destination[1] AS "destination.longitude",
			   start_time, end_time
		FROM rides
		WHERE id = $1
	`
	queryUpdateRideData = `
		UPDATE rides
		SET %s, updated_at = NOW()
		WHERE %s 
	`
	queryInsertRideCommission = `
		INSERT INTO ride_commissions(ride_id, platform_fee, driver_commission)
		VALUES($1, $2, $3)
	`
)
