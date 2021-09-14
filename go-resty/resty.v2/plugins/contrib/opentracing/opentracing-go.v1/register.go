package opentracing

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/go-resty/resty/v2"
	"github.com/opentracing/opentracing-go"
)

// Register registers opentracing middleware for resty.
func Register(ctx context.Context, client *resty.Client) error {

	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling opentracing middleware in resty")

	client.OnBeforeRequest(ot)

	logger.Debug("opentracing middleware successfully enabled in resty")

	return nil
}

func ot(client *resty.Client, request *resty.Request) error {

	ctx := request.Context()

	if span := opentracing.SpanFromContext(ctx); span != nil {
		return opentracing.GlobalTracer().Inject(
			span.Context(),
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(request.Header))
	}

	return nil
}
