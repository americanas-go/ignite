package otelecho // import "github.com/americanas-go/ignite/labstack/echo.v4/plugins/contrib/opentelemetry/otelecho.v1"

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/labstack/echo.v4"
)

const (
	root    = echo.PluginsRoot + ".otel"
	enabled = ".enabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable the opentelemetry integration")
}
