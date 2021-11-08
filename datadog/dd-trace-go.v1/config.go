package datadog

import (
	"math"
	"time"

	"github.com/americanas-go/config"
)

const (
	root                  = "ignite.datadog"
	service               = root + ".service"
	env                   = root + ".env"
	version               = root + ".version"
	tracerEnabled         = root + ".tracer.enabled"
	profilerEnabled       = root + ".profiler.enabled"
	tags                  = root + ".tags"
	host                  = root + ".host"
	port                  = root + ".port"
	lambdaMode            = root + ".lambdaMode"
	analytics             = root + ".analytics"
	analyticsRate         = root + ".analyticsRate"
	debugMode             = root + ".debugMode"
	debugStack            = root + ".debugStack"
	httpClientRoot        = root + ".httpClient"
	maxIdleConnPerHost    = httpClientRoot + ".maxIdleConnPerHost"
	maxIdleConn           = httpClientRoot + ".maxIdleConn"
	maxConnsPerHost       = httpClientRoot + ".maxConnsPerHost"
	idleConnTimeout       = httpClientRoot + ".idleConnTimeout"
	disableKeepAlives     = httpClientRoot + ".disableKeepAlives"
	forceHTTP2            = httpClientRoot + ".forceHTTP2"
	tlsHandshakeTimeout   = httpClientRoot + ".TLSHandshakeTimeout"
	timeout               = httpClientRoot + ".timeout"
	keepAlive             = httpClientRoot + ".keepAlive"
	expectContinueTimeout = httpClientRoot + ".expectContinueTimeout"
	dualStack             = httpClientRoot + ".dualStack"
	logRoot               = root + ".log"
	logLevel              = logRoot + ".level"
)

func init() {
	config.Add(service, "", "service name for datadog")
	config.Add(env, "", "service env")
	config.Add(tracerEnabled, true, "enables tracer datadog")
	config.Add(profilerEnabled, true, "enables profiler datadog")
	config.Add(logLevel, "DEBUG", "log level")
	config.Add(tags, map[string]string{}, "sets a key/value pair which will be set as a tag on all spans created by tracer. This option may be used multiple times")
	config.Add(host, "localhost", "sets the address where the agent is located")
	config.Add(port, "8126", "sets the port where the agent is located")
	config.Add(lambdaMode, false, "enables lambda mode on the tracer, for use with AWS Lambda")
	config.Add(analytics, false, "allows specifying whether Trace Search & Analytics should be enabled for integrations")
	config.Add(analyticsRate, math.NaN(), "sets the global sampling rate for sampling APM events")
	config.Add(debugMode, false, "enables debug mode on the tracer, resulting in more verbose logging.")
	config.Add(debugStack, true, "can be used to globally enable or disable the collection of stack traces when spans finish with errors.")
	config.Add(maxIdleConnPerHost, 1, "http max idle connections per host")
	config.Add(maxIdleConn, 100, "http max idle connections")
	config.Add(maxConnsPerHost, 20, "http max connections per host")
	config.Add(idleConnTimeout, 90*time.Second, "http idle connections timeout")
	config.Add(disableKeepAlives, true, "http disable keep alives")
	config.Add(forceHTTP2, true, "http force http2")
	config.Add(tlsHandshakeTimeout, 10*time.Second, "TLS handshake timeout")
	config.Add(timeout, 30*time.Second, "timeout")
	config.Add(keepAlive, 15*time.Second, "keep alive")
	config.Add(expectContinueTimeout, 1*time.Second, "expect continue timeout")
	config.Add(dualStack, true, "dual stack")
	config.Add(version, "0.0.0", "application version")
}

// IsTracerEnabled returns config value from key ignite.datadog.trace.enabled where default is true.
func IsTracerEnabled() bool {
	return config.Bool(tracerEnabled)
}

// IsProfilerEnabled returns config value from key ignite.datadog.profiler.enabled where default is true.
func IsProfilerEnabled() bool {
	return config.Bool(profilerEnabled)
}

// Service returns config value from key ignite.datadog.service where default is empty.
func Service() string {
	return config.String(service)
}
