package pkg_otel

import (
	"context"
	"nebeng-jek/pkg/logger"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

func NewOpenTelemetry(serviceHost, serviceName, appEnv string) *OpenTelemetry {
	var (
		attributeName = semconv.ServiceNameKey.String(serviceName)
		ctx           = context.Background()
	)

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

	exp, err := initGRPCExporters(ctx, serviceHost)
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
	}

	err = initTracerProvider(res, exp.TraceExporter)
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
	}
	err = initMeterProvider(res, exp.MetricExporter)
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

// Initializes an OTLP exporter, and configures the corresponding trace provider.
func initTracerProvider(res *resource.Resource, exporter sdktrace.SpanExporter) error {
	// Register the trace exporter with a TracerProvider, using a batch
	// span processor to aggregate spans before export.
	bsp := sdktrace.NewBatchSpanProcessor(exporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)

	// Set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// Shutdown will flush any remaining spans and shut down the exporter.
	return nil
}

// Initializes an OTLP exporter, and configures the corresponding meter provider.
func initMeterProvider(res *resource.Resource, exporter sdkmetric.Exporter) error {
	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(
			sdkmetric.NewPeriodicReader(
				exporter,
				// Default is 1m. Set to 3s for demonstrative purposes.
				sdkmetric.WithInterval(3*time.Second),
			),
		),
		sdkmetric.WithResource(res),
	)
	otel.SetMeterProvider(meterProvider)

	return nil
}

// EndAPM shutdown the tracer
func (o *OpenTelemetry) EndAPM() error {
	if tp, ok := otel.GetTracerProvider().(*sdktrace.TracerProvider); ok {
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
