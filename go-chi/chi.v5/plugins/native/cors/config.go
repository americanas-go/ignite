package cors

import (
	"net/http"

	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/go-chi/chi.v5"
)

const (
	root               = chi.PluginsRoot + ".cors"
	enabled            = root + ".enabled"
	allowedOrigins     = root + ".allowed.origins"
	allowedHeaders     = root + ".allowed.headers"
	allowedMethods     = root + ".allowed.methods"
	allowedCredentials = root + ".allowed.credentials"
	exposedHeaders     = root + ".exposed.headers"
	maxAge             = root + ".maxage"
)

func init() {
	config.Add(enabled, true, "enable/disable cors middleware")
	config.Add(allowedOrigins, []string{"*"}, "cors allow origins")
	config.Add(allowedHeaders, []string{"Origin", "Content-Type", "Accept"},
		"cors allow headers")
	config.Add(allowedMethods,
		[]string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		"cors allow methods")
	config.Add(allowedCredentials, true, "cors allow credentials")
	config.Add(exposedHeaders, []string{}, "cors expose headers")
	config.Add(maxAge, 5200, "cors max age (seconds)")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func getAllowedOrigins() []string {
	return config.Strings(allowedOrigins)
}

func getAllowedMethods() []string {
	return config.Strings(allowedMethods)
}

func getAllowedHeaders() []string {
	return config.Strings(allowedHeaders)
}

func getAllowedCredentials() bool {
	return config.Bool(allowedCredentials)
}

func getExposedHeaders() []string {
	return config.Strings(exposedHeaders)
}

func getMaxAge() int {
	return config.Int(maxAge)
}
