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

func Type() string {
	return config.String(tp)
}

func LogrusFormatter() string {
	return config.String(logrusFormatter)
}
