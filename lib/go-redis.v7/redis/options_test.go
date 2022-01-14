package redis

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
					Password:           "",
					MaxRetries:         0,
					MinRetryBackoff:    8 * time.Millisecond,
					MaxRetryBackoff:    512 * time.Millisecond,
					DialTimeout:        5 * time.Second,
					ReadTimeout:        3 * time.Second,
					WriteTimeout:       3 * time.Second,
					PoolSize:           10,
					MinIdleConns:       2,
					MaxConnAge:         0,
					PoolTimeout:        4 * time.Second,
					IdleTimeout:        5 * time.Minute,
					IdleCheckFrequency: 1 * time.Minute,
					Client: ClientOptions{
						Addr:    "127.0.0.1:6379",
						Network: "tcp",
						DB:      0,
					},
					Cluster: ClusterOptions{
						Enabled:        false,
						Addrs:          []string{"127.0.0.1:6379"},
						MaxRedirects:   8,
						ReadOnly:       false,
						RouteByLatency: false,
						RouteRandomly:  false,
					},
					Sentinel: SentinelOptions{
						MasterName: "",
						Addrs:      []string{},
						Password:   "",
					},
					Plugins: PluginsOptions{
						Datadog: DatadogOptions{
							Enabled:       false,
							ServiceName:   "Redis",
							AnalyticsRate: -1,
						},
						Health: HealthOptions{
							Enabled:     true,
							Name:        "Redis",
							Description: "default connection",
							Required:    true,
						},
						Newrelic: NewrelicOptions{
							Enabled: false,
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
