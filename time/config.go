package time

import (
	"time"

	"github.com/americanas-go/config"
)

const (
	root     = "ignite.time"
	format   = root + ".timestamp"
	location = root + ".location"
)

var (
	fmt string
	loc *time.Location
)

func init() {
	config.Add(format, "2006/01/02 15:04:05.000", "time format")
	config.Add(location, time.UTC.String(), "time location")
}

// Format returns config value from key ignite.time.timestamp where default is 2006/01/02 15:04:05.000.
func Format() string {
	if fmt == "" {
		fmt = config.String(format)
	}
	return fmt
}

// Location returns config value from key ignite.time.location where default is UTC.
func Location() *time.Location {
	if loc == nil {
		var err error
		locStr := config.String(location)
		loc, err = time.LoadLocation(locStr)
		if err != nil {
			panic(err)
		}
	}
	return loc
}
