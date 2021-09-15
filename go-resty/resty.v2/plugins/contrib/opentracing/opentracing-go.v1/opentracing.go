package opentracing

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/go-resty/resty/v2"
	"github.com/opentracing/opentracing-go"
)

type Opentracing struct {
	options *Options
}

func NewOpentracingWithConfigPath(path string) (*Opentracing, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewOpentracingWithOptions(o), nil
}

func NewOpentracingWithOptions(options *Options) *Opentracing {
	return &Opentracing{options: options}
}

func Register(ctx context.Context, client *resty.Client) error {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	plugin := NewOpentracingWithOptions(o)
	return plugin.Register(ctx, client)
}

func (i *Opentracing) Register(ctx context.Context, client *resty.Client) error {

	if i.options.Enabled {
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
