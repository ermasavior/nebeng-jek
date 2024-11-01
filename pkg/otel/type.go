package pkg_otel

import (
	"go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type OpenTelemetry struct {
	Tracer trace.Tracer
	Meter  metric.Meter
	Logger log.Logger
}
