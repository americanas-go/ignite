package vault

import "context"

type Manager interface {
	Options() *ManagerOptions
	Close(context.Context) error
	Configure(context.Context, map[string]interface{}) error
}
