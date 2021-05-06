package elasticsearch

import (
	"time"

	"github.com/americanas-go/config"
)

const (
	root                  = "ignite.elasticsearch"
	addresses             = root + ".addresses"
	username              = root + ".username"
	password              = root + ".password"
	cloudID               = root + ".cloudID"
	apiKey                = root + ".APIKey"
	caCert                = root + ".CACert"
	retryOnStatus         = root + ".retryOnStatus"
	disableRetry          = root + ".disableRetry"
	enableRetryOnTimeout  = root + ".enableRetryOnTimeout"
	maxRetries            = root + ".maxRetries"
	discoverNodesOnStart  = root + ".discoverNodesOnStart"
	discoverNodesInterval = root + ".discoverNodesInterval"
	enableMetrics         = root + ".enableMetrics"
	enableDebugLogger     = root + ".enableDebugLogger"
	retryBackoff          = root + ".retryBackoff"
	PluginsRoot           = root + ".plugins"
)

func init() {
	config.Add(addresses, []string{"http://127.0.0.1:9200"}, "a list of Elasticsearch nodes to use")
	config.Add(username, "", "username for HTTP Basic Authentication")
	config.Add(password, "", "password for HTTP Basic Authentication", config.WithHide())
	config.Add(cloudID, "", "endpoint for the Elastic Service (https://elastic.co/cloud)")
	config.Add(apiKey, "", "base64-encoded token for authorization; if set, overrides username and password")
	config.Add(caCert, "", "PEM-encoded certificate authorities")
	config.Add(retryOnStatus, []int{502, 503, 504}, "List of status codes for retry")
	config.Add(disableRetry, false, "")
	config.Add(enableRetryOnTimeout, false, "")
	config.Add(maxRetries, 3, "")
	config.Add(discoverNodesOnStart, false, "discover nodes when initializing the client")
	config.Add(discoverNodesInterval, 0*time.Millisecond, "discover nodes periodically. Default: 0 (disabled)")
	config.Add(enableMetrics, false, "enable the metrics collection")
	config.Add(enableDebugLogger, false, "enable the debug logging")
	config.Add(retryBackoff, 5*time.Millisecond, "optional backoff duration")
}
