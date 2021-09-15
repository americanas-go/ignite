package client

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/valyala/fasthttp"
)

func NewClientWithConfigPath(ctx context.Context, path string) (*fasthttp.Client, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewClientWithOptions(ctx, options), nil
}

func NewClientWithOptions(ctx context.Context, o *Options) *fasthttp.Client {

	client := &fasthttp.Client{
		Name:                          o.Name,
		NoDefaultUserAgentHeader:      o.NoDefaultUserAgentHeader, // Don't send: User-Agent: fasthttp
		DialDualStack:                 o.DialDualStack,
		MaxConnsPerHost:               o.MaxConnsPerHost,
		MaxIdleConnDuration:           o.MaxIdleConnDuration,
		MaxConnDuration:               o.MaxConnDuration,
		MaxIdemponentCallAttempts:     o.MaxIdemponentCallAttempts,
		ReadBufferSize:                o.ReadBufferSize,  // Make sure to set this big enough that your whole request can be read at once.
		WriteBufferSize:               o.WriteBufferSize, // Same but for your response.
		ReadTimeout:                   o.ReadTimeout,
		WriteTimeout:                  o.WriteTimeout,
		MaxResponseBodySize:           o.MaxResponseBodySize,
		DisableHeaderNamesNormalizing: o.DisableHeaderNamesNormalizing, // If you set the case on your headers correctly you can enable this.
		MaxConnWaitTimeout:            o.MaxConnWaitTimeout,
	}

	return client
}

func NewClient(ctx context.Context) *fasthttp.Client {

	logger := log.FromContext(ctx)

	o, err := NewOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewClientWithOptions(ctx, o)
}
