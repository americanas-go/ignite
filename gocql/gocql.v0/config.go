package gocql

import (
	"time"

	"github.com/americanas-go/config"
)

const (
	root                     = "ignite.gocql"
	hosts                    = ".hosts"
	port                     = ".port"
	dc                       = ".dc"
	pu                       = ".username"
	pp                       = ".password"
	cqlVersion               = ".CQLVersion"
	protoVersion             = ".protoVersion"
	timeout                  = ".timeout"
	connectTimeout           = ".connecttimeout"
	keyspace                 = ".keyspace"
	numConns                 = ".numConns"
	consistency              = ".consistency"
	socketKeepalive          = ".socketKeepAlive"
	maxPreparedStmts         = ".maxPreparedStmts"
	maxRoutingKeyInfo        = ".maxRoutingKeyInfo"
	pageSize                 = ".pageSize"
	defaultTimestamp         = ".defaultTimestamp"
	reconnectInterval        = ".reconnectInterval"
	maxWaitSchemaAgreement   = ".maxWaitSchemaAgreement"
	disableInitialHostLookup = ".disableInitialHostLookup"
	writeCoalesceWaitTime    = ".writeCoalesceWaitTime"
	PluginsRoot              = root + ".plugins"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+hosts, []string{"127.0.0.1"}, "addresses for the initial connections")
	config.Add(path+port, 9042, "define port")
	config.Add(path+dc, "", "define DC")
	config.Add(path+pu, "", "define username")
	config.Add(path+pp, "", "define password", config.WithHide())
	config.Add(path+cqlVersion, "3.0.0", "define cql version")
	config.Add(path+protoVersion, 0, "define version of the native protocol to use")
	config.Add(path+timeout, 600*time.Millisecond, "connection timeout")
	config.Add(path+connectTimeout, 600*time.Millisecond, "initial connection timeout, used during initial dial to server")
	config.Add(path+keyspace, "", "initial keyspace (optional)")
	config.Add(path+numConns, 2, "number of connections per host")
	config.Add(path+consistency, "QUORUM", "default consistency level (default: Quorum) (values: ANY, ONE, TWO, THREE, QUORUM, ALL, LOCAL_QUORUM, EACH_QUORUM, LOCAL_ONE)")
	config.Add(path+socketKeepalive, 0*time.Millisecond, "The keepalive period to use, enabled if > 0 (default: 0)")
	config.Add(path+maxPreparedStmts, 1000, "Sets the maximum cache size for prepared statements globally for gocql")
	config.Add(path+maxRoutingKeyInfo, 1000, "Sets the maximum cache size for query info about statements for each session")
	config.Add(path+pageSize, 5000, "Default page size to use for created sessions")
	config.Add(path+defaultTimestamp, true, "Sends a client side timestamp for all requests which overrides the timestamp at which it arrives at the server. (default: true, only enabled for protocol 3 and above)")
	config.Add(path+reconnectInterval, 10*time.Millisecond, "If not zero, gocql attempt to reconnect known DOWN nodes in every ReconnectInterval")
	config.Add(path+maxWaitSchemaAgreement, 60*time.Second, "The maximum amount of time to wait for schema agreement in a cluster after receiving a schema change frame")
	config.Add(path+disableInitialHostLookup, true, "If true then the driver will not attempt to get host info from the system.peers table")
	config.Add(path+writeCoalesceWaitTime, 200*time.Microsecond, "The time to wait for frames before flushing the frames connection to Cassandra")
}
