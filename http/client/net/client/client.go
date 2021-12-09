package client

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/americanas-go/log"
)

// NewClientWithConfigPath returns a http client with options from config path.
func NewClientWithConfigPath(ctx context.Context, path string) (*http.Client, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewClientWithOptions(ctx, options), nil
}

// NewClientWithOptions returns a new http client with options.
func NewClientWithOptions(ctx context.Context, options *Options) *http.Client {

	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:  options.DialTimeout,
			Deadline: time.Time{},
			// LocalAddr:     nil,
			// FallbackDelay: 0,
			KeepAlive: options.KeepAlive,
			// Resolver:      nil,
			// Control:       nil,
		}).DialContext,
		/*
			TLSClientConfig: &tls.Config{
				Rand:                        nil,
				Time:                        nil,
				Certificates:                []tls.Certificate{
					{
						Certificate:                  nil,
						PrivateKey:                   nil,
						SupportedSignatureAlgorithms: nil,
						OCSPStaple:                   nil,
						SignedCertificateTimestamps:  nil,
						Leaf:                         nil,
					},
				},
				NameToCertificate:           nil,
				GetCertificate:              nil,
				GetClientCertificate:        nil,
				GetConfigForClient:          nil,
				VerifyPeerCertificate:       nil,
				VerifyConnection:            nil,
				RootCAs:                     nil,
				NextProtos:                  nil,
				ServerName:                  "",
				ClientAuth:                  0,
				ClientCAs:                   nil,
				InsecureSkipVerify:          false,
				CipherSuites:                nil,
				PreferServerCipherSuites:    false,
				SessionTicketsDisabled:      false,
				SessionTicketKey:            [32]byte{},
				ClientSessionCache:          nil,
				MinVersion:                  0,
				MaxVersion:                  0,
				CurvePreferences:            nil,
				DynamicRecordSizingDisabled: false,
				Renegotiation:               0,
				KeyLogWriter:                nil,
			},
		*/
		TLSHandshakeTimeout: options.TLSHandshakeTimeout,
		DisableKeepAlives:   options.DisableKeepAlives,
		DisableCompression:  options.DisableCompression,
		MaxIdleConns:        options.MaxIdleConn,
		MaxIdleConnsPerHost: options.MaxIdleConnPerHost,
		MaxConnsPerHost:     options.MaxConnsPerHost,
		IdleConnTimeout:     options.IdleConnTimeout,
		// ResponseHeaderTimeout:  0,
		ExpectContinueTimeout: options.ExpectContinueTimeout,
		// MaxResponseHeaderBytes: 0,
		// WriteBufferSize:        0,
		// ReadBufferSize:         0,
		ForceAttemptHTTP2: options.ForceHTTP2,
	}

	return &http.Client{
		Transport: tr,
		// CheckRedirect: nil,
		// Jar:           nil,
		Timeout: options.Timeout,
	}
}

// NewClient returns a new http client with default options.
func NewClient(ctx context.Context) *http.Client {

	logger := log.FromContext(ctx)

	o, err := NewOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewClientWithOptions(ctx, o)
}
