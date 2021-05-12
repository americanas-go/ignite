package compressor

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/google.golang.org/grpc.v1/client"
)

const (
	root  = client.PluginsRoot + ".compressor"
	level = root + ".level"
)

func init() {
	config.Add(level, -1, "sets gzip level")
}

func Level() int {
	return config.Int(level)
}
