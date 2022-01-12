package ignite

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		run  func() any
		want any
	}{
		{
			name: "Returns new instance of *CustomOptions",
			run:  func() any { return New[*CustomOptions]() },
			want: &CustomOptions{},
		},
		{
			name: "Returns new instance of CustomOptions",
			run:  func() any { return New[CustomOptions]() },
			want: CustomOptions{},
		},
		{
			name: "Returns new instance of *int",
			run:  func() any { return New[*int]() },
			want: new(int),
		},
		{
			name: "Returns new instance of int",
			run:  func() any { return New[int]() },
			want: 0,
		},
		{
			name: "Returns new instance of *string",
			run:  func() any { return New[*string]() },
			want: new(string),
		},
		{
			name: "Returns new instance of string",
			run:  func() any { return New[string]() },
			want: "",
		},
		{
			name: "Returns new instance of *map[string]string",
			run:  func() any { return New[*map[string]string]() },
			want: new(map[string]string),
		},
		{
			name: "Returns new instance of map[string]string",
			run:  func() any { return New[map[string]string]() },
			want: map[string]string{},
		},
		{
			name: "Returns new instance of []*string",
			run:  func() any { return New[[]*string]() },
			want: []*string{},
		},
		{
			name: "Returns new instance of []string",
			run:  func() any { return New[[]string]() },
			want: []string{},
		},
		{
			name: "Returns new instance of []*float64",
			run:  func() any { return New[[]*float64]() },
			want: []*float64{},
		},
		{
			name: "Returns new instance of []float64",
			run:  func() any { return New[[]float64]() },
			want: []float64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.run()
			if got == nil || reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("\nwant\t%v of type %T\ngot \t%v of type %T", tt.want, tt.want, got, got)
			}
		})
	}
}
