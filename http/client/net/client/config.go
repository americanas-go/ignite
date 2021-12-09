package client

import (
	"time"

	"github.com/americanas-go/config"
)

const (
	root                  = "ignite.http.client"
	maxIdleConnPerHost    = ".maxIdleConnPerHost"
	maxIdleConn           = ".maxIdleConn"
	maxConnsPerHost       = ".maxConnsPerHost"
	idleConnTimeout       = ".idleConnTimeout"
	disableKeepAlives     = ".disableKeepAlives"
	disableCompression    = ".disableCompression"
	forceHTTP2            = ".forceHTTP2"
	tlsHandshakeTimeout   = ".TLSHandshakeTimeout"
	timeout               = ".timeout"
	dialTimeout           = ".dialTimeout"
	keepAlive             = ".keepAlive"
	expectContinueTimeout = ".expectContinueTimeout"
	dualStack             = ".dualStack"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+maxIdleConnPerHost, 1, "http max idle connections per host")
	config.Add(path+maxIdleConn, 100, "http max idle connections")
	config.Add(path+maxConnsPerHost, 20, "http max connections per host")
	config.Add(path+idleConnTimeout, 90*time.Second, "http idle connections timeout")
	config.Add(path+disableKeepAlives, true, "http disable keep alives")
	config.Add(path+disableCompression, false, "http disable keep alives")
	config.Add(path+forceHTTP2, true, "http force http2")
	config.Add(path+tlsHandshakeTimeout, 10*time.Second, "TLS handshake timeout")
	config.Add(path+timeout, 30*time.Second, "timeout")
	config.Add(path+dialTimeout, 5*time.Second, "dial timeout")
	config.Add(path+keepAlive, 15*time.Second, "keep alive")
	config.Add(path+expectContinueTimeout, 1*time.Second, "expect continue timeout")
	config.Add(path+dualStack, true, "dual stack")
}
