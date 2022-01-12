package ignite

import (
	"os"
	"reflect"
	"testing"
)

type CustomOptions struct {
	Enabled bool   `default:"true" desc:"whether enabled"`
	Host    string `default:"127.0.0.1" desc:"app host"`
	Port    int    `default:"7777" desc:"app port"`
	Plugins struct {
		Custom struct {
			Enabled bool     `default:"true" desc:"custom plugin options"`
			Count   int      `config:"counterNumber" default:"777" desc:"custom plugin options"`
			Servers []string `default:"localhost,10.0.1.1" desc:"custom plugin options"`
			Ports   []int    `default:"9999,8888,8080" desc:"custom plugin options"`
		}
	} `config:"middlewares"`
}

func (co *CustomOptions) Root() string {
	return "ignite.custom"
}

func (co *CustomOptions) PostLoad() error {
	return nil
}

func TestLoadOptions(t *testing.T) {
	os.Setenv("IGNITE_CUSTOM_ENABLED", "true")
	os.Setenv("IGNITE_CUSTOM_HOST", "localhost")
	os.Setenv("IGNITE_CUSTOM_PORT", "9999")
	os.Setenv("IGNITE_CUSTOM_MIDDLEWARES_CUSTOM_ENABLED", "true")
	os.Setenv("IGNITE_CUSTOM_MIDDLEWARES_CUSTOM_COUNTER__NUMBER", "18")
	defer os.Clearenv()
	tests := []struct {
		name    string
		want    func() *CustomOptions
		wantErr func(error) bool
	}{
		{
			name: "Returns new ignite options",
			want: func() *CustomOptions {
				co := New[*CustomOptions]()
				co.Enabled = true
				co.Host = "localhost"
				co.Port = 9999
				co.Plugins.Custom.Enabled = true
				co.Plugins.Custom.Count = 18
				co.Plugins.Custom.Servers = []string{
					"localhost",
					"10.0.1.1",
				}
				co.Plugins.Custom.Ports = []int{
					9999,
					8888,
					8080,
				}
				return co
			},
			wantErr: func(e error) bool { return e == nil },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, e := Load[*CustomOptions]()
			want := tt.want()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("\nwant\t%v\ngot \t%v", want, got)
			}
			if !tt.wantErr(e) {
				t.Errorf("Unexpected error %v", e)
			}
		})
	}
}

func TestLoadOptionsWithPath(t *testing.T) {
	tests := []struct {
		name    string
		want    func() *CustomOptions
		wantErr func(error) bool
	}{
		{
			name: "Returns new ignite options",
			want: func() *CustomOptions {
				co := New[*CustomOptions]()
				co.Enabled = true
				co.Host = "127.0.0.1"
				co.Port = 7777
				co.Plugins.Custom.Enabled = true
				co.Plugins.Custom.Count = 777
				co.Plugins.Custom.Servers = []string{
					"localhost",
					"10.0.1.1",
				}
				co.Plugins.Custom.Ports = []int{
					9999,
					8888,
					8080,
				}
				return co
			},
			wantErr: func(e error) bool { return e == nil },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, e := LoadWithPath[*CustomOptions]("test.custom")
			want := tt.want()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("\nwant\t%v\ngot \t%v", want, got)
			}
			if !tt.wantErr(e) {
				t.Errorf("Unexpected error %v", e)
			}
		})
	}
}
