package repository_postgres

const (
	queryGetRiderByMSISDN = `
		SELECT id, name, phone_number FROM riders
		WHERE phone_number = $1
	`
	queryInsertNewRide = `
		INSERT INTO rides(rider_id, status, pickup_location, destination)
		VALUES ($1, $2, POINT($3, $4), POINT($5, $6))
		RETURNING id
	`
)
