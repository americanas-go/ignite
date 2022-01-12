package contrib

import (
	"context"
	"strconv"

	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	iresty "github.com/americanas-go/ignite/go-resty/resty.v2/resty"
	"github.com/americanas-go/log"
	"github.com/go-resty/resty/v2"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// Adds Datadog APM support.
func Datadog(ctx context.Context, w *iresty.Wrapper) error {
	o := w.Options.Plugins.Datadog
	if !o.Enabled || !datadog.IsTracerEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating resty in datadog")

	w.Instance.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
		spanOptions := []ddtrace.StartSpanOption{
			tracer.ResourceName(request.URL),
			tracer.SpanType(ext.SpanTypeHTTP),
			tracer.Tag(ext.HTTPMethod, request.Method),
			tracer.Tag(ext.HTTPURL, request.URL),
		}
		//add custom tags
		for k, v := range o.Tags {
			spanOptions = append(spanOptions, tracer.Tag(k, v))
		}
		reqCtx := request.Context()
		span, ctx := tracer.StartSpanFromContext(reqCtx, o.OperationName, spanOptions...)

		// pass the span through the request context
		request.SetContext(ctx)

		return tracer.Inject(span.Context(), tracer.HTTPHeadersCarrier(request.Header))
	})

	w.Instance.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
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
