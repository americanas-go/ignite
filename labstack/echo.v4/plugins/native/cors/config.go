package cors

import (
	"net/http"

	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/labstack/echo.v4"
	e "github.com/labstack/echo/v4"
)

const (
	root             = echo.PluginsRoot + ".cors"
	enabled          = root + ".enabled"
	allowOrigins     = root + ".allow.origins"
	allowHeaders     = root + ".allow.headers"
	allowMethods     = root + ".allow.methods"
	allowCredentials = root + ".allow.credentials"
	exposeHeaders    = root + ".expose.headers"
	maxAge           = root + ".maxage"
)

func init() {
	config.Add(enabled, true, "enable/disable cors middleware")
	config.Add(allowOrigins, []string{"*"}, "cors allow origins")
	config.Add(allowHeaders, []string{e.HeaderOrigin, e.HeaderContentType, e.HeaderAccept},
		"cors allow headers")
	config.Add(allowMethods,
		[]string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		"cors allow methods")
	config.Add(allowCredentials, true, "cors allow credentials")
	config.Add(exposeHeaders, []string{}, "cors expose headers")
	config.Add(maxAge, 5200, "cors max age (seconds)")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func GetAllowOrigins() []string {
	return config.Strings(allowOrigins)
}

func GetAllowMethods() []string {
	return config.Strings(allowMethods)
}

func GetAllowHeaders() []string {
	return config.Strings(allowHeaders)
}

func GetAllowCredentials() bool {
	return config.Bool(allowCredentials)
}

func GetExposeHeaders() []string {
	return config.Strings(exposeHeaders)
}

func GetMaxAge() int {
	return config.Int(maxAge)
}
