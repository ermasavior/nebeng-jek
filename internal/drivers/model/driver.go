package model

type DriverMessage struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}
