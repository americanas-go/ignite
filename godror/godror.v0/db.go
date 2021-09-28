package godror

import (
	"context"
	"database/sql"

	"github.com/americanas-go/ignite/time"
	"github.com/americanas-go/log"
	"github.com/godror/godror"
)

type Plugin func(context.Context, *sql.DB) error

func NewDBWithConfigPath(ctx context.Context, path string, plugins ...Plugin) (*sql.DB, error) {
	opts, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewDBWithOptions(ctx, opts, plugins...)
}

func NewDBWithOptions(ctx context.Context, o *Options, plugins ...Plugin) (db *sql.DB, err error) {

	logger := log.FromContext(ctx)

	var P godror.ConnectionParams
	P.ConnectString = o.ConnectString
	if o.Username != "" && o.Password != "" {
		P.Username, P.Password = o.Username, godror.NewPassword(o.Password)
	}
	P.SessionTimeout = o.SessionTimeout
	P.MaxLifeTime = o.MaxLifetime
	P.MaxSessions = o.MaxSessions
	P.MinSessions = o.MinSessions
	P.MaxSessionsPerShard = o.MaxSessionsPerShard
	P.Timezone = time.Location()
	P.WaitTimeout = o.WaitTimeout
	P.SessionIncrement = o.SessionIncrement
	// P.SetSessionParamOnInit("NLS_NUMERIC_CHARACTERS", ",.")
	// P.SetSessionParamOnInit("NLS_LANGUAGE", "FRENCH")

	db = sql.OpenDB(godror.NewConnector(P))

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
