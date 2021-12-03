package memdb

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/hashicorp/go-memdb"
)

// NewMemDB returns a new MemDB with the schema.
func NewMemDB(ctx context.Context, schema *memdb.DBSchema) (db *memdb.MemDB, err error) {

	logger := log.FromContext(ctx)

	db, err = memdb.NewMemDB(schema)
	if err != nil {
		return nil, err
	}

	logger.Info("Connected to go-memdb")

	return db, err

}
