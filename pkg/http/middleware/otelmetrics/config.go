package otel_middleware

import (
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type config struct {
	recorder   Recorder
	attributes func(serverName, route, requestMethod string) []attribute.KeyValue
}

func getDefaultConfig() *config {
	return &config{
		attributes: DefaultAttributes,
	}
}

var DefaultAttributes = func(serverName, route, requestMethod string) []attribute.KeyValue {
	attrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String(requestMethod),
	}

	if serverName != "" {
		attrs = append(attrs, semconv.HTTPServerNameKey.String(serverName))
	}
	if route != "" {
		attrs = append(attrs, semconv.HTTPRouteKey.String(route))
	}
	return attrs
}
