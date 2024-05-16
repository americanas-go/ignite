package azure

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"testing"
)

func TestNewCredentialWithOptions(t *testing.T) {
	type args struct {
		ctx     context.Context
		options *Options
	}
	tests := []struct {
		name string
		args args
		want azcore.TokenCredential
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := NewCredentialWithOptions(tt.args.ctx, tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCredentialWithOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCredential(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want azcore.TokenCredential
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := NewCredential(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestNewCredential() = %v, want %v", got, tt.want)
			}
		})
	}
}
