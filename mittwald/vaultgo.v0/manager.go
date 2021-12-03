package vault

import "context"

// Manager vault manager interface.
type Manager interface {
	Options() *ManagerOptions
	Close(context.Context) error
	Configure(context.Context, map[string]interface{}) error
}
