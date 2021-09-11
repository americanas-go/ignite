package datadog

import (
	"github.com/americanas-go/config"
	awstrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/aws/aws-sdk-go-v2/aws"
)

type Options struct {
	Enabled      bool
	TraceOptions []awstrace.Option
}

func NewOptions(traceOptions ...awstrace.Option) (*Options, error) {
	o := &Options{TraceOptions: traceOptions}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func NewOptionsWithPath(path string, traceOptions ...awstrace.Option) (opts *Options, err error) {

	opts, err = NewOptions(traceOptions...)
	if err != nil {
		return nil, err
	}

	err = config.UnmarshalWithPath(path, opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}
