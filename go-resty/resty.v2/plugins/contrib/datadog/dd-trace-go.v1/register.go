package datadog

import (
	"context"
	"math"
	"net/url"
	"strconv"

	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	"github.com/americanas-go/log"
	"github.com/go-resty/resty/v2"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func Register(ctx context.Context, client *resty.Client) error {

	if !IsEnabled() || !datadog.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating resty in datadog")

	client.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {

		var extHost string
		if client.HostURL != "" {
			extHost = client.HostURL
		} else {
			u, err := url.Parse(request.URL)
			if err != nil {
				extHost = u.Hostname()
			} else {
				extHost = request.URL
			}
		}

		opts := []ddtrace.StartSpanOption{
			tracer.ResourceName(request.URL),
			tracer.SpanType(ext.SpanTypeHTTP),
			tracer.Tag(ext.HTTPMethod, request.Method),
			tracer.Tag(ext.HTTPURL, request.URL),
			tracer.Tag("external.resource", request.Method+" "+extHost),
			tracer.Measured(),
		}
		if spanctx, err := tracer.Extract(tracer.HTTPHeadersCarrier(request.Header)); err == nil {
			opts = append(opts, tracer.ChildOf(spanctx))
		}
		if anRate := datadog.AnalyticsRate(); !math.IsNaN(anRate) {
			opts = append(opts, tracer.AnalyticsRate(anRate))
		}
		for h, t := range datadog.HeaderTags() {
			if head := request.Header.Get(h); head != "" {
				opts = append(opts, tracer.Tag(t, head))
			}
		}

		_, ctx := tracer.StartSpanFromContext(request.Context(), "http.external.request", opts...)

		// pass the span through the request context
		request.SetContext(ctx)

		return nil
	})

	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {

		ctx := resp.Request.Context()

		span, ok := tracer.SpanFromContext(ctx)
		if ok {
			span.SetTag(ext.HTTPCode, strconv.Itoa(resp.StatusCode()))

			var fnOpts []tracer.FinishOption

			if resp.IsError() {
				if e, ok := resp.Error().(error); ok {
					fnOpts = append(fnOpts, tracer.WithError(e))
				} else {
					span.SetTag(ext.Error, resp.Error())
				}
			}

			span.Finish(fnOpts...)
		}

		return nil
	})

	logger.Debug("resty successfully integrated in datadog")

	return nil
}
