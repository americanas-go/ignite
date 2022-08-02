package zap

import "github.com/americanas-go/config"

const (
	root             = "ignite.zap"
	consoleRoot      = root + ".console"
	consoleEnabled   = consoleRoot + ".enabled"
	consoleFormatter = consoleRoot + ".formatter"
	consoleLevel     = consoleRoot + ".level"
	fileRoot         = root + ".file"
	fileEnabled      = fileRoot + ".enabled"
	fileLevel        = fileRoot + ".level"
	filePath         = fileRoot + ".path"
	fileName         = fileRoot + ".name"
	fileMaxSize      = fileRoot + ".maxsize"
	fileCompress     = fileRoot + ".compress"
	fileMaxAge       = fileRoot + ".maxage"
	fileFormatter    = fileRoot + ".formatter"
)

func init() {

	config.Add(consoleEnabled, true, "enable/disable console logging")
	config.Add(consoleLevel, "INFO", "console log level")
	config.Add(consoleFormatter, "TEXT", "formatter TEXT/JSON")
	config.Add(fileEnabled, false, "enable/disable file logging")
	config.Add(fileLevel, "INFO", "console log level")
	config.Add(filePath, "/tmp", "log path")
	config.Add(fileName, "application.log", "log filename")
	config.Add(fileMaxSize, 100, "log file max size (MB)")
	config.Add(fileCompress, true, "log file compress")
	config.Add(fileMaxAge, 28, "log file max age (days)")
	config.Add(fileFormatter, "TEXT", "formatter TEXT/JSON")

}
