package otel_middleware

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// Recorder knows how to record and measure the metrics. This
// has the required methods to be used with the HTTP
// middlewares.
type otelRecorder struct {
	attemptsCounter metric.Int64UpDownCounter
	totalDuration   metric.Int64Histogram
	requestSize     metric.Int64Histogram
	responseSize    metric.Int64Histogram
}

type Recorder interface {
	// AddRequests increments the number of requests being processed.
	AddRequests(ctx context.Context, quantity int64, attributes []attribute.KeyValue)

	// ObserveHTTPRequestDuration measures the duration of an HTTP request.
	ObserveHTTPRequestDuration(ctx context.Context, duration time.Duration, attributes []attribute.KeyValue)

	// ObserveHTTPRequestSize measures the size of an HTTP request in bytes.
	ObserveHTTPRequestSize(ctx context.Context, sizeBytes int64, attributes []attribute.KeyValue)

	// ObserveHTTPResponseSize measures the size of an HTTP response in bytes.
	ObserveHTTPResponseSize(ctx context.Context, sizeBytes int64, attributes []attribute.KeyValue)
}

func getRecorder(meter metric.Meter) Recorder {
	attemptsCounter, _ := meter.Int64UpDownCounter("http.server.request_count", metric.WithDescription("Number of Requests"))
	totalDuration, _ := meter.Int64Histogram("http.server.duration", metric.WithDescription("Time Taken by request"), metric.WithUnit("ms"))
	requestSize, _ := meter.Int64Histogram("http.server.request_content_length", metric.WithDescription("Request Size"), metric.WithUnit("bytes"))
	responseSize, _ := meter.Int64Histogram("http.server.response_content_length", metric.WithDescription("Response Size"), metric.WithUnit("bytes"))

	return &otelRecorder{
		attemptsCounter: attemptsCounter,
		totalDuration:   totalDuration,
		requestSize:     requestSize,
		responseSize:    responseSize,
	}
}

// AddRequests increments the number of requests being processed.
func (r *otelRecorder) AddRequests(ctx context.Context, quantity int64, attributes []attribute.KeyValue) {
	r.attemptsCounter.Add(ctx, quantity, metric.WithAttributes(attributes...))
}

// ObserveHTTPRequestDuration measures the duration of an HTTP request.
func (r *otelRecorder) ObserveHTTPRequestDuration(ctx context.Context, duration time.Duration, attributes []attribute.KeyValue) {
	r.totalDuration.Record(ctx, int64(duration/time.Millisecond), metric.WithAttributes(attributes...))
}

// ObserveHTTPRequestSize measures the size of an HTTP request in bytes.
func (r *otelRecorder) ObserveHTTPRequestSize(ctx context.Context, sizeBytes int64, attributes []attribute.KeyValue) {
	r.requestSize.Record(ctx, sizeBytes, metric.WithAttributes(attributes...))
}

// ObserveHTTPResponseSize measures the size of an HTTP response in bytes.
func (r *otelRecorder) ObserveHTTPResponseSize(ctx context.Context, sizeBytes int64, attributes []attribute.KeyValue) {
	r.responseSize.Record(ctx, sizeBytes, metric.WithAttributes(attributes...))
}
