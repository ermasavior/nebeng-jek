package constants

const (
	TypeApplicationJSON = "application/json"
	ExchangeTypeFanout  = "fanout"

	NewRideRequestsExchange    = "ride.new_request"
	DriverAcceptedRideExchange = "ride.driver_accepted"
	RideReadyToPickupExchange  = "ride.ready_to_pickup"
	RideStartedExchange        = "ride.trip_started"
	RideEndedExchange          = "ride.trip_ended"
	RidePaidExchange           = "ride.paid"

	UserLocationLiveTrackExchange = "user.live_track"
)
