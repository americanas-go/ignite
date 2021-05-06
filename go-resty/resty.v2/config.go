package resty

import (
	"time"

	"github.com/americanas-go/config"
)

const (
	root                           = "ignite.resty"
	debug                          = ".debug"
	closeConnection                = ".closeConnection"
	connectionTimeout              = ".connectionTimeout"
	keepAlive                      = ".keepAlive"
	fallbackDelay                  = ".fallbackDelay"
	requestTimeout                 = ".requestTimeout"
	transportDisableCompression    = ".transport.disableCompression"
	transportDisableKeepAlives     = ".transport.disableKeepAlives"
	transportMaxIdleConnsPerHost   = ".transport.maxIdleConnsPerHost"
	transportResponseHeaderTimeout = ".transport.responseHeaderTimeout"
	transportForceAttemptHTTP2     = ".transport.forceAttemptHTTP2"
	transportMaxIdleConns          = ".transport.maxIdleConns"
	transportMaxConnsPerHost       = ".transport.maxConnsPerHost"
	transportIdleConnTimeout       = ".transport.idleConnTimeout"
	transportTLSHandshakeTimeout   = ".transport.TLSHandshakeTimeout"
	transportExpectContinueTimeout = ".transport.expectContinueTimeout"
	PluginsRoot                    = root + ".plugins"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+debug, false, "defines global debug request")
	config.Add(path+closeConnection, false, "defines global http close connection")
	config.Add(path+connectionTimeout, 3*time.Minute, "defines global http connection timeout")
	config.Add(path+keepAlive, 30*time.Second, "defines global http keepalive")
	config.Add(path+fallbackDelay, 300*time.Millisecond, "defines global fallbackDelay")
	config.Add(path+requestTimeout, 30*time.Second, "defines global http request timeout")
	config.Add(path+transportDisableCompression, false, "enabled/disable transport compression")
	config.Add(path+transportDisableKeepAlives, false, "enabled/disable transport keep alives")
	config.Add(path+transportMaxIdleConnsPerHost, 2, "define transport max idle conns per host")
	config.Add(path+transportResponseHeaderTimeout, 2*time.Second, "define transport response header timeout")
	config.Add(path+transportForceAttemptHTTP2, true, "define transport force attempt http2")
	config.Add(path+transportMaxIdleConns, 100, "define transport max idle conns")
	config.Add(path+transportMaxConnsPerHost, 100, "define transport max conns per host")
	config.Add(path+transportIdleConnTimeout, 90*time.Second, "define transport idle conn timeout")
	config.Add(path+transportTLSHandshakeTimeout, 10*time.Second, "define transport TLS handshake timeout")
	config.Add(path+transportExpectContinueTimeout, 1*time.Second, "define transport expect continue timeout")
}
