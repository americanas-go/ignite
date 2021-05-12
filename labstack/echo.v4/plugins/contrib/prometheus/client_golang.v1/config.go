package prometheus

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/labstack/echo.v4"
)

const (
	ConfigRoot = echo.PluginsRoot + ".prometheus"
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

func GetRoute() string {
	return config.String(route)
}
