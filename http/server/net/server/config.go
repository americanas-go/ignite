package server

import (
	"time"

	"github.com/americanas-go/config"
)

const (
	root              = "ignite.http.server"
	serverAddress     = ".addr"
	maxHeaderBytes    = ".maxHeaderBytes"
	readHeaderTimeout = ".readHeaderTimeout"
	readTimeout       = ".readTimeout"
	writeTimeout      = ".writeTimeout"
	idleTimeout       = ".idleTimeout"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+serverAddress, ":8081", "server address")
	config.Add(path+maxHeaderBytes, 1048576, "max header timeout")
	config.Add(path+readHeaderTimeout, 1*time.Second, "read header timeout")
	config.Add(path+readTimeout, 1*time.Second, "read timeout")
	config.Add(path+writeTimeout, 7*time.Second, "write timeout ")
	config.Add(path+idleTimeout, 30*time.Second, "idle timeout")
}
