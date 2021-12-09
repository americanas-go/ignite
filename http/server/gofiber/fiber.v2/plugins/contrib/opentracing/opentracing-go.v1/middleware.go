package opentracing

import (
	"fmt"
	"net/http"
	"net/url"
	"unsafe"

	"github.com/gofiber/fiber/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

type contextKey struct{}

var activeSpanKey = contextKey{}

type mwOptions struct {
	opNameFunc    func(c *fiber.Ctx) string
	spanObserver  func(span opentracing.Span, c *fiber.Ctx)
	urlTagFunc    func(u *url.URL) string
	componentName string
}

var getString = func(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func opentracingMiddleware() fiber.Handler {

	// Return new handler
	return func(c *fiber.Ctx) error {

		tracer := opentracing.GlobalTracer()

		opts := mwOptions{
			componentName: "fiber",
			opNameFunc: func(c *fiber.Ctx) string {
				return "HTTP " + c.Method() + " " + c.Path()
			},
			spanObserver: func(span opentracing.Span, c *fiber.Ctx) {

			},
			urlTagFunc: func(u *url.URL) string {
				return u.String()
			},
		}

		hdr := make(http.Header)

		c.Request().Header.VisitAll(func(k, v []byte) {
			hdr.Set(getString(k), getString(v))
		})

		carrier := opentracing.HTTPHeadersCarrier(hdr)
		ctx, _ := tracer.Extract(opentracing.HTTPHeaders, carrier)
		op := opts.opNameFunc(c)
		sp := opentracing.StartSpan(op, ext.RPCServerOption(ctx))
		defer sp.Finish()

		ext.HTTPMethod.Set(sp, c.Method())

		u, _ := url.ParseRequestURI(c.OriginalURL())

		ext.HTTPUrl.Set(sp, opts.urlTagFunc(u))
		opts.spanObserver(sp, c)
		ext.Component.Set(sp, opts.componentName)

		c.Context().SetUserValue(fmt.Sprintf("%v", activeSpanKey), sp)

		err := tracer.Inject(sp.Context(), opentracing.HTTPHeaders, carrier)

		if err != nil {
			panic("SpanContext Inject Error!")
		}

		if err := c.Next(); err != nil {
			sp.SetTag("error", true)
			return err
		}

		sp.SetTag("error", false)
		ext.HTTPStatusCode.Set(sp, uint16(c.Response().StatusCode()))

		return nil

	}
}
