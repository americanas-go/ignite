package buntdb

import "github.com/americanas-go/config"

const (
	root                 = "ignite.buntdb"
	path                 = ".path"
	syncPolicy           = ".syncPolicy"
	autoShrinkPercentage = ".autoShrink.percentage"
	autoShrinkMinSize    = ".autoShrink.minSize"
	autoShrinkDisabled   = ".autoShrink.disabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(p string) {
	config.Add(p+path, ":memory:", "open opens a database at the provided path")
	config.Add(p+syncPolicy, 1, "adjusts how often the data is synced to disk (Never: 0, EverySecond: 1, Always: 2)")
	config.Add(p+autoShrinkPercentage, 100, "is used by the background process to trigger a shrink of the aof file when the size of the file is larger than the percentage of the result of the previous shrunk file")
	config.Add(p+autoShrinkMinSize, 32*1024*102, "defines the minimum size of the aof file before an automatic shrink can occur")
	config.Add(p+autoShrinkDisabled, false, "turns off automatic background shrinking")
}
