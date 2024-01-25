package contrib // import "github.com/americanas-go/ignite/go.mongodb.org/mongo-driver.v1/plugins/contrib/opentelemetry/otelmongo.v1

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/go.mongodb.org/mongo-driver.v1"
)

const (
	root    = mongo.PluginsRoot + ".otel"
	enabled = ".enabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable the opentelemetry integration")
}
