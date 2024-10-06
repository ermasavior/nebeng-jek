package pkg_context

import "context"

type KeyContext string

const (
	keyMSISDN = KeyContext("msisdn")
	keyAPIKey = KeyContext("api_key")
)

func SetMSISDNToContext(ctx context.Context, msisdn string) context.Context {
	return context.WithValue(ctx, keyMSISDN, msisdn)
}

func GetMSISDNFromContext(ctx context.Context) string {
	val := ctx.Value(keyMSISDN)

	if msisdn, ok := val.(string); ok {
		return msisdn
	}

	return ""
}

func SetAPIKeyToContext(ctx context.Context, apiKey string) context.Context {
	return context.WithValue(ctx, keyAPIKey, apiKey)
}

func GetAPIKeyFromContext(ctx context.Context) string {
	val := ctx.Value(keyAPIKey)

	if apiKey, ok := val.(string); ok {
		return apiKey
	}

	return ""
}
