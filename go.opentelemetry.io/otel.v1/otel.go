package otel

import (
	"context"
	embeddedm "go.opentelemetry.io/otel/metric/embedded"
	"go.opentelemetry.io/otel/propagation"
	embeddedt "go.opentelemetry.io/otel/trace/embedded"
	"os"
	"sync"

	"github.com/americanas-go/log"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	noopt "go.opentelemetry.io/otel/trace/noop"
	"google.golang.org/grpc/credentials"
)

var tracerProvider embeddedt.TracerProvider
var metricProvider embeddedm.MeterProvider

// StartTracerProvider starts the tracer provider like StartTracerProviderWithOptions but with default Options.
func StartTracerProvider(ctx context.Context, startOptions ...sdktrace.TracerProviderOption) {

	o, err := NewOptions()
	if err != nil {
		panic(err)
	}

	StartTracerProviderWithOptions(ctx, o, startOptions...)
}

var tracerOnce sync.Once

// StartTracerProviderWithOptions starts the tracer provider with the given set of options. Calling
// it multiple times will have no effect. If an error occours during tracer initialization,
// a Noop trace provider will be used instead.
func StartTracerProviderWithOptions(ctx context.Context, options *Options, tracerProviderOptions ...sdktrace.TracerProviderOption) {

	if !IsTracerEnabled() {
		return
	}

	tracerOnce.Do(func() {
		logger := log.FromContext(ctx)

		var exporter *otlptrace.Exporter
		var err error

		switch options.Protocol {
		case "grpc":
			exporter, err = startGRPCTracer(ctx, options)
		case "http":
			exporter, err = startHTTPTracer(ctx, options)
		default:
			exporter, err = startHTTPTracer(ctx, options)
		}

		if err != nil {
			logger.Error("error creating opentelemetry exporter: ", err)
			otel.SetTracerProvider(noopt.NewTracerProvider())
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
			otel.SetTracerProvider(noopt.NewTracerProvider())
			return
		}

		tracerProviderOptions = append(tracerProviderOptions,
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(rs),
		)

		prov := sdktrace.NewTracerProvider(tracerProviderOptions...)

		otel.SetTracerProvider(prov)
		otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
		tracerProvider = prov

		provider := sdkmetric.NewMeterProvider(
			sdkmetric.WithResource(rs),
		)

		otel.SetMeterProvider(provider)

		log.Infof("started opentelemetry tracer: %s", options.Service)
	})
}

func startHTTPTracer(ctx context.Context, options *Options) (*otlptrace.Exporter, error) {
	var exporterOpts []otlptracehttp.Option
	if _, ok := os.LookupEnv("OTEL_EXPORTER_OTLP_ENDPOINT"); !ok { // Only using WithEndpoint when the environment variable is not set
		exporterOpts = append(exporterOpts, otlptracehttp.WithEndpoint(options.Endpoint)) //TODO see https://github.com/open-telemetry/opentelemetry-go/issues/3730
	}

	if IsInsecure() {
		exporterOpts = append(exporterOpts, otlptracehttp.WithInsecure())
	}

	exporter, err := otlptracehttp.New(
		ctx,
		exporterOpts...,
	)
	if err != nil {
		return nil, err
	}

	return exporter, nil
}

func startGRPCTracer(ctx context.Context, options *Options) (*otlptrace.Exporter, error) {
	var exporterOpts []otlptracegrpc.Option
	if _, ok := os.LookupEnv("OTEL_EXPORTER_OTLP_ENDPOINT"); !ok { // Only using WithEndpoint when the environment variable is not set
		exporterOpts = append(exporterOpts, otlptracegrpc.WithEndpoint(options.Endpoint)) //TODO see https://github.com/open-telemetry/opentelemetry-go/issues/3730
	}

	if IsInsecure() {
		exporterOpts = append(exporterOpts, otlptracegrpc.WithInsecure())
	} else {
		creds, err := credentials.NewClientTLSFromFile(options.TLS.Cert, "")
		if err != nil {
			return nil, errors.Wrap(err, "error creating tls credentials")
		}
		exporterOpts = append(exporterOpts, otlptracegrpc.WithTLSCredentials(creds))
	}

	exporter, err := otlptracegrpc.New(
		ctx,
		exporterOpts...,
	)
	if err != nil {
		return nil, err
	}

	return exporter, nil
}
