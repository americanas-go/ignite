package semaphore

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/labstack/echo/v4"
	"golang.org/x/sync/semaphore"
)

var (
	sem *semaphore.Weighted
)

func Register(ctx context.Context, instance *echo.Echo) error {

	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling semaphore middleware in echo")

	instance.Use(semaphoreMiddleware(int64(GetLimit())))

	logger.Debug("semaphore middleware successfully enabled in echo")

	return nil
}

func semaphoreMiddleware(limit int64) echo.MiddlewareFunc {

	sem = semaphore.NewWeighted(limit)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			logger := log.FromContext(c.Request().Context())

			if !sem.TryAcquire(1) {
				logger.Errorf("the http server has reached the limit of %v simultaneous connections", limit)
				return echo.ErrServiceUnavailable
			}
			defer sem.Release(1)

			return next(c)
		}
	}
}
