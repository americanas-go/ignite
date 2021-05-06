package server

import (
	"time"

	"github.com/americanas-go/config"
)

const (
	root              = "ignite.http.server"
	serverAddress     = root + ".addr"
	maxHeaderBytes    = root + ".maxHeaderBytes"
	readHeaderTimeout = root + ".readHeaderTimeout"
	readTimeout       = root + ".readTimeout"
	writeTimeout      = root + ".writeTimeout"
	idleTimeout       = root + ".idleTimeout"
)

func init() {
	config.Add(serverAddress, ":8081", "server address")
	config.Add(maxHeaderBytes, 1048576, "max header timeout")
	config.Add(readHeaderTimeout, 1*time.Second, "read header timeout")
	config.Add(readTimeout, 1*time.Second, "read timeout")
	config.Add(writeTimeout, 7*time.Second, "write timeout ")
	config.Add(idleTimeout, 30*time.Second, "idle timeout")
}
