package semaphore

import (
	"context"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	e "github.com/labstack/echo/v4"
	"golang.org/x/sync/semaphore"
)

func Register(ctx context.Context, server *echo.Server) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewSemaphoreWithOptions(o)
	return h.Register(ctx, server)
}

type Semaphore struct {
	options *Options
	sem     *semaphore.Weighted
}

func NewSemaphoreWithOptions(options *Options) *Semaphore {
	return &Semaphore{options: options}
}

func NewSemaphoreWithConfigPath(path string) (*Semaphore, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewSemaphoreWithOptions(o), nil
}

func NewSemaphore() *Semaphore {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewSemaphoreWithOptions(o)
}

func (i *Semaphore) Register(ctx context.Context, server *echo.Server) error {

	if !i.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling semaphore middleware in echo")

	server.Use(i.semaphoreMiddleware(i.options.Limit))

	logger.Debug("semaphore middleware successfully enabled in echo")

	return nil
}

func (i *Semaphore) semaphoreMiddleware(limit int64) e.MiddlewareFunc {

	i.sem = semaphore.NewWeighted(limit)

	return func(next e.HandlerFunc) e.HandlerFunc {
		return func(c e.Context) error {

			logger := log.FromContext(c.Request().Context())

			if !i.sem.TryAcquire(1) {
				logger.Errorf("the http server has reached the limit of %v simultaneous connections", limit)
				return e.ErrServiceUnavailable
			}
			defer i.sem.Release(1)

			return next(c)
		}
	}
}
