package aws

import (
	"context"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		ctx     context.Context
		options *Options
	}
	tests := []struct {
		name string
		args args
		want aws.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfigWithOptions(tt.args.ctx, tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfigWithOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDefaultConfig(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want aws.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
