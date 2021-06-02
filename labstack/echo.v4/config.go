package echo

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/golang.org/x/net.v0/http2/server"
)

const (
	root         = "ignite.echo"
	hideBanner   = ".hideBanner"
	disableHTTP2 = ".disableHTTP2"
	tp           = ".type"
	port         = ".port"
	protocol     = ".protocol"
	hc2Root      = ".hc2"
	tlsRoot      = ".tls"
	tlsEnabled   = tlsRoot + ".enabled"
	tlsType      = tlsRoot + ".type"
	tlsAutoRoot  = tlsRoot + ".auto"
	tlsAutoHost  = tlsAutoRoot + ".host"
	tlsFileRoot  = tlsRoot + ".file"
	tlsFileCert  = tlsFileRoot + ".cert"
	tlsFileKey   = tlsFileRoot + ".key"
	PluginsRoot  = root + ".plugins"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	server.ConfigAdd(path + hc2Root)
	config.Add(path+hideBanner, true, "hide/show banner")
	config.Add(path+disableHTTP2, false, "disable http2")
	config.Add(path+port, 8080, "Server http port")
	config.Add(path+tp, "REST", "defines type for applicaton")
	config.Add(path+protocol, "HTTP", "defines protocol HTTP/H2C")
	config.Add(path+tlsEnabled, false, "enable/disable tls")
	config.Add(path+tlsType, "AUTO", "defines tls type. AUTO/FILE")
	config.Add(path+tlsAutoHost, "localhost", "defines tls auto host")
	config.Add(path+tlsFileCert, "", "defines tls cert")
	config.Add(path+tlsFileKey, "", "defines tls key")
}
