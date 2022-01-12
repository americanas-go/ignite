package contrib

import (
	"context"
	"net/http"

	iresty "github.com/americanas-go/ignite/go-resty/resty.v2/resty"
	newrelic "github.com/americanas-go/ignite/newrelic/go-agent.v3"
	"github.com/americanas-go/log"
	"github.com/go-resty/resty/v2"
	nr "github.com/newrelic/go-agent/v3/newrelic"
)

// Adds Newrelic apm support.
func Newrelic(ctx context.Context, w *iresty.Wrapper) error {
	o := w.Options.Plugins.Newrelic
	if !o.Enabled || !newrelic.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating resty in newrelic")

	w.Instance.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {

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

	w.Instance.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {

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
