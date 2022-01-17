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
				return &Options{
					Debug:             false,
					CloseConnection:   false,
					Host:              "http://localhost",
					ConnectionTimeout: 3 * time.Minute,
					KeepAlive:         30 * time.Second,
					RequestTimeout:    30 * time.Second,
					FallbackDelay:     300 * time.Millisecond,
					Transport: Transport{
						DisableCompression:    false,
						DisableKeepAlives:     false,
						ExpectContinueTimeout: 1 * time.Second,
						ForceAttemptHTTP2:     true,
						IdleConnTimeout:       90 * time.Second,
						MaxConnsPerHost:       100,
						MaxIdleConns:          100,
						MaxIdleConnsPerHost:   2,
						ResponseHeaderTimeout: 2 * time.Second,
						TLSHandshakeTimeout:   10 * time.Second,
					},
					Plugins: Plugins{
						Datadog: Datadog{
							Enabled:       false,
							OperationName: "",
						},
						Health: Health{
							Enabled:     false,
							Description: "default connection",
							Endpoint:    "/resource-status",
							Host:        "",
							Name:        "rest api",
							Required:    true,
						},
						Log: Log{
							Enabled: false,
							Level:   "DEBUG",
						},
						Newrelic: Newrelic{
							Enabled: false,
						},
						Opentracing: Opentracing{
							Enabled: false,
						},
						RequestID: RequestID{
							Enabled: true,
						},
						Retry: Retry{
							Enabled:     true,
							Count:       0,
							MaxWaitTime: 2 * time.Second,
							WaitTime:    200 * time.Millisecond,
						},
					},
				}

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
