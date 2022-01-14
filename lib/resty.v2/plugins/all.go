package plugins

import (
	"github.com/americanas-go/ignite"
	"github.com/americanas-go/ignite/lib/resty.v2/plugins/contrib"
	"github.com/americanas-go/ignite/lib/resty.v2/plugins/extra"
	"github.com/americanas-go/ignite/lib/resty.v2/resty"
)

// all plugins.
// they are activated by configuration.
//
// For example, enabling Log plugin:
//   ignite:
//     resty:
//        plugins:
//          log:
//            enabled: true
//            level: INFO
var All = []ignite.Plugin[*resty.Wrapper, *resty.Options]{
	contrib.Datadog, contrib.Health, contrib.Log, contrib.Newrelic, extra.RequestID, extra.Retry,
}
