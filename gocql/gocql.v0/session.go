package gocql

import (
	"context"
	"strings"

	"github.com/americanas-go/log"
	"github.com/gocql/gocql"
)

type Plugin func(context.Context, *gocql.Session) error

func NewSession(ctx context.Context, plugins ...Plugin) (*gocql.Session, error) {

	logger := log.FromContext(ctx)

	o, err := NewOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewSessionWithOptions(ctx, o, plugins...)
}

func NewSessionWithOptions(ctx context.Context, o *Options, plugins ...Plugin) (session *gocql.Session, err error) {

	logger := log.FromContext(ctx)

	cluster := gocql.NewCluster(o.Hosts...)

	if o.Port > 0 {
		cluster.Port = o.Port
	}

	if o.CQLVersion != "" {
		cluster.CQLVersion = o.CQLVersion
	}

	if o.ProtoVersion > 0 {
		cluster.ProtoVersion = o.ProtoVersion
	}

	if o.Timeout > 0 {
		cluster.Timeout = o.Timeout
	}

	if o.ConnectTimeout > 0 {
		cluster.ConnectTimeout = o.ConnectTimeout
	}

	if o.Keyspace != "" {
		cluster.Keyspace = o.Keyspace
	}

	if o.NumConns > 0 {
		cluster.NumConns = o.NumConns
	}

	if o.SocketKeepalive > 0 {
		cluster.SocketKeepalive = o.SocketKeepalive
	}

	if o.MaxPreparedStmts > 0 {
		cluster.MaxPreparedStmts = o.MaxPreparedStmts
	}

	if o.MaxRoutingKeyInfo > 0 {
		cluster.MaxRoutingKeyInfo = o.MaxRoutingKeyInfo
	}

	if o.PageSize > 0 {
		cluster.PageSize = o.PageSize
	}

	cluster.DefaultTimestamp = o.DefaultTimestamp

	if o.ReconnectInterval > 0 {
		cluster.ReconnectInterval = o.ReconnectInterval
	}

	if o.MaxWaitSchemaAgreement > 0 {
		cluster.MaxWaitSchemaAgreement = o.MaxWaitSchemaAgreement
	}

	cluster.DisableInitialHostLookup = o.DisableInitialHostLookup

	if o.WriteCoalesceWaitTime > 0 {
		cluster.WriteCoalesceWaitTime = o.WriteCoalesceWaitTime
	}

	if o.Username != "" || o.Password != "" {
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: o.Username,
			Password: o.Password,
		}
	}

	if o.Consistency != "" {
		cluster.Consistency = gocql.ParseConsistency(o.Consistency)
	}

	if o.DC != "" {
		cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy(o.DC)
		cluster.HostFilter = gocql.DataCentreHostFilter(o.DC)
	}

	session, err = cluster.CreateSession()

	if err != nil {
		return nil, err
	}

	for _, plugin := range plugins {
		if err := plugin(ctx, session); err != nil {
			panic(err)
		}
	}

	logger.Infof("Connected to Cassandra server: %v", strings.Join(o.Hosts, ","))

	return session, err
}
