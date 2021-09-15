package compressor

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/google.golang.org/grpc.v1/client"
)

const (
	root  = client.PluginsRoot + ".compressor"
	level = ".level"
)

func ConfigAdd(path string) {
	config.Add(path+level, -1, "sets gzip level")
}
