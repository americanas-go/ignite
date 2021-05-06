package resty

import (
	"time"
)

type OptionsTransport struct {
	DisableCompression    bool
	DisableKeepAlives     bool
	MaxIdleConnsPerHost   int
	ResponseHeaderTimeout time.Duration
	ForceAttemptHTTP2     bool `config:"forceAttemptHTTP2"`
	MaxIdleConns          int
	MaxConnsPerHost       int
	IdleConnTimeout       time.Duration
	TLSHandshakeTimeout   time.Duration
	ExpectContinueTimeout time.Duration
}
