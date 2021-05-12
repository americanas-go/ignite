package bodydump

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/labstack/echo.v4"
)

const (
	enabled = echo.PluginsRoot + ".bodydump.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable body dump middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
