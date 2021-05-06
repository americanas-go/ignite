package logrus

import "github.com/americanas-go/config"

const (
	root           = "ignite.logrus"
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
	redisRoot      = root + ".redis"
	redisEnabled   = redisRoot + ".enabled"
	redisHost      = redisRoot + ".host"
	redisKey       = redisRoot + ".key"
	redisFormat    = redisRoot + ".format"
	redisApp       = redisRoot + ".app"
	redisPort      = redisRoot + ".port"
	redisDb        = redisRoot + ".db"
	formatter      = root + ".formatter"
	timeFormat     = root + ".time.format"
)

func init() {

	config.Add(consoleEnabled, true, "enable/disable console logging")
	config.Add(consoleLevel, "INFO", "console log level")
	config.Add(fileEnabled, false, "enable/disable file logging")
	config.Add(fileLevel, "INFO", "console log level")
	config.Add(filePath, "/tmp", "log path")
	config.Add(fileName, "application.l", "log filename")
	config.Add(fileMaxSize, 100, "log file max size (MB)")
	config.Add(fileCompress, true, "log file compress")
	config.Add(fileMaxAge, 28, "log file max age (days)")
	config.Add(redisEnabled, false, "enable/disable redis logging")
	config.Add(redisHost, "localhost", "redis host")
	config.Add(redisKey, "", "redis key")
	config.Add(redisFormat, "v1", "redis format")
	config.Add(redisApp, "logger", "redis app")
	config.Add(redisPort, 6379, "redis port")
	config.Add(redisDb, 0, "redis db")
	config.Add(formatter, "TEXT", "formatter TEXT/JSON/AWS_CLOUD_WATCH")
	config.Add(timeFormat, "2006/01/02 15:04:05.000", "defines time format")

}
