package prometheus

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/gofiber/fiber.v2"
)

const (
	ConfigRoot = fiber.PluginsRoot + ".prometheus"
	enabled    = ConfigRoot + ".enabled"
	route      = ConfigRoot + ".route"
)

func init() {
	config.Add(enabled, true, "enable/disable prometheus integration")
	config.Add(route, "/metrics", "define prometheus metrics url")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func getRoute() string {
	return config.String(route)
}
