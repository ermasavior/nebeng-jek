package model

type RiderData struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	MSISDN string `json:"phone_number"`
}
