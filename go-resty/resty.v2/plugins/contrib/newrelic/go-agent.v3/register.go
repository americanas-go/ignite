package newrelic

import (
	"context"

	"net/http"

	"github.com/americanas-go/log"
	"github.com/go-resty/resty/v2"
	newrelic "github.com/jvitoroc/ignite/newrelic/go-agent.v3"
	nr "github.com/newrelic/go-agent/v3/newrelic"
)

func Register(ctx context.Context, client *resty.Client) error {

	if !IsEnabled() || !newrelic.IsEnabled() {
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
