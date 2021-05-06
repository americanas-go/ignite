package newrelic

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

const NewRelicTransaction = "__newrelic_transaction__"

func FromContext(ctx context.Context) *newrelic.Transaction {
	txn := newrelic.FromContext(ctx)
	if txn == nil {
		return ctx.Value(NewRelicTransaction).(*newrelic.Transaction)
	}
	return txn
}
