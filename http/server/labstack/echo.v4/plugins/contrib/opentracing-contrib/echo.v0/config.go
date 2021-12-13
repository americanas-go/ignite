package opentracing

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/http/server/labstack/echo.v4"
)

const (
	root    = echo.PluginsRoot + ".opentracing"
	enabled = ".enabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable opentracing")
}
