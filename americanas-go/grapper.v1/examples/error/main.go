package main

import (
	"context"
	"os"

	"github.com/americanas-go/errors"
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

	wrp := grapper.New[string]("example", logger.New[string], hystrix.New[string])

	r, err = wrp.Exec(ctx,
		func(ctx context.Context) (string, error) {
			l := log.FromContext(ctx)
			l.Info("executed business rule with error")
			return "", errors.New("an error ocurred")
		},
		func(ctx context.Context, err error) (string, error) {
			l := log.FromContext(ctx)
			if err != nil {
				l.Info("executed fallback business rule")
				return "string", nil
			}
			return "", err
		})

	if err != nil {
		log.Errorf(err.Error())
	}

	log.Infof(r)
}
