package health

import (
	"github.com/americanas-go/config"
	ginats "github.com/jvitoroc/ignite/nats-io/nats.go.v1"
)

const (
	root        = ginats.ExtRoot + ".health"
	name        = root + ".name"
	description = root + ".description"
	required    = root + ".required"
	enabled     = root + ".enabled"
)

func init() {
	config.Add(name, "nats", "health name")
	config.Add(description, "default connection", "define health description")
	config.Add(required, true, "define health description")
	config.Add(enabled, true, "enable/disable health")
}
