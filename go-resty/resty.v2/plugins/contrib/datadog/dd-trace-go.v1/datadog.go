package datadog

import (
	"context"
	"strconv"

	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	"github.com/americanas-go/log"
	"github.com/go-resty/resty/v2"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type DataDog struct {
	options *Options
}

func NewDatadogWithConfigPath(path string) (*DataDog, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewDataDogWithOptions(o), nil
}

func NewDataDogWithOptions(options *Options) *DataDog {
	return &DataDog{options: options}
}

func Register(ctx context.Context, client *resty.Client) error {
	o, err := NewOptions()
	if err != nil {
		return err
	}
	datadog := NewDataDogWithOptions(o)
	return datadog.Register(ctx, client)
}

func (d *DataDog) Register(ctx context.Context, client *resty.Client) error {
	if !d.options.Enabled || !datadog.IsTracerEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating resty in datadog")

	client.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
		spanOptions := []ddtrace.StartSpanOption{
			tracer.ResourceName(request.URL),
			tracer.SpanType(ext.SpanTypeHTTP),
			tracer.Tag(ext.HTTPMethod, request.Method),
			tracer.Tag(ext.HTTPURL, request.URL),
		}

		spanOptions = append(spanOptions, d.options.SpanOptions...)

		reqCtx := request.Context()
		span, ctx := tracer.StartSpanFromContext(reqCtx, d.options.OperationName, spanOptions...)

		// pass the span through the request context
		request.SetContext(ctx)

		return tracer.Inject(span.Context(), tracer.HTTPHeadersCarrier(request.Header))
	})

	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		ctx := resp.Request.Context()

		span, ok := tracer.SpanFromContext(ctx)

		if ok {
			span.SetTag(ext.HTTPCode, strconv.Itoa(resp.StatusCode()))
			span.SetTag(ext.Error, resp.Error())
			span.Finish()
		}

		return nil
	})

	logger.Debug("resty successfully integrated in datadog")

	return nil
}
