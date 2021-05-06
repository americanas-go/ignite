package fx

import (
	"go.uber.org/fx"
)

func NewApp(opts ...fx.Option) *fx.App {
	opts = append([]fx.Option{fx.Logger(NewLogger())}, opts...)
	return fx.New(opts...)
}
