package health

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/http/server/labstack/echo.v4"
)

const (
	root    = echo.PluginsRoot + ".health"
	enabled = ".enabled"
	route   = ".route"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable health route")
	config.Add(path+route, "/health", "define status url")
}
