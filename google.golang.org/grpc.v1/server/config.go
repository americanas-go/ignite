package server

import "github.com/americanas-go/config"

const (
	root                  = "ignite.grpc.server"
	port                  = ".port"
	maxConcurrentStreams  = ".maxConcurrentStreams"
	initialWindowSize     = ".initialWindowSize"
	initialConnWindowSize = ".initialConnWindowSize"
	tlsRoot               = ".tls"
	tlsEnabled            = tlsRoot + ".enabled"
	tlsType               = tlsRoot + ".type"
	tlsAutoRoot           = tlsRoot + ".auto"
	tlsAutoHost           = tlsAutoRoot + ".host"
	tlsFileRoot           = tlsRoot + ".file"
	tlsFileCert           = tlsFileRoot + ".cert"
	tlsFileKey            = tlsFileRoot + ".key"
	tlsFileCA             = tlsFileRoot + ".ca"
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
	config.Add(path+tlsEnabled, false, "use TLS - required for HTTP2")
	config.Add(path+tlsType, "AUTO", "defines tls type. AUTO/FILE")
	config.Add(path+tlsAutoHost, "localhost", "defines tls auto host")
	config.Add(path+tlsFileCert, "", "path to the CRT/PEM file")
	config.Add(path+tlsFileKey, "", "path to the private key file")
	config.Add(path+tlsFileCA, "", "path to the certificate authority (CA)")
}
