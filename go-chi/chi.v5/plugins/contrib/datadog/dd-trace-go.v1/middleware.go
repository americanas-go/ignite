package datadog

import (
	"fmt"
	"net/http"
	"strconv"

	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	"github.com/americanas-go/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// Middleware returns middleware that will trace incoming requests.
func Middleware(spanOpts ...ddtrace.StartSpanOption) func(next http.Handler) http.Handler {
	log.Debug("contrib/go-chi/chi.v5: Configuring Middleware")
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			opts := []ddtrace.StartSpanOption{
				tracer.SpanType(ext.SpanTypeWeb),
				tracer.Tag(ext.HTTPMethod, r.Method),
				tracer.Tag(ext.HTTPURL, r.URL.Path),
				tracer.Measured(),
			}

			if spanctx, err := tracer.Extract(tracer.HTTPHeadersCarrier(r.Header)); err == nil {
				opts = append(opts, tracer.ChildOf(spanctx))
			}

			for h, t := range datadog.HeaderTags() {
				if head := r.Header.Get(h); head != "" {
					opts = append(opts, tracer.Tag(t, head))
				}
			}

			opts = append(opts, spanOpts...)
			span, ctx := tracer.StartSpanFromContext(r.Context(), "http.request", opts...)
			defer span.Finish()

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			// pass the span through the request context and serve the request to the next middleware
			next.ServeHTTP(ww, r.WithContext(ctx))

			// set the resource name as we get it only once the handler is executed
			resourceName := chi.RouteContext(r.Context()).RoutePattern()
			if resourceName == "" {
				resourceName = "unknown"
			}
			resourceName = r.Method + " " + resourceName
			span.SetTag(ext.ResourceName, resourceName)

			// set the status code
			status := ww.Status()
			// 0 status means one has not yet been sent in which case net/http library will write StatusOK
			if ww.Status() == 0 {
				status = http.StatusOK
			}
			span.SetTag(ext.HTTPCode, strconv.Itoa(status))

			if status > 399 {
				span.SetTag(ext.Error, fmt.Errorf("%d: %s", status, http.StatusText(status)))
			}
		})
	}
}
