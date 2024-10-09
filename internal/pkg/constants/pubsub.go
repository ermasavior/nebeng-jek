package constants

const (
	TypeApplicationJSON = "application/json"
	ExchangeTypeFanout  = "fanout"

	NewRideRequestsExchange    = "ride.new_request"
	DriverAcceptedRideExchange = "ride.driver_accepted"
	RideReadyToPickupExchange  = "ride.ready_to_pickup"
	RideStartedExchange        = "ride.in_progress"
	RideEndedExchange          = "ride.arrived_at_destination"
)
