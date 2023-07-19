package otelresty // import "github.com/americanas-go/ignite/go-resty/resty.v2/plugins/contrib/dubonzi/otelresty.v1"

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/go-resty/resty.v2"
)

const (
	root       = resty.PluginsRoot + ".otel"
	enabled    = ".enabled"
	tracerName = ".tracerName"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable the opentelemetry integration")
	config.Add(path+tracerName, "resty.request", "defines the name of the tracer used to create spans")
}
