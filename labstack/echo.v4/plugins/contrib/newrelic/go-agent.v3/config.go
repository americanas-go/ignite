package newrelic

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/labstack/echo.v4"
)

const (
	root                       = echo.PluginsRoot + ".newrelic"
	enabled                    = root + ".enabled"
	middlewareRoot             = root + ".middleware"
	middlewareRequestIDEnabled = middlewareRoot + ".requestId.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable newrelic integration")
	config.Add(middlewareRequestIDEnabled, true, "enable/disable request id middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func IsEnabledRequestID() bool {
	return config.Bool(middlewareRequestIDEnabled)
}
