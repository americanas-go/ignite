package logrus

import (
	"github.com/americanas-go/config"
)

const (
	root           = "ignite.logrus"
	formatter      = root + ".formatter"
	consoleRoot    = root + ".console"
	consoleEnabled = consoleRoot + ".enabled"
	consoleLevel   = consoleRoot + ".level"
	fileRoot       = root + ".file"
	fileEnabled    = fileRoot + ".enabled"
	fileLevel      = fileRoot + ".level"
	filePath       = fileRoot + ".path"
	fileName       = fileRoot + ".name"
	fileMaxSize    = fileRoot + ".maxsize"
	fileCompress   = fileRoot + ".compress"
	fileMaxAge     = fileRoot + ".maxage"
	timeFormat     = root + ".time.format"
)

func init() {
	config.Add(formatter, TextFormatter, "log output formatter (TEXT/JSON/AWS_CLOUD_WATCH)")
	config.Add(consoleEnabled, true, "enable/disable console logging")
	config.Add(consoleLevel, "INFO", "console log level")
	config.Add(fileEnabled, false, "enable/disable file logging")
	config.Add(fileLevel, "INFO", "console log level")
	config.Add(filePath, "/tmp", "log path")
	config.Add(fileName, "application.log", "log filename")
	config.Add(fileMaxSize, 100, "log file max size (MB)")
	config.Add(fileCompress, true, "log file compress")
	config.Add(fileMaxAge, 28, "log file max age (days)")
	config.Add(timeFormat, "2006/01/02 15:04:05.000", "defines time format")
}

func Formatter() string {
	return config.String(formatter)
}
