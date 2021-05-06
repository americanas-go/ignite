package godror

import (
	"context"
	"database/sql"

	"github.com/americanas-go/log"
	_ "github.com/godror/godror"
)

type Plugin func(context.Context, *sql.DB) error

func NewDBWithOptions(ctx context.Context, o *Options, plugins ...Plugin) (db *sql.DB, err error) {

	logger := log.FromContext(ctx)

	db, err = sql.Open("godror", o.DataSourceName)
	if err != nil {
		return nil, err
	}
	// defer db.Close()

	db.SetConnMaxLifetime(o.ConnMaxLifetime)
	db.SetMaxIdleConns(o.MaxIdleConns)
	db.SetMaxOpenConns(o.MaxOpenConns)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	for _, plugin := range plugins {
		if err := plugin(ctx, db); err != nil {
			panic(err)
		}
	}

	logger.Info("Connected to Oracle (godror) server")

	return db, err
}

func NewDB(ctx context.Context, plugins ...Plugin) (*sql.DB, error) {

	logger := log.FromContext(ctx)

	o, err := NewOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewDBWithOptions(ctx, o, plugins...)
}
