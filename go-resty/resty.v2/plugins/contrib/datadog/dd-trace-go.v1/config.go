package datadog

import (
	"github.com/americanas-go/config"
	girest "github.com/americanas-go/ignite/go-resty/resty.v2"
)

const (
	ConfigRoot = girest.PluginsRoot + ".datadog"
	enabled    = ConfigRoot + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable datadog integration")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
