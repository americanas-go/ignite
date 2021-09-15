package datadog

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/aws/aws-sdk-go.v2"
)

const (
	root    = aws.PluginsRoot + ".datadog"
	enabled = ".enabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(enabled, true, "enable/disable datadog integration")
}
