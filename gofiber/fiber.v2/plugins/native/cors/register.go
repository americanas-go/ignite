package cors

import (
	"context"
	"strings"

	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	f "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {

	if !IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling cors middleware in fiber")

	return nil, func(ctx context.Context, app *f.App) error {
		app.Use(cors.New(cors.Config{
			AllowOrigins:     strings.Join(getAllowOrigins(), ","),
			AllowMethods:     strings.Join(getAllowMethods(), ","),
			AllowHeaders:     strings.Join(getAllowHeaders(), ","),
			AllowCredentials: getAllowCredentials(),
			ExposeHeaders:    strings.Join(getExposeHeaders(), ","),
			MaxAge:           getMaxAge(),
		}))

		logger.Debug("cors middleware successfully enabled in fiber")

		return nil
	}
}
