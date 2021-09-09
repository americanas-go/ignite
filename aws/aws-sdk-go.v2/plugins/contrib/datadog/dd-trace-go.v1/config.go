package datadog

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/aws/aws-sdk-go.v2"
)

const (
	enabled = aws.PluginsRoot + ".datadog.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable datadog integration")
}

// IsEnabled returns config value from key ignite.aws.plugins.datadog.enabled where default is true.
func IsEnabled() bool {
	return config.Bool(enabled)
}
