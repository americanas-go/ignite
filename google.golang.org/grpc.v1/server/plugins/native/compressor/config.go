package compressor

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/google.golang.org/grpc.v1/server"
)

const (
	root  = server.PluginsRoot + ".compressor"
	level = root + ".level"
)

func init() {
	config.Add(level, -1, "sets gzip level")
}

func Level() int {
	return config.Int(level)
}
