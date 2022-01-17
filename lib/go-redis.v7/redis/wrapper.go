package redis

import (
	"context"
	"strings"

	"github.com/americanas-go/log"
	"github.com/go-redis/redis/v7"
)

// wraps redis client, cluster client.
type Wrapper struct {
	Client        *redis.Client
	ClusterClient *redis.ClusterClient
	Options       *Options
}

// Init redis client or cluster client
func (w *Wrapper) Init(ctx context.Context, o *Options) error {
	w.Options = o
	if o.Cluster.Enabled {
		return w.initCluster(ctx)
	}
	return w.initClient(ctx)
}

// returns the client or cluster client as universal client
func (w *Wrapper) UniversalClient() redis.UniversalClient {
	if w.Options.Cluster.Enabled {
		return w.ClusterClient
	}
	return w.Client
}

// returns redis options according to client or cluster client
func (w *Wrapper) RedisOptions() *redis.Options {
	if w.Options.Cluster.Enabled {
		o := w.Options
		return &redis.Options{
			Addr:               o.Cluster.Addrs[0],
			Network:            "tcp",
			Password:           o.Password,
			MaxRetries:         o.MaxRetries,
			MinRetryBackoff:    o.MinRetryBackoff,
			MaxRetryBackoff:    o.MaxRetryBackoff,
			DialTimeout:        o.DialTimeout,
			DB:                 0,
			ReadTimeout:        o.ReadTimeout,
			WriteTimeout:       o.WriteTimeout,
			PoolSize:           o.PoolSize,
			MinIdleConns:       o.MinIdleConns,
			MaxConnAge:         o.MaxConnAge,
			PoolTimeout:        o.PoolTimeout,
			IdleTimeout:        o.IdleTimeout,
			IdleCheckFrequency: o.IdleCheckFrequency,
		}
	}
	return w.Client.Options()
}

func (w *Wrapper) initCluster(ctx context.Context) error {
	logger := log.FromContext(ctx)
	o := w.Options
	w.ClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              o.Cluster.Addrs,
		MaxRedirects:       o.Cluster.MaxRedirects,
		ReadOnly:           o.Cluster.ReadOnly,
		RouteByLatency:     o.Cluster.RouteByLatency,
		RouteRandomly:      o.Cluster.RouteRandomly,
		Password:           o.Password,
		MaxRetries:         o.MaxRetries,
		MinRetryBackoff:    o.MinRetryBackoff,
		MaxRetryBackoff:    o.MaxRetryBackoff,
		DialTimeout:        o.DialTimeout,
		ReadTimeout:        o.ReadTimeout,
		WriteTimeout:       o.WriteTimeout,
		PoolSize:           o.PoolSize,
		MinIdleConns:       o.MinIdleConns,
		MaxConnAge:         o.MaxConnAge,
		PoolTimeout:        o.PoolTimeout,
		IdleTimeout:        o.IdleTimeout,
		IdleCheckFrequency: o.IdleCheckFrequency,
	})

	ping := w.ClusterClient.Ping()
	if ping.Err() != nil {
		return ping.Err()
	}

	logger.Infof("Connected to Redis Cluster server: %s status: %s", strings.Join(w.ClusterClient.Options().Addrs, ","), ping.String())

	return nil
}

func (w *Wrapper) initClient(ctx context.Context) error {
	o := w.Options
	logger := log.FromContext(ctx)

	if redisSentinel(o) {
		w.Client = failOverClient(o)
	} else {
		w.Client = standaloneClient(o)
	}

	ping := w.Client.Conn().Ping()
	if ping.Err() != nil {
		return ping.Err()
	}

	logger.Infof("Connected to Redis server: %s %s", w.Client.Options().Addr, ping.String())

	return nil
}

func failOverClient(o *Options) *redis.Client {
	return redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:         o.Sentinel.MasterName,
		SentinelAddrs:      o.Sentinel.Addrs,
		SentinelPassword:   o.Sentinel.Password,
		Password:           o.Password,
		MaxRetries:         o.MaxRetries,
		MinRetryBackoff:    o.MinRetryBackoff,
		MaxRetryBackoff:    o.MaxRetryBackoff,
		DialTimeout:        o.DialTimeout,
		DB:                 o.Client.DB,
		ReadTimeout:        o.ReadTimeout,
		WriteTimeout:       o.WriteTimeout,
		PoolSize:           o.PoolSize,
		MinIdleConns:       o.MinIdleConns,
		MaxConnAge:         o.MaxConnAge,
		PoolTimeout:        o.PoolTimeout,
		IdleTimeout:        o.IdleTimeout,
		IdleCheckFrequency: o.IdleCheckFrequency,
	})
}

func standaloneClient(o *Options) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:               o.Client.Addr,
		Network:            o.Client.Network,
		Password:           o.Password,
		MaxRetries:         o.MaxRetries,
		MinRetryBackoff:    o.MinRetryBackoff,
		MaxRetryBackoff:    o.MaxRetryBackoff,
		DialTimeout:        o.DialTimeout,
		DB:                 o.Client.DB,
		ReadTimeout:        o.ReadTimeout,
		WriteTimeout:       o.WriteTimeout,
		PoolSize:           o.PoolSize,
		MinIdleConns:       o.MinIdleConns,
		MaxConnAge:         o.MaxConnAge,
		PoolTimeout:        o.PoolTimeout,
		IdleTimeout:        o.IdleTimeout,
		IdleCheckFrequency: o.IdleCheckFrequency,
	})
}

func redisSentinel(o *Options) bool {
	return o.Sentinel.MasterName != "" || len(o.Sentinel.Addrs) > 0
}
