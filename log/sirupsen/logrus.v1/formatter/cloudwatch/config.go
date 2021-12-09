package cloudwatch

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/log/sirupsen/logrus.v1"
)

const (
	root             = logrus.FormattersRoot + ".cloudwatch"
	prefixFields     = root + ".prefixFields"
	disableSorting   = root + ".disableSorting"
	quoteEmptyFields = root + ".quoteEmptyFields"
)

func init() {
	config.Add(prefixFields, []string{"RequestId"}, "defines fields will always be placed in front of the other fields")
	config.Add(disableSorting, false, "defines fields are sorted by default for a consistent output")
	config.Add(quoteEmptyFields, true, "will wrap empty fields in quotes if true")
}
