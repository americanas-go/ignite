package text

import (
	"github.com/americanas-go/config"
)

const (
	root                      = "ignite.logrus.formatters.text"
	forceColors               = root + ".forceColors"
	disableColors             = root + ".disableColors"
	disableQuote              = root + ".disableQuote"
	forceQuote                = root + ".forceQuote"
	environmentOverrideColors = root + ".environmentOverrideColors"
	disableTimestamp          = root + ".disableTimestamp"
	fullTimestamp             = root + ".fullTimestamp"
	timestampFormat           = root + ".timestampFormat"
	disableSorting            = root + ".disableSorting"
	disableLevelTruncation    = root + ".disableLevelTruncation"
	padLevelText              = root + ".padLevelText"
	quoteEmptyFields          = root + ".quoteEmptyFields"
)

func init() {
	config.Add(forceColors, false, "set to true to bypass checking for a TTY before outputting colors")
	config.Add(disableColors, false, "force disabling colors")
	config.Add(forceQuote, false, "force quoting of all values")
	config.Add(disableQuote, false, "disables quoting for all values. will have a lower priority than ForceQuote")
	config.Add(environmentOverrideColors, false, "override coloring based on CLICOLOR and CLICOLOR_FORCE. - https://bixense.com/clicolors/")
	config.Add(disableTimestamp, false, "disable timestamp logging")
	config.Add(fullTimestamp, true, "enable logging the full timestamp when a TTY is attached instead of just the time passed since beginning of execution")
	config.Add(timestampFormat, "2006/01/02 15:04:05.000", "to use for display when a full timestamp is printed")
	config.Add(disableSorting, false, "the fields are sorted by default for a consistent output")
	config.Add(disableLevelTruncation, true, "disables the truncation of the level text to 4 characters")
	config.Add(padLevelText, false, "adds padding the level text so that all the levels output at the same length")
	config.Add(quoteEmptyFields, false, "will wrap empty fields in quotes if true")
}
