package datadog

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/labstack/echo.v4"
)

const (
	enabled = echo.PluginsRoot + ".datadog.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable datadog middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
