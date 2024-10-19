package repository

import "context"

type RidesPubsubRepository interface {
	BroadcastMessage(context.Context, string, interface{}) error
}
