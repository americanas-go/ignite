package retry

import (
	"time"

	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/go-resty/resty.v2"
)

const (
	root        = resty.PluginsRoot + ".retry"
	enabled     = root + ".enabled"
	count       = root + ".count"
	waitTime    = root + ".waitTime"
	maxWaitTime = root + ".maxWaitTime"
)

func init() {
	config.Add(enabled, true, "enable/disable retry")
	config.Add(count, 0, "defines global max http retries")
	config.Add(waitTime, 200*time.Millisecond, "defines global retry wait time")
	config.Add(maxWaitTime, 2*time.Second, "defines global max retry wait time")
}
