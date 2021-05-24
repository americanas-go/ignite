package semaphore

import (
	"context"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	e "github.com/labstack/echo/v4"
	"golang.org/x/sync/semaphore"
)

var (
	sem *semaphore.Weighted
)

func Register(ctx context.Context, server *echo.Server) error {

	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling semaphore middleware in echo")

	server.Use(semaphoreMiddleware(int64(GetLimit())))

	logger.Debug("semaphore middleware successfully enabled in echo")

	return nil
}

func semaphoreMiddleware(limit int64) e.MiddlewareFunc {

	sem = semaphore.NewWeighted(limit)

	return func(next e.HandlerFunc) e.HandlerFunc {
		return func(c e.Context) error {

			logger := log.FromContext(c.Request().Context())

			if !sem.TryAcquire(1) {
				logger.Errorf("the http server has reached the limit of %v simultaneous connections", limit)
				return e.ErrServiceUnavailable
			}
			defer sem.Release(1)

			return next(c)
		}
	}
}
