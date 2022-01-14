package resty

import (
	"fmt"
	"testing"
	"time"

	"github.com/americanas-go/ignite"
)

func TestOptions(t *testing.T) {
	tests := []struct {
		name  string
		setup func()
		want  func() *Options
	}{
		{
			name:  "loads options",
			setup: func() {},
			want: func() *Options {
				o := ignite.New[*Options]()
				o.Debug = false
				o.CloseConnection = false
				o.Host = "http://localhost"
				o.ConnectionTimeout = 3 * time.Minute
				o.KeepAlive = 30 * time.Second
				o.RequestTimeout = 30 * time.Second
				o.FallbackDelay = 300 * time.Millisecond
				t := &(o.Transport)
				t.DisableCompression = false
				t.DisableKeepAlives = false
				t.ExpectContinueTimeout = 1 * time.Second
				t.ForceAttemptHTTP2 = true
				t.IdleConnTimeout = 90 * time.Second
				t.MaxConnsPerHost = 100
				t.MaxIdleConns = 100
				t.MaxIdleConnsPerHost = 2
				t.ResponseHeaderTimeout = 2 * time.Second
				t.TLSHandshakeTimeout = 10 * time.Second
				// Plugins ---
				p := &(o.Plugins)
				// Datadog
				p.Datadog.Enabled = false
				p.Datadog.OperationName = ""
				// Health
				p.Health.Enabled = false
				p.Health.Description = "default connection"
				p.Health.Endpoint = "/resource-status"
				p.Health.Host = ""
				p.Health.Name = "rest api"
				p.Health.Required = true
				// Log
				p.Log.Enabled = false
				p.Log.Level = "DEBUG"
				// Newrelic
				p.Newrelic.Enabled = false
				// Opentracing
				p.Opentracing.Enabled = false
				// Retry
				p.Retry.Enabled = true
				p.Retry.Count = 0
				p.Retry.MaxWaitTime = 2 * time.Second
				p.Retry.WaitTime = 200 * time.Millisecond
				//RequestID
				p.RequestID.Enabled = true
				return o
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got, e := ignite.Load[*Options]()
			if e != nil {
				t.Errorf("Unexpected error %v", e)
			}
			want := tt.want()
			if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", want) {
				t.Errorf("\nwant\t%v\ngot \t%v", want, got)
			}

		})
	}
}
