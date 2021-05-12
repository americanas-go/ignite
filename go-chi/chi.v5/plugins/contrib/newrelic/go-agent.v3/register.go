package newrelic

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/americanas-go/log"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jvitoroc/ignite/go-chi/chi.v5"
	newrelic "github.com/jvitoroc/ignite/newrelic/go-agent.v3"
	nr "github.com/newrelic/go-agent/v3/newrelic"
)

func Register(ctx context.Context) (*chi.Config, error) {
	if !IsEnabled() || !newrelic.IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling newrelic middleware in chi")

	return &chi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			nrMiddleware,
		},
	}, nil
}

func nrMiddleware(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		url := r.URL.String()
		path := r.URL.Path
		txnName := strings.Join([]string{r.Method, path}, " ")

		txn := newrelic.Application().StartTransaction(txnName)
		defer txn.End()

		txn.SetWebRequestHTTP(r)

		if isWebResponseEnabled() {
			w = txn.SetWebResponse(w)
		}

		txn.AddAttribute("request.url", fmt.Sprintf("http://%s%s", r.Host, url))

		qs := r.URL.Query()
		for key, value := range qs {
			txn.AddAttribute(key, strings.Join(value, "|"))
		}

		if reqID := middleware.GetReqID(ctx); reqID != "" {
			txn.AddAttribute("request.id", reqID)
		}

		r = nr.RequestWithTransactionContext(r, txn)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
