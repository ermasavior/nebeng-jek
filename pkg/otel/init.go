package pkg_otel

import (
	"nebeng-jek/pkg/logger"
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

func NewOpenTelemetry(serviceHost, serviceName, appEnv string) *OpenTelemetry {
	ctx := context.Background()

	conn, err := initConn(serviceHost)
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
	}

	var attributeName = semconv.ServiceNameKey.String(serviceName)

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// The service name used to display traces in backends
			attributeName,
			attribute.String("env", appEnv),
			attribute.String("version", "ver.1"),
			attribute.String("platform", "go"),
		),
	)
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
	}

	err = initTracerProvider(ctx, res, conn)
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
	}
	err = initMeterProvider(ctx, res, conn)
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
	}

	tracer := otel.Tracer(serviceName)
	meter := otel.Meter(serviceName)

	return &OpenTelemetry{
		tracer: tracer,
		Meter:  meter,
	}
}

// EndAPM shutdown the tracer
func (o *OpenTelemetry) EndAPM() error {
	if tp, ok := otel.GetTracerProvider().(*sdkTrace.TracerProvider); ok {
		err := tp.Shutdown(context.Background())
		return err
	}

	// shutdown the meter
	if mp, ok := otel.GetMeterProvider().(*sdkmetric.MeterProvider); ok {
		err := mp.Shutdown(context.Background())
		return err
	}

	return nil
}

// StartTransaction starts a new OpenTelemetry span with the given name from a context.
func (o *OpenTelemetry) StartTransaction(ctx context.Context, name string, attributes ...trace.SpanStartOption) (context.Context, interface{}) {
	ctx, span := o.tracer.Start(ctx, name, attributes...)
	return ctx, span
}

// EndTransaction ends the given OpenTelemetry span.
func (o *OpenTelemetry) EndTransaction(txn interface{}) {
	span := txn.(trace.Span)
	span.End()
}
