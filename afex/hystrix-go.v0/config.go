package hystrix

import (
	"github.com/americanas-go/config"
)

const (
	root                  = "ignite.elasticsearch"
	addresses             = ".addresses"
	pu                    = ".username"
	pp                    = ".password"
	cloudID               = ".cloudID"
	apiKey                = ".APIKey"
	caCert                = ".CACert"
	retryOnStatus         = ".retryOnStatus"
	disableRetry          = ".disableRetry"
	enableRetryOnTimeout  = ".enableRetryOnTimeout"
	maxRetries            = ".maxRetries"
	discoverNodesOnStart  = ".discoverNodesOnStart"
	discoverNodesInterval = ".discoverNodesInterval"
	enableMetrics         = ".enableMetrics"
	enableDebugLogger     = ".enableDebugLogger"
	retryBackoff          = ".retryBackoff"
	PluginsRoot           = root + ".plugins"
)

const (
	root           = "pkg.lib.hystrix"
	prometheusRoot = root + ".prometheus"
	namespace      = prometheusRoot + ".namespace"
	labels         = prometheusRoot + ".labels"

	hystrixEnabled                = ".enabled"
	hystrixCommand                = ".name"
	hystrixTimeout                = ".timeout"
	hystrixRequestVolumeThreshold = ".requestVolumeThreshold"
	hystrixErrorPercentThreshold  = ".errorPercentThreshold"
	hystrixMaxConcurrentRequests  = ".maxConcurrentRequests"
	hystrixSleepWindow            = ".sleepWindow"
)

func init() {
	config.Add(namespace, "hystrix", "defines hystrix namespace")
	config.Add(labels, map[string]interface{}{}, "defines hystrix labels")
}

func CmdConfigAdd(path string, name string) {
	path += ".hystrix"
	config.Add(path+hystrixEnabled, true, "enable/disable circuit breaker when necessary")
	config.Add(path+hystrixCommand, name, "defines hystrix command name")
	config.Add(path+hystrixTimeout, 10000, "defines how long to wait for command to complete, in milliseconds")
	config.Add(path+hystrixRequestVolumeThreshold, 10, "defines the minimum number of requests needed before a circuit can be tripped due to health")
	config.Add(path+hystrixErrorPercentThreshold, 5, "defines percentage of requests to open the circuit once the rolling measure of errors exceeds it")
	config.Add(path+hystrixMaxConcurrentRequests, 20, "defines how many commands of the same type can run at the same time")
	config.Add(path+hystrixSleepWindow, 5000, "defines how long, in milliseconds, to wait after a circuit opens before testing for recovery")
}
