package resty

import (
	"time"
)

type Options struct {
	ConnectionTimeout time.Duration `default:"3m" desc:"defines http connection timeout"`
	CloseConnection   bool          `default:"false" desc:"defines http close connection"`
	Debug             bool          `default:"false" desc:"defines debug request"`
	FallbackDelay     time.Duration `default:"300ms" desc:"defines fallbackDelay"`
	Host              string        `default:"http://localhost" desc:"defines host request"`
	KeepAlive         time.Duration `default:"30s" desc:"defines http keepalive"`
	RequestTimeout    time.Duration `default:"30s" desc:"defines http request timeout"`
	Transport         Transport
	Plugins           Plugins
}

func (o *Options) Root() string {
	return "ignite.resty"
}

func (o *Options) PostLoad() error {
	return nil
}

type Transport struct {
	DisableCompression    bool          `default:"false" desc:"enabled/disable transport compression"`
	DisableKeepAlives     bool          `default:"false" desc:"enabled/disable transport keep alives"`
	ExpectContinueTimeout time.Duration `default:"1s" desc:"define transport expect continue timeout"`
	ForceAttemptHTTP2     bool          `config:"forceAttemptHTTP2" default:"true" desc:"define transport force attempt http2"`
	IdleConnTimeout       time.Duration `default:"90s" desc:"define transport idle conn timeout"`
	MaxIdleConnsPerHost   int           `default:"2" desc:"define transport max idle conns per host"`
	MaxIdleConns          int           `default:"100" desc:"define transport max idle conns"`
	MaxConnsPerHost       int           `default:"100" desc:"define transport max conns per host"`
	ResponseHeaderTimeout time.Duration `default:"2s" desc:"define transport response header timeout"`
	TLSHandshakeTimeout   time.Duration `default:"10s" desc:"define transport TLS handshake timeout"`
}
type Plugins struct {
	Datadog     Datadog
	Health      Health
	Log         Log
	Newrelic    Newrelic
	Opentracing Opentracing
	RequestID   RequestID `config:"requestId"`
	Retry       Retry
}

type Datadog struct {
	Enabled       bool              `default:"false" desc:"enable/disable datadog integration"`
	OperationName string            `desc:"Datadog operation name"`
	Tags          map[string]string `desc:"Datadog span tags"`
}

type Health struct {
	Enabled     bool   `default:"false" desc:"enable/disable health"`
	Name        string `default:"rest api" desc:"health name"`
	Host        string `default:"" desc:"health host"`
	Endpoint    string `default:"/resource-status" desc:"health endpoint"`
	Description string `default:"default connection" desc:"health check"`
	Required    bool   `default:"true" desc:"if its required"`
}
type Log struct {
	Enabled bool   `default:"false" desc:"enable/disable log"`
	Level   string `default:"DEBUG" desc:"sets log level INFO/DEBUG/TRACE"`
}
type Newrelic struct {
	Enabled bool `default:"false" desc:"enable/disable newrelic integration"`
}
type Opentracing struct {
	Enabled bool `default:"false" desc:"enable/disable opentracing"`
}
type RequestID struct {
	Enabled bool `default:"true" desc:"enable/disable requestId"`
}
type Retry struct {
	Enabled     bool          `default:"true" desc:"enable/disable retry"`
	Count       int           `default:"0" desc:"defines global max http retries"`
	MaxWaitTime time.Duration `default:"2s" desc:"defines global max retry wait time"`
	WaitTime    time.Duration `default:"200ms" desc:"defines global retry wait time"`
}
