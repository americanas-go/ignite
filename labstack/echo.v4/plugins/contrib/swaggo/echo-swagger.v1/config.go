package swagger

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/labstack/echo.v4"
)

const (
	root    = echo.PluginsRoot + ".swagger"
	enabled = ".enabled"
	route   = ".route"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable swagger integration")
	config.Add(path+route, "/swagger/*", "define swagger metrics url")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func GetRoute() string {
	return config.String(route)
}
