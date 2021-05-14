package newrelic

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

const NewRelicTransaction = "__newrelic_transaction__"

func FromContext(ctx context.Context) *newrelic.Transaction {
	txn := newrelic.FromContext(ctx)
	if txn == nil {
		if txn, ok := ctx.Value(NewRelicTransaction).(*newrelic.Transaction); ok {
			return txn
		}
	}
	return txn
}
