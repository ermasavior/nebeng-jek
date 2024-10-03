package model

type RiderData struct {
	ID     int64  `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	MSISDN string `json:"phone_number" db:"phone_number"`
}
