package resty

import (
	"context"
	"net"
	"net/http"

	"github.com/americanas-go/config"
	"github.com/americanas-go/log"
	"github.com/go-resty/resty/v2"
)

type Plugin func(context.Context, *resty.Client) error

func NewClient(ctx context.Context, plugins ...Plugin) (*resty.Client, error) {
	opts, err := NewOptions()
	if err != nil {
		return nil, err
	}
	return NewClientWithOptions(ctx, opts, plugins...), nil
}

func NewClientWithOptions(ctx context.Context, options *Options, plugins ...Plugin) *resty.Client {

	logger := log.FromContext(ctx)

	logger.Tracef("creating resty client for host %s", options.Host)

	client := resty.New()

	dialer := &net.Dialer{
		Timeout:       options.ConnectionTimeout,
		FallbackDelay: options.FallbackDelay,
		KeepAlive:     options.KeepAlive,
	}

	transport := &http.Transport{
		DisableCompression:    config.Bool(transportDisableCompression),
		DisableKeepAlives:     config.Bool(transportDisableKeepAlives),
		MaxIdleConnsPerHost:   config.Int(transportMaxConnsPerHost),
		ResponseHeaderTimeout: config.Duration(transportResponseHeaderTimeout),
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     config.Bool(transportForceAttemptHTTP2),
		MaxIdleConns:          config.Int(transportMaxIdleConns),
		MaxConnsPerHost:       config.Int(transportMaxConnsPerHost),
		IdleConnTimeout:       config.Duration(transportIdleConnTimeout),
		TLSHandshakeTimeout:   config.Duration(transportTLSHandshakeTimeout),
		ExpectContinueTimeout: config.Duration(transportExpectContinueTimeout),
	}

	if options.Transport != nil {
		updateTransportWithOptions(transport, options.Transport)
	}

	client.
		SetTransport(transport).
		SetTimeout(options.RequestTimeout).
		SetDebug(options.Debug).
		SetHostURL(options.Host).
		SetCloseConnection(options.CloseConnection)

	for _, plugin := range plugins {
		if err := plugin(ctx, client); err != nil {
			panic(err)
		}
	}

	logger.Debugf("resty client created for host %s", options.Host)

	return client
}

func updateTransportWithOptions(transport *http.Transport, options *OptionsTransport) {
	transport.DisableCompression = options.DisableCompression
	transport.DisableKeepAlives = options.DisableKeepAlives
	transport.ForceAttemptHTTP2 = options.ForceAttemptHTTP2

	if options.MaxIdleConnsPerHost > 0 {
		transport.MaxIdleConnsPerHost = options.MaxIdleConnsPerHost
	}

	if options.ResponseHeaderTimeout > 0 {
		transport.ResponseHeaderTimeout = options.ResponseHeaderTimeout
	}

	if options.MaxIdleConns > 0 {
		transport.MaxIdleConns = options.MaxIdleConns
	}

	if options.MaxConnsPerHost > 0 {
		transport.MaxConnsPerHost = options.MaxConnsPerHost
	}

	if options.IdleConnTimeout > 0 {
		transport.IdleConnTimeout = options.IdleConnTimeout
	}

	if options.TLSHandshakeTimeout > 0 {
		transport.TLSHandshakeTimeout = options.TLSHandshakeTimeout
	}

	if options.ExpectContinueTimeout > 0 {
		transport.ExpectContinueTimeout = options.ExpectContinueTimeout
	}
}
