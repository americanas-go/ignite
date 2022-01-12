package contrib

import (
	"context"

	iresty "github.com/americanas-go/ignite/go-resty/resty.v2/resty"
	"github.com/americanas-go/log"
	"github.com/go-resty/resty/v2"
	"github.com/opentracing/opentracing-go"
)

// Adds Opentracing support.
func Opentracing(ctx context.Context, w *iresty.Wrapper) error {
	o := w.Options.Plugins.Opentracing
	if !o.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling opentracing middleware in resty")

	w.Instance.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
		if span := opentracing.SpanFromContext(request.Context()); span != nil {
			return opentracing.GlobalTracer().Inject(
				span.Context(),
				opentracing.HTTPHeaders,
				opentracing.HTTPHeadersCarrier(request.Header))
		}

		return nil
	})

	logger.Debug("opentracing middleware successfully enabled in resty")

	return nil
}
