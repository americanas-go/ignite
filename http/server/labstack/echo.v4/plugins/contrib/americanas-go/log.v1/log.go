package log

import (
	"context"
	"strconv"
	"time"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	e "github.com/labstack/echo/v4"
)

// Register registers a new logger plugin for echo server.
func Register(ctx context.Context, server *echo.Server) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewLogWithOptions(o)
	return h.Register(ctx, server)
}

// Log represents logger plugin for echo server.
type Log struct {
	options *Options
}

// NewLogWithOptions returns a new logger plugin with options.
func NewLogWithOptions(options *Options) *Log {
	return &Log{options: options}
}

// NewLogWithConfigPath returns a new logger plugin with options from config path.
func NewLogWithConfigPath(path string) (*Log, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewLogWithOptions(o), nil
}

// NewLog returns a new logger plugin with default options.
func NewLog() *Log {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewLogWithOptions(o)
}

// Register registers this logger plugin for echo server.
func (i *Log) Register(ctx context.Context, server *echo.Server) error {
	if !i.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling logger middleware in echo")

	server.Use(loggerMiddleware(i.options.Level))

	logger.Debug("logger middleware successfully enabled in echo")

	return nil
}

// loggerMiddleware returns a middleware that logs HTTP requests.
func loggerMiddleware(level string) e.MiddlewareFunc {
	return func(next e.HandlerFunc) e.HandlerFunc {
		return func(c e.Context) error {

			req := c.Request()
			res := c.Response()
			start := time.Now()

			ctx := req.Context()

			id := req.Header.Get(e.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(e.HeaderXRequestID)
			}

			logger := log.FromContext(ctx).
				WithField("requestId", id).
				WithField("requestUri", req.RequestURI)

			ctx = logger.ToContext(ctx)
			c.SetRequest(req.WithContext(ctx))

			defer func() {
				stop := time.Now()

				reqSize := req.Header.Get(e.HeaderContentLength)
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
			}()

			var err error
			if err = next(c); err != nil {
				c.Error(err)
			}

			return nil
		}
	}
}
