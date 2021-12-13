package elasticsearch

import (
	"time"

	"github.com/americanas-go/config"
)

const (
	root                  = "ignite.elasticsearch"
	addresses             = ".addresses"
	username              = ".username"
	password              = ".password"
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

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+addresses, []string{"http://127.0.0.1:9200"}, "a list of Elasticsearch nodes to use")
	config.Add(path+username, "", "username for HTTP Basic Authentication")
	config.Add(path+password, "", "password for HTTP Basic Authentication", config.WithHide())
	config.Add(path+cloudID, "", "endpoint for the Elastic Service (https://elastic.co/cloud)")
	config.Add(path+apiKey, "", "base64-encoded token for authorization; if set, overrides username and password")
	config.Add(path+caCert, "", "PEM-encoded certificate authorities")
	config.Add(path+retryOnStatus, []int{502, 503, 504}, "List of status codes for retry")
	config.Add(path+disableRetry, false, "")
	config.Add(path+enableRetryOnTimeout, false, "")
	config.Add(path+maxRetries, 3, "")
	config.Add(path+discoverNodesOnStart, false, "discover nodes when initializing the client")
	config.Add(path+discoverNodesInterval, 0*time.Millisecond, "discover nodes periodically. Default: 0 (disabled)")
	config.Add(path+enableMetrics, false, "enable the metrics collection")
	config.Add(path+enableDebugLogger, false, "enable the debug logging")
	config.Add(path+retryBackoff, 5*time.Millisecond, "optional backoff duration")
}
