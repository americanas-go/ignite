package opentelemetry

import (
	"context"
	"sync"

	"github.com/americanas-go/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

// StartTracer starts the tracer like StartTraceWithOptions but with default Options.
func StartTracer(ctx context.Context, startOptions ...sdktrace.TracerProviderOption) {

	o, err := NewOptions()
	if err != nil {
		panic(err)
	}

	StartTracerWithOptions(ctx, o, startOptions...)
}

var tracerOnce sync.Once

// StartTracerWithOptions starts the tracer with the given set of options. Calling
// it multiple times will have no effect. If an error occours during tracer initialization,
// a Noop trace provider will be used instead.
func StartTracerWithOptions(ctx context.Context, options *Options, startOptions ...sdktrace.TracerProviderOption) {

	if !IsTracerEnabled() {
		return
	}

	tracerOnce.Do(func() {
		logger := log.FromContext(ctx)

		exporterOpts := []otlptracehttp.Option{otlptracehttp.WithEndpoint(options.Endpoint)}
		if IsInsecure() {
			exporterOpts = append(exporterOpts, otlptracehttp.WithInsecure())
		}

		exporter, err := otlptracehttp.New(
			ctx,
			exporterOpts...,
		)
		if err != nil {
			logger.Error("error creating opentelemetry exporter: ", err)
			otel.SetTracerProvider(trace.NewNoopTracerProvider())
			return
		}

		attrs := make([]attribute.KeyValue, len(options.Tags))
		for k, v := range options.Tags {
			attrs = append(attrs, attribute.KeyValue{
				Key:   attribute.Key(k),
				Value: attribute.StringValue(v),
			})
		}

		rs, err := resource.New(ctx,
			resource.WithSchemaURL(semconv.SchemaURL),
			resource.WithAttributes(
				semconv.ServiceNameKey.String(options.Service),
				semconv.ServiceVersionKey.String(options.Version),
				attribute.String("env", options.Env),
			),
			resource.WithAttributes(attrs...),
		)
		if err != nil {
			logger.Error("error creating opentelemetry resource: ", err)
			otel.SetTracerProvider(trace.NewNoopTracerProvider())
			return
		}

		startOptions = append(startOptions,
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(rs),
		)

		prov := sdktrace.NewTracerProvider(startOptions...)

		otel.SetTracerProvider(prov)

		log.Infof("started opentelemetry tracer: %s", options.Service)
	})
}
