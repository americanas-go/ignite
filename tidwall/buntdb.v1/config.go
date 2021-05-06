package buntdb

import "github.com/americanas-go/config"

const (
	root                 = "ignite.buntdb"
	path                 = root + ".path"
	syncPolicy           = root + ".syncPolicy"
	autoShrinkPercentage = root + ".autoShrink.percentage"
	autoShrinkMinSize    = root + ".autoShrink.minSize"
	autoShrinkDisabled   = root + ".autoShrink.disabled"
)

func init() {
	config.Add(path, ":memory:", "open opens a database at the provided path")
	config.Add(syncPolicy, 1, "adjusts how often the data is synced to disk (Never: 0, EverySecond: 1, Always: 2)")
	config.Add(autoShrinkPercentage, 100, "is used by the background process to trigger a shrink of the aof file when the size of the file is larger than the percentage of the result of the previous shrunk file")
	config.Add(autoShrinkMinSize, 32*1024*102, "defines the minimum size of the aof file before an automatic shrink can occur")
	config.Add(autoShrinkDisabled, false, "turns off automatic background shrinking")
}
