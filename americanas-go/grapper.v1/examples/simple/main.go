package main

import (
	"context"
	"os"

	"github.com/americanas-go/ignite"
	"github.com/americanas-go/ignite/americanas-go/grapper.v1"
	"github.com/americanas-go/ignite/americanas-go/grapper.v1/plugins/contrib/afex/hystrix-go.v0"
	logger "github.com/americanas-go/ignite/americanas-go/grapper.v1/plugins/contrib/americanas-go/log.v1"
	"github.com/americanas-go/ignite/rs/zerolog.v1"
	"github.com/americanas-go/log"
)

func init() {
	os.Setenv("IGNITE_ZEROLOG_LEVEL", "TRACE")
}

func main() {

	ctx := context.Background()

	ignite.Boot()
	zerolog.NewLogger()

	var r string
	var err error

	wrp, _ := grapper.NewAnyErrorWrapper[string](ctx, logger.NewAnyError[string], hystrix.NewAnyError[string])

	r, err = wrp.Exec(ctx, "xpto",
		func(ctx context.Context) (string, error) {
			l := log.FromContext(ctx)
			l.Info("executed business rule")
			return "string", nil
		}, nil)

	if err != nil {
		log.Errorf(err.Error())
	}

	log.Infof(r)
}
