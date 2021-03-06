package cors

import (
	"net/http"

	"github.com/americanas-go/config"
	gifiber "github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/gofiber/fiber/v2"
)

const (
	root               = gifiber.PluginsRoot + ".cors"
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
	config.Add(path+allowedHeaders, []string{fiber.HeaderOrigin, fiber.HeaderContentType, fiber.HeaderAccept},
		"cors allowed headers")
	config.Add(path+allowedMethods,
		[]string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		"cors allowed methods")
	config.Add(path+allowedCredentials, true, "cors allowed credentials")
	config.Add(path+exposedHeaders, []string{}, "cors exposed headers")
	config.Add(path+maxAge, 5200, "cors max age (seconds)")
}
