package server

import (
	"time"

	"github.com/americanas-go/config"
)

const (
	root                         = "ignite.http2.server"
	maxHandlers                  = ".maxHandlers"
	maxConcurrentStreams         = ".maxConcurrentStreams"
	maxReadFrameSize             = ".maxReadFrameSize"
	permitProhibitedCipherSuites = ".permitProhibitedCipherSuites"
	maxUploadBufferPerConnection = ".maxUploadBufferPerConnection"
	maxUploadBufferPerStream     = ".maxUploadBufferPerStream"
	idleTimeout                  = ".idleTimeout"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+maxHandlers, 2, "limits the number of http.Handler ServeHTTP goroutines which may run at a time over all connections")
	config.Add(path+maxConcurrentStreams, 250, "optionally specifies the number of concurrent streams that each client may have open at a time")
	config.Add(path+maxReadFrameSize, 1<<20, "optionally specifies the largest frame this server is willing to read. A valid value is between 16k and 16M, inclusive")
	config.Add(path+permitProhibitedCipherSuites, false, "if true, permits the use of cipher suites prohibited by the HTTP/2 spec")
	config.Add(path+idleTimeout, 30*time.Second, "specifies how long until idle clients should be closed with a GOAWAY frame")
	config.Add(path+maxUploadBufferPerConnection, 1<<20, "is the size of the initial flow control window for each connections")
	config.Add(path+maxUploadBufferPerStream, 1<<20, "is the size of the initial flow control window for each stream")
}
