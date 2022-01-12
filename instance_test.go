package ignite

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/americanas-go/config"
)

var (
	pluginError = errors.New("Ops! This should not be happening!")
)

// custom client
type CustomClient struct {
	Plugins []Plugin[*CustomClient, *CustomClientOptions]
	Options *CustomClientOptions
}

// custom client Init
func (tc *CustomClient) Init(ctx context.Context, o *CustomClientOptions) error {
	tc.Options = o
	return nil
}

type CustomClientOptions struct {
	Enabled bool
	Host    string
	Port    int
	Plugins struct {
		Custom struct {
			Enabled    bool
			MaxRetries int
		}
	}
}

func (o *CustomClientOptions) Root() string {
	return "ignite.customClient"
}

func (o *CustomClientOptions) PostLoad() error {
	return nil
}

func CustomPlugin(ctx context.Context, tc *CustomClient) error {
	tc.Plugins = append(tc.Plugins, CustomPlugin)
	// plugin code
	return nil
}

func CustomPluginWithError(ctx context.Context, tc *CustomClient) error {
	return pluginError
}

func TestSetup(t *testing.T) {
	tests := []struct {
		name    string
		plugins []Plugin[*CustomClient, *CustomClientOptions]
		want    func(*CustomClient) bool
		wantErr func(error) bool
	}{
		{
			name: "Returns new ignite test client with set options and plugins",
			plugins: []Plugin[*CustomClient, *CustomClientOptions]{
				CustomPlugin,
			},
			want: func(c *CustomClient) bool {
				return c != nil && len(c.Plugins) == 1
			},
			wantErr: func(e error) bool { return e == nil },
		},
		{
			name: "Returns error when registering plugin ",
			plugins: []Plugin[*CustomClient, *CustomClientOptions]{
				CustomPluginWithError,
			},
			want: func(tc *CustomClient) bool {
				return tc != nil
			},
			wantErr: func(e error) bool { return e != nil && e == pluginError },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tc, e := Setup(context.Background(), tt.plugins...)
			if !tt.want(tc) {
				t.Errorf("Unexpected value %v", tc)
			}
			if !tt.wantErr(e) {
				t.Errorf("Unexpected error %v", e)
			}
		})
	}
}

func TestSetupWithConfigPath(t *testing.T) {
	enabled := true
	host := "localhost"
	port := 9999
	os.Setenv("CUSTOM_PATH_ENABLED", fmt.Sprint(enabled))
	os.Setenv("CUSTOM_PATH_HOST", host)
	os.Setenv("CUSTOM_PATH_PORT", fmt.Sprint(port))
	config.Load()
	tests := []struct {
		name    string
		plugins []Plugin[*CustomClient, *CustomClientOptions]
		want    func(*CustomClient) bool
		wantErr func(error) bool
	}{
		{
			name: "Returns new ignite test client with set options and plugins",
			plugins: []Plugin[*CustomClient, *CustomClientOptions]{
				CustomPlugin,
			},
			want: func(c *CustomClient) bool {
				return c != nil && len(c.Plugins) == 1 && c.Options.Enabled == enabled &&
					c.Options.Host == host && c.Options.Port == port
			},
			wantErr: func(e error) bool { return e == nil },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tc, e := SetupWithConfigPath(context.Background(), "custom.path", tt.plugins...)
			if !tt.want(tc) {
				t.Errorf("Unexpected value %v", tc)
			}
			if !tt.wantErr(e) {
				t.Errorf("Unexpected error %v", e)
			}
		})
	}
}
