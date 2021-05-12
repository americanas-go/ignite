package gzip

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/labstack/echo.v4"
)

const (
	enabled = echo.PluginsRoot + ".gzip.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable gzip middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
