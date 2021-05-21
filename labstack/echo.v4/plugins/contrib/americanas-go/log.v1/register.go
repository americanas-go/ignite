package log

import (
	"context"
	"strconv"
	"time"

	"github.com/americanas-go/log"
	"github.com/labstack/echo/v4"
)

func Register(ctx context.Context, instance *echo.Echo) error {
	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling logger middleware in echo")

	instance.Use(loggerMiddleware(Level()))

	logger.Debug("logger middleware successfully enabled in echo")

	return nil
}

// loggerMiddleware returns a middleware that logs HTTP requests.
func loggerMiddleware(level string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			start := time.Now()

			ctx := req.Context()

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			logger := log.FromContext(ctx).
				WithField("requestId", id).
				WithField("requestUri", req.RequestURI)

			ctx = logger.ToContext(ctx)
			c.SetRequest(req.WithContext(ctx))

			var err error
			if err = next(c); err != nil {
				c.Error(err)
			}

			stop := time.Now()

			reqSize := req.Header.Get(echo.HeaderContentLength)
			if reqSize == "" {
				reqSize = "0"
			}

			var method func(format string, args ...interface{})

			switch level {
			case "TRACE":
				method = logger.Tracef
			case "INFO":
				method = logger.Infof
			default:
				method = logger.Debugf
			}

			method("%s %s %s %-7s %s %3d %s %s %13v %s %s",
				id,
				c.RealIP(),
				req.Host,
				req.Method,
				req.RequestURI,
				res.Status,
				reqSize,
				strconv.FormatInt(res.Size, 10),
				stop.Sub(start).String(),
				req.Referer(),
				req.UserAgent(),
			)

			return err
		}
	}
}
