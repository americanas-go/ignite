package newrelic

import (
	"context"
	"fmt"
	c "github.com/go-chi/chi/v5"
	"net/http"
	"strings"

	"github.com/americanas-go/ignite/go-chi/chi.v5"
	newrelic "github.com/americanas-go/ignite/newrelic/go-agent.v3"
	"github.com/americanas-go/log"
	"github.com/go-chi/chi/v5/middleware"
	nr "github.com/newrelic/go-agent/v3/newrelic"
)

// Register registers a default newrelic plugin as a middleware in a new chi config.
func Register(ctx context.Context, mux *c.Mux) (*chi.Config, error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}
	n := NewNewrelicWithOptions(o)
	return n.Register(ctx, mux)
}

// NewRelic struct which represents a new relic plugin for chi
type Newrelic struct {
	options *Options
}

// NewNewrelicWithConfigPath returns a new relic plugin with options from config path.
func NewNewrelicWithConfigPath(path string) (*Newrelic, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewNewrelicWithOptions(o), nil
}

// NewNewrelicWithOptions returns a new relic plugin with options.
func NewNewrelicWithOptions(options *Options) *Newrelic {
	return &Newrelic{options: options}
}

// Register registers the newrelic plugin as a middleware in a new chi config.
func (d *Newrelic) Register(ctx context.Context, mux *c.Mux) (*chi.Config, error) {
	if !d.options.Enabled || !newrelic.IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling newrelic middleware in chi")

	return &chi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			d.nrMiddleware,
		},
	}, nil
}

func (d *Newrelic) nrMiddleware(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		url := r.URL.String()
		path := r.URL.Path
		txnName := strings.Join([]string{r.Method, path}, " ")

		txn := newrelic.Application().StartTransaction(txnName)
		defer txn.End()

		txn.SetWebRequestHTTP(r)

		if d.options.WebResponseEnabled {
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
