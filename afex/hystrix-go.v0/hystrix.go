package hystrix

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/americanas-go/config"
	"github.com/americanas-go/log"
)

func ConfigureCommands(cmds []string) error {
	for _, cmd := range cmds {
		if err := ConfigureCommand(cmd); err != nil {
			return err
		}
	}
	return nil
}

func ConfigureCommand(cmd string) error {

	CommandConfigAdd(cmd)
	config.Load()

	options, err := NewOptionsFromCommand(cmd)
	if err != nil {
		return err
	}
	if err := ConfigureCommandWithOptions(cmd, options); err != nil {
		return err
	}
	return nil
}

func ConfigureCommandWithOptions(cmd string, options *Options) error {
	hystrix.ConfigureCommand(cmd, hystrix.CommandConfig{
		Timeout:                options.Timeout,
		MaxConcurrentRequests:  options.MaxConcurrentRequests,
		RequestVolumeThreshold: options.RequestVolumeThreshold,
		SleepWindow:            options.SleepWindow,
		ErrorPercentThreshold:  options.ErrorPercentThreshold,
	})
	hystrix.SetLogger(log.GetLogger())
	return nil
}
