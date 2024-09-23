package otel_middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

// Middleware returns middleware that will trace incoming requests.
// The service parameter should describe the name of the (virtual)
// server handling the request.
func Middleware(service string, otelMeter metric.Meter) gin.HandlerFunc {
	cfg := getDefaultConfig()

	recorder := cfg.recorder
	if recorder == nil {
		recorder = getRecorder(otelMeter)
	}
	return func(ginCtx *gin.Context) {
		ctx := ginCtx.Request.Context()

		route := ginCtx.FullPath()
		if len(route) <= 0 {
			route = "nonconfigured"
		}

		start := time.Now()
		reqAttributes := cfg.attributes(service, route, ginCtx.Request.Method)

		defer func() {
			resAttributes := append(reqAttributes[0:0], reqAttributes...)

			// status code metrics
			code := ginCtx.Writer.Status()
			resAttributes = append(resAttributes, semconv.HTTPStatusCodeKey.Int(code))

			// number of requests metrics
			recorder.AddRequests(ctx, 1, resAttributes)

			// request and response size metrics
			requestSize := computeApproximateRequestSize(ginCtx.Request)
			recorder.ObserveHTTPRequestSize(ctx, requestSize, resAttributes)
			recorder.ObserveHTTPResponseSize(ctx, int64(ginCtx.Writer.Size()), resAttributes)

			// request latency metrics
			recorder.ObserveHTTPRequestDuration(ctx, time.Since(start), resAttributes)
		}()

		ginCtx.Next()
	}
}

func computeApproximateRequestSize(r *http.Request) int64 {
	s := 0
	if r.URL != nil {
		s = len(r.URL.Path)
	}

	s += len(r.Method)
	s += len(r.Proto)
	for name, values := range r.Header {
		s += len(name)
		for _, value := range values {
			s += len(value)
		}
	}
	s += len(r.Host)

	// N.B. r.Form and r.MultipartForm are assumed to be included in r.URL.

	if r.ContentLength != -1 {
		s += int(r.ContentLength)
	}
	return int64(s)
}
