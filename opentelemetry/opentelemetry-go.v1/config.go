package opentelemetry

import "github.com/americanas-go/config"

const (
	root    = "ignite.otel"
	enabled = root + ".enabled"
	service = root + ".service"
	env     = root + ".env"
	version = root + ".version"
	host    = root + ".host"
	port    = root + ".port"
	tags    = root + ".tags"
)

func init() {
	config.Add(service, "", "service name for opentelemetry spans")
	config.Add(enabled, true, "enables the opentelemetry tracer")
	config.Add(env, "", "service env")
	config.Add(version, "0.0.0", "service version")
	config.Add(host, "localhost", "host address of the opentelemetry agent")
	config.Add(port, "4318", "port of the opentelemetry agent")
	config.Add(tags, map[string]string{}, "sets a key/value pair which will be set as a tag on all spans created by tracer. This option may be used multiple times")
}

// IsTracerEnabled returns config value from key ignite.opentelemetry.enabled where default is true.
func IsTracerEnabled() bool {
	return config.Bool(enabled)
}

// Service returns config value from key ignite.opentelemetry.service where default is empty.
func Service() string {
	return config.String(service)
}
