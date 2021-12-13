package cors

import (
	"net/http"

	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/http/server/go-chi/chi.v5"
)

const (
	root               = chi.PluginsRoot + ".cors"
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
	config.Add(path+allowedOrigins, []string{"*"}, "cors allow origins")
	config.Add(path+allowedHeaders, []string{"Origin", "Content-Type", "Accept"},
		"cors allow headers")
	config.Add(path+allowedMethods,
		[]string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		"cors allow methods")
	config.Add(path+allowedCredentials, true, "cors allow credentials")
	config.Add(path+exposedHeaders, []string{}, "cors expose headers")
	config.Add(path+maxAge, 5200, "cors max age (seconds)")
}
