package resty

import (
	"context"
	"net"
	"net/http"

	"github.com/americanas-go/log"
	"github.com/go-resty/resty/v2"
)

type Wrapper struct {
	Instance *resty.Client
	Options  *Options
}

func (w *Wrapper) Init(ctx context.Context, o *Options) error {
	w.Options = o
	logger := log.FromContext(ctx)

	logger.Tracef("creating resty client for host %s", o.Host)

	w.Instance = resty.New()

	dialer := &net.Dialer{
		Timeout:       o.ConnectionTimeout,
		FallbackDelay: o.FallbackDelay,
		KeepAlive:     o.KeepAlive,
	}

	transport := &http.Transport{
		DisableCompression:    o.Transport.DisableCompression,
		DisableKeepAlives:     o.Transport.DisableKeepAlives,
		MaxIdleConnsPerHost:   o.Transport.MaxIdleConnsPerHost,
		ResponseHeaderTimeout: o.Transport.ResponseHeaderTimeout,
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     o.Transport.ForceAttemptHTTP2,
		MaxIdleConns:          o.Transport.MaxIdleConns,
		MaxConnsPerHost:       o.Transport.MaxConnsPerHost,
		IdleConnTimeout:       o.Transport.IdleConnTimeout,
		TLSHandshakeTimeout:   o.Transport.TLSHandshakeTimeout,
		ExpectContinueTimeout: o.Transport.ExpectContinueTimeout,
	}

	w.Instance.
		SetTransport(transport).
		SetTimeout(o.RequestTimeout).
		SetDebug(o.Debug).
		SetHostURL(o.Host).
		SetCloseConnection(o.CloseConnection)

	logger.Debugf("resty client created for host %s", o.Host)
	return nil
}
