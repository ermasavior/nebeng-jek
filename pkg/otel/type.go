package pkg_otel

import (
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type OpenTelemetry struct {
	tracer trace.Tracer
	Meter  metric.Meter
}

type exporters struct {
	TraceExporter  sdktrace.SpanExporter
	MetricExporter sdkmetric.Exporter
}
