package server

import "github.com/americanas-go/config"

const (
	root                  = "ignite.grpc.server"
	port                  = ".port"
	maxConcurrentStreams  = ".maxConcurrentStreams"
	initialWindowSize     = ".initialWindowSize"
	initialConnWindowSize = ".initialConnWindowSize"
	tlsEnabled            = ".tls.enabled"
	certFile              = ".tls.certFile"
	keyFile               = ".tls.keyFile"
	caFile                = ".tls.caFile"
	PluginsRoot           = root + ".plugins"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+port, 9090, "server grpc port")
	config.Add(path+maxConcurrentStreams, 1024*1024*2, "server grpc max concurrent streams")
	config.Add(path+initialWindowSize, 1024*1024*2, "sets the initial window size for a stream")
	config.Add(path+initialConnWindowSize, 1024*1024*2, "sets the initial window size for a connection")
	config.Add(path+tlsEnabled, false, "use TLS - required for HTTP2.")
	config.Add(path+certFile, "", "path to the CRT/PEM file.")
	config.Add(path+keyFile, "", "path to the private key file.")
	config.Add(path+caFile, "", "path to the certificate authority (CA).")
}
