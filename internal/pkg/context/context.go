package pkg_context

import "context"

type KeyContext string

const (
	keyMSISDN = KeyContext("msisdn")
)

func SetMSISDNToContext(ctx context.Context, msisdn string) context.Context {
	return context.WithValue(ctx, keyMSISDN, msisdn)
}

func GetMSISDNFromContext(ctx context.Context) string {
	val := ctx.Value(keyMSISDN)

	return val.(string)
}