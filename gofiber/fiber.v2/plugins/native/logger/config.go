package logger

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/gofiber/fiber.v2"
)

const (
	enabled = fiber.PluginsRoot + ".logger.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable logger middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
