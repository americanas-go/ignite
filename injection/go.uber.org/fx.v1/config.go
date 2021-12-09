package fx

import "github.com/americanas-go/config"

const (
	root     = "ignite.fx"
	logLevel = root + ".log.level"
)

func init() {
	config.Add(logLevel, "DEBUG", "define log level")
}

func LogLevel() string {
	return config.String(logLevel)
}
