package etag

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/gofiber/fiber.v2"
)

const (
	enabled = fiber.PluginsRoot + ".etag.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable etag middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
