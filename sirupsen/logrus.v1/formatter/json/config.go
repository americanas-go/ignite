package json

import (
	"github.com/americanas-go/config"
)

const (
	root              = "ignite.logrus.formatters.json"
	timestampFormat   = root + ".timestampFormat"
	disableTimestamp  = root + ".disableTimestamp"
	disableHTMLEscape = root + ".disableHTMLEscape"
	prettyPrint       = root + ".prettyPrint"
)

func init() {
	config.Add(timestampFormat, "2006/01/02 15:04:05.000", "sets the format used for marshaling timestamps")
	config.Add(disableTimestamp, false, "allows disabling automatic timestamps in output")
	config.Add(disableHTMLEscape, false, "allows disabling html escaping in output")
	config.Add(prettyPrint, false, "will indent all json logs")
}
