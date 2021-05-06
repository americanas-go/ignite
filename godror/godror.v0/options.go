package godror

import (
	"time"

	"github.com/americanas-go/config"
)

type Options struct {
	DataSourceName  string        `config:"datasourcename"`
	ConnMaxLifetime time.Duration `config:"connmaxlifetime"`
	MaxIdleConns    int           `config:"maxidleconns"`
	MaxOpenConns    int           `config:"maxopenconns"`
}

func NewOptions() (*Options, error) {
	o := &Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func NewOptionsWithPath(path string) (opts *Options, err error) {

	opts, err = NewOptions()
	if err != nil {
		return nil, err
	}

	err = config.UnmarshalWithPath(path, opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}
