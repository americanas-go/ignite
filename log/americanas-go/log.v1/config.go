package log

import "github.com/americanas-go/config"

const (
	root            = "ignite.log"
	tp              = root + ".type"
	logrusRoot      = root + ".logrus"
	logrusFormatter = logrusRoot + ".formatter"
)

func init() {
	config.Add(tp, "LOGRUS", "defines log implementation LOGRUS/ZAP/ZEROLOG")
	config.Add(logrusFormatter, "TEXT", "defines logrus formatter TEXT/JSON/CLOUDWATCH")
}

// Type returns the log implentation type from the configuration via the "ignite.log.type" key.
// Supported values are: `LOGRUS` (default), `ZAP`, `ZEROLOG`.
func Type() string {
	return config.String(tp)
}

// LogrusFormatter returns the logrus formatter configured via the "ignite.logrus.formatter" key.
// Supported values are: `TEXT` (default), `JSON`, `CLOUDWATCH`.
func LogrusFormatter() string {
	return config.String(logrusFormatter)
}
