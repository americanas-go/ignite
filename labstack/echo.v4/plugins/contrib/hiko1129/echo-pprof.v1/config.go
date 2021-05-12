package echo_pprof_v1

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/labstack/echo.v4"
)

const (
	enabled = echo.PluginsRoot + ".pprof.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable pprof integration")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
