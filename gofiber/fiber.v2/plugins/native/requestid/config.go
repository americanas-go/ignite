package requestid

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/gofiber/fiber.v2"
)

const (
	enabled = fiber.PluginsRoot + ".requestid.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable requestid middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
