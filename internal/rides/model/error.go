package model

const (
	ErrMsgFailGetRideData           = "error get ride data"
	ErrMsgFailUpdateRideData        = "error update ride data"
	ErrMsgFailBroadcastMessage      = "error broadcasting message"
	ErrMsgFailRemoveAvailableDriver = "error remove available driver"
	ErrMsgFailGetDriverData         = "error get driver data"
	ErrMsgFailGetDriverMSISDN       = "error get driver msisdn"
	ErrMsgFailGetRiderMSISDN        = "error get rider msisdn"
	ErrMsgFailProcessPayment        = "error processing payment"

	ErrMsgInvalidRideStatus  = "invalid ride status"
	ErrMsgInvalidFare        = "invalid fare, must not be empty"
	ErrMsgInvalidDistance    = "invalid distance, must not be empty"
	ErrMsgInvalidCustomPrice = "custom price must be lower than fare price"

	ErrMsgFailedHTTPRequest = "error request http to client"
)
