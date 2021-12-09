package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/americanas-go/log"
	"github.com/wesovilabs/beyond/api"
	c "github.com/wesovilabs/beyond/api/context"
)

type TracingAdvice struct {
	prefix string
}

func (a *TracingAdvice) Before(ctx *c.BeyondContext) {

	params := make([]string, ctx.Params().Count())
	ctx.Params().ForEach(func(index int, arg *c.Arg) {
		params[index] = fmt.Sprintf("%s:%v", arg.Name(), arg.Value())
	})
	printTrace(ctx, a.prefix, params)
}

func NewTracingAdvice() api.Before {
	return &TracingAdvice{}
}

func NewTracingAdviceWithPrefix(prefix string) func() api.Before {
	return func() api.Before {
		return &TracingAdvice{
			prefix: prefix,
		}
	}
}

func printTrace(ctx *c.BeyondContext, prefix string, params []string) {

	logger := log.FromContext(context.Background())

	if prefix == "" {
		logger.Infof("%s.%s(%s)\n", ctx.Pkg(), ctx.Function(), strings.Join(params, ","))
		return
	}
	logger.Infof("%s %s.%s(%s)\n", prefix, ctx.Pkg(), ctx.Function(), strings.Join(params, ","))
}
