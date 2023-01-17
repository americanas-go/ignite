package client

import (
	"time"

	"github.com/americanas-go/config"
)

const (
	root                          = "ignite.fasthttp.client"
	name                          = ".name"
	noDefaultUserAgentHeader      = ".noDefaultUserAgentHeader"
	maxConnsPerHost               = ".maxConnsPerHost"
	maxConnWaitTimeout            = ".maxConnWaitTimeout"
	readBufferSize                = ".readBufferSize"
	writeBufferSize               = ".writeBufferSize"
	readTimeout                   = ".readTimeout"
	writeTimeout                  = ".writeTimeout"
	maxIdleConnDuration           = ".maxIdleConnDuration"
	maxConnDuration               = ".maxConnDuration"
	disableHeaderNamesNormalizing = ".disableHeaderNamesNormalizing"
	dialDualStack                 = ".dialDualStack"
	maxResponseBodySize           = ".maxResponseBodySize"
	maxIdemponentCallAttempts     = ".maxIdemponentCallAttempts"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+name, "fasthttp-client", "used in User-Agent request header")
	config.Add(path+noDefaultUserAgentHeader, false, "User-Agent header to be excluded from the Request")
	config.Add(path+maxConnsPerHost, 512, "the maximum number of concurrent connections")
	config.Add(path+readBufferSize, 0, "per-connection buffer size for responses' reading")
	config.Add(path+writeBufferSize, 0, "per-connection buffer size for requests' writing")
	config.Add(path+maxConnWaitTimeout, 0*time.Second, "maximum amount of time to wait for a connection to be free")
	config.Add(path+readTimeout, 0*time.Second, "maximum duration for full response reading (including body)")
	config.Add(path+writeTimeout, 0*time.Second, "maximum duration for full request writing (including body)")
	config.Add(path+maxIdleConnDuration, 10*time.Second, "the default duration before idle keep-alive")
	config.Add(path+maxConnDuration, 0*time.Second, "Keep-alive connections are closed after this duration.")
	config.Add(path+disableHeaderNamesNormalizing, false, "header names are passed as-is without normalization")
	config.Add(path+dialDualStack, false, "attempt to connect to both ipv4 and ipv6 addresses if set to true")
	config.Add(path+maxResponseBodySize, 52428800, "maximum response body size")
	config.Add(path+maxIdemponentCallAttempts, 5, "maximum number of attempts for idempotent calls")
}
