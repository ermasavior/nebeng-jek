package pkg_otel

import (
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type OpenTelemetry struct {
	tracer trace.Tracer
	Meter  metric.Meter
}
