package errorhandler

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/labstack/echo.v4"
)

const (
	enabled = echo.PluginsRoot + ".errorhandler.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable custom error handler")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
