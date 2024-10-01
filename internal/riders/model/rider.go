package model

type RiderMessage struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}
