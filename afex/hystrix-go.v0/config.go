package hystrix

import (
	"strings"

	"github.com/americanas-go/config"
)

const (
	root                          = "ignite.hystrix"
	PluginsRoot                   = root + ".plugins"
	cmdRoot                       = root + ".commands"
	hystrixEnabled                = ".enabled"
	hystrixCommand                = ".name"
	hystrixTimeout                = ".timeout"
	hystrixRequestVolumeThreshold = ".requestVolumeThreshold"
	hystrixErrorPercentThreshold  = ".errorPercentThreshold"
	hystrixMaxConcurrentRequests  = ".maxConcurrentRequests"
	hystrixSleepWindow            = ".sleepWindow"
)

func CommandConfigsAdd(cmds []string) {
	for _, cmd := range cmds {
		CommandConfigAdd(cmd)
	}
}

func CommandConfigAdd(cmd string) {
	path := strings.Join([]string{cmdRoot, cmd}, ".")
	config.Add(path+hystrixEnabled, true, "enable/disable circuit breaker when necessary")
	config.Add(path+hystrixCommand, cmd, "defines hystrix command cmd")
	config.Add(path+hystrixTimeout, 10000, "defines how long to wait for command to complete, in milliseconds")
	config.Add(path+hystrixRequestVolumeThreshold, 10, "defines the minimum number of requests needed before a circuit can be tripped due to health")
	config.Add(path+hystrixErrorPercentThreshold, 5, "defines percentage of requests to open the circuit once the rolling measure of errors exceeds it")
	config.Add(path+hystrixMaxConcurrentRequests, 20, "defines how many commands of the same type can run at the same time")
	config.Add(path+hystrixSleepWindow, 5000, "defines how long, in milliseconds, to wait after a circuit opens before testing for recovery")
}

func IsCommandEnabled(cmd string) bool {
	path := strings.Join([]string{cmdRoot, cmd}, ".")
	return config.Bool(path + hystrixEnabled)
}
