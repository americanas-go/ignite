package fiber

import (
	"github.com/americanas-go/config"
	"github.com/gofiber/fiber/v2"
)

const (
	root                      = "ignite.fiber"
	port                      = ".port"
	tp                        = ".type"
	configRoot                = ".config"
	prefork                   = configRoot + ".prefork"
	serverHeader              = configRoot + ".serverHeader"
	strictRouting             = configRoot + ".strictRouting"
	caseSensitive             = configRoot + ".caseSensitive"
	immutable                 = configRoot + ".immutable"
	unescapePath              = configRoot + ".unescapePath"
	etag                      = configRoot + ".etag"
	bodyLimit                 = configRoot + ".bodyLimit"
	concurrency               = configRoot + ".concurrency"
	readTimeout               = configRoot + ".readTimeout"
	writeTimeout              = configRoot + ".writeTimeout"
	idleTimeout               = configRoot + ".idleTimeout"
	readBufferSize            = configRoot + ".readBufferSize"
	writeBufferSize           = configRoot + ".writeBufferSize"
	compressedFileSuffix      = configRoot + ".compressedFileSuffix"
	proxyHeader               = configRoot + ".proxyHeader"
	getOnly                   = configRoot + ".GETOnly"
	reduceMemoryUsage         = configRoot + ".reduceMemoryUsage"
	network                   = configRoot + ".network"
	disableKeepalive          = configRoot + ".disableKeepalive"
	disableDefaultDate        = configRoot + ".disableDefaultDate"
	disableDefaultContentType = configRoot + ".disableDefaultContentType"
	disableHeaderNormalizing  = configRoot + ".disableHeaderNormalizing"
	disableStartupMessage     = configRoot + ".disableStartupMessage"
	PluginsRoot               = root + ".plugins"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+port, 8082, "server http port")
	config.Add(path+tp, "REST", "defines type for applicaton ")
	config.Add(path+prefork, false, "enables use of the SO_REUSEPORT socket option. This will spawn multiple Go processes listening on the same port. learn more about socket sharding.")
	config.Add(path+serverHeader, "", "enables the Server HTTP header with the given value.")
	config.Add(path+strictRouting, false, "when enabled, the router treats /foo and /foo/ as different. Otherwise, the router treats /foo and /foo/ as the same.")
	config.Add(path+caseSensitive, false, "when enabled, /Foo and /foo are different routes. When disabled, /Fooand /foo are treated the same.")
	config.Add(path+immutable, false, "when enabled, all values returned by context methods are immutable. By default, they are valid until you return from the handler; see issue #185.")
	config.Add(path+unescapePath, false, "converts all encoded characters in the route back before setting the path for the context, so that the routing can also work with URL encoded special characters")
	config.Add(path+etag, false, "enable or disable etag header generation, since both weak and strong etags are generated using the same hashing method (CRC-32). Weak ETags are the default when enabled.")
	config.Add(path+bodyLimit, 4*1024*1024, "sets the maximum allowed size for a request body, if the size exceeds the configured limit, it sends 413 - Request Entity Too Large response.")
	config.Add(path+concurrency, 256*1024, "maximum number of concurrent connections.")
	config.Add(path+readTimeout, "0s", "the amount of time allowed to read the full request, including the body. The default timeout is unlimited.")
	config.Add(path+writeTimeout, "0s", "the maximum duration before timing out writes of the response. The default timeout is unlimited.")
	config.Add(path+idleTimeout, "0s", "the maximum amount of time to wait for the next request when keep-alive is enabled. If IdleTimeout is zero, the value of ReadTimeout is used.")
	config.Add(path+readBufferSize, 4096, "per-connection buffer size for requests' reading. This also limits the maximum header size. Increase this buffer if your clients send multi-KB RequestURIs and/or multi-KB headers (for example, BIG cookies).")
	config.Add(path+writeBufferSize, 4096, "per-connection buffer size for responses' writing.")
	config.Add(path+compressedFileSuffix, ".fiber.gz", "adds a suffix to the original file name and tries saving the resulting compressed file under the new file name.")
	config.Add(path+proxyHeader, "", "this will enable c.IP() to return the value of the given header key. By default c.IP()will return the Remote IP from the TCP connection, this property can be useful if you are behind a load balancer e.g. X-Forwarded-*.")
	config.Add(path+getOnly, false, "rejects all non-GET requests if set to true. This option is useful as anti-DoS protection for servers accepting only GET requests. The request size is limited by ReadBufferSize if GETOnly is set.")
	config.Add(path+reduceMemoryUsage, false, "aggressively reduces memory usage at the cost of higher CPU usage if set to true")
	config.Add(path+network, fiber.NetworkTCP4, "known networks are \"tcp\", \"tcp4\" (IPv4-only), \"tcp6\" (IPv6-only)")
	config.Add(path+disableKeepalive, false, "disable keep-alive connections, the Server will close incoming connections after sending the first response to the client")
	config.Add(path+disableDefaultDate, false, "when set to true causes the default date header to be excluded from the response.")
	config.Add(path+disableDefaultContentType, false, "when set to true, causes the default Content-Type header to be excluded from the Response.")
	config.Add(path+disableHeaderNormalizing, false, "by default all header names are normalized: conteNT-tYPE -> Content-Type")
	config.Add(path+disableStartupMessage, false, "when set to true, it will not print out debug information")
}
