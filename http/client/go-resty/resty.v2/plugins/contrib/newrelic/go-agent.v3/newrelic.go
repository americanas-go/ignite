package newrelic

import (
	"context"

	"net/http"

	newrelic "github.com/americanas-go/ignite/apm/newrelic/go-agent.v3"
	"github.com/americanas-go/log"
	"github.com/go-resty/resty/v2"
	nr "github.com/newrelic/go-agent/v3/newrelic"
)

// Newrelic represents a newrelic plugin for resty client.
type Newrelic struct {
	options *Options
}

// NewNewrelicWithConfigPath returns a new newrelic plugin with options from config path.
func NewNewrelicWithConfigPath(path string) (*Newrelic, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewNewrelicWithOptions(o), nil
}

// NewNewrelicWithOptions returns a new newrelic plugin with options.
func NewNewrelicWithOptions(options *Options) *Newrelic {
	return &Newrelic{options: options}
}

// Register registers a new newrelic plugin on resty client.
func Register(ctx context.Context, client *resty.Client) error {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	plugin := NewNewrelicWithOptions(o)
	return plugin.Register(ctx, client)
}

// Register registers this newrelic plugin on resty client.
func (i *Newrelic) Register(ctx context.Context, client *resty.Client) error {

	if !i.options.Enabled || !newrelic.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating resty in newrelic")

	client.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {

		rctx := request.Context()

		txn := newrelic.FromContext(rctx)
		if txn == nil {
			logger.Warnf("there is no transaction in context for newrelic")
			return nil
		}

		txn.InsertDistributedTraceHeaders(request.Header)

		req, _ := http.NewRequest(request.Method, client.HostURL, nil)
		req.Header = request.Header

		s := nr.StartExternalSegment(txn, req)
		ctx := context.WithValue(rctx, "nrext", s)

		request.SetContext(ctx)

		return nil
	})

	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {

		ctx := resp.Request.Context()

		s, ok := ctx.Value("nrext").(*nr.ExternalSegment)
		if ok {
			s.End()
		}

		return nil
	})

	logger.Debug("resty successfully integrated in newrelic")

	return nil
}
