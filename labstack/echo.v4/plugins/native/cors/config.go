package cors

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/labstack/echo.v4"
)

const (
	root               = echo.PluginsRoot + ".cors"
	enabled            = ".enabled"
	allowedRoot        = ".allowed"
	allowedOrigins     = allowedRoot + ".origins"
	allowedHeaders     = allowedRoot + ".headers"
	allowedMethods     = allowedRoot + ".methods"
	allowedCredentials = allowedRoot + ".credentials"
	exposedHeaders     = ".exposed.headers"
	maxAge             = ".maxAge"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable cors middleware")
	config.Add(path+allowedOrigins, []string{"*"}, "cors allowed origins")
	config.Add(path+allowedHeaders, []string{"*"},
		"cors allowed headers")
	config.Add(path+allowedMethods,
		[]string{"*"},
		"cors allowed methods")
	config.Add(path+allowedCredentials, true, "cors allowed credentials")
	config.Add(path+exposedHeaders, []string{}, "cors exposed headers")
	config.Add(path+maxAge, 5200, "cors max age (seconds)")
}
