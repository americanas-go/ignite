package log

import "github.com/americanas-go/config"

const (
	root = "ignite.log"
	tp   = root + ".type"
)

func init() {
	config.Add(tp, "LOGRUS", "defines log implementation LOGRUS/ZAP/ZEROLOG")
}

// Type returns the log implentation type from the configuration via the "ignite.log.type" key.
// Supported values are: `LOGRUS` (default), `ZAP`, `ZEROLOG`.
func Type() string {
	return config.String(tp)
}
