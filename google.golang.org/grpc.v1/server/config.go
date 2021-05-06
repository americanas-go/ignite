package server

import "github.com/americanas-go/config"

const (
	root                  = "ignite.grpc.server"
	port                  = root + ".port"
	maxConcurrentStreams  = root + ".maxConcurrentStreams"
	initialWindowSize     = root + ".initialWindowSize"
	initialConnWindowSize = root + ".initialConnWindowSize"
	tlsEnabled            = root + ".tls.enabled"
	certFile              = root + ".tls.certFile"
	keyFile               = root + ".tls.keyFile"
	caFile                = root + ".tls.caFile"
	PluginsRoot           = root + ".plugins"
)

func init() {
	config.Add(port, 9090, "server grpc port")
	config.Add(maxConcurrentStreams, 1024*1024*2, "server grpc max concurrent streams")
	config.Add(initialWindowSize, 1024*1024*2, "sets the initial window size for a stream")
	config.Add(initialConnWindowSize, 1024*1024*2, "sets the initial window size for a connection")
	config.Add(tlsEnabled, false, "use TLS - required for HTTP2.")
	config.Add(certFile, "", "path to the CRT/PEM file.")
	config.Add(keyFile, "", "path to the private key file.")
	config.Add(caFile, "", "path to the certificate authority (CA).")
}
