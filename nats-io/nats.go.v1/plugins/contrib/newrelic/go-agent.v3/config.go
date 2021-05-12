package newrelic

import (
	"github.com/americanas-go/config"
	ginats "github.com/jvitoroc/ignite/nats-io/nats.go.v1"
)

const (
	root    = ginats.ExtRoot + ".newrelic"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable newrelic")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
