package gocql

import (
	"time"

	"github.com/americanas-go/config"
)

const (
	root                     = "ignite.gocql"
	hosts                    = root + ".hosts"
	port                     = root + ".port"
	dc                       = root + ".dc"
	username                 = root + ".username"
	password                 = root + ".password"
	cqlVersion               = root + ".CQLVersion"
	protoVersion             = root + ".protoVersion"
	timeout                  = root + ".timeout"
	connectTimeout           = root + ".connecttimeout"
	keyspace                 = root + ".keyspace"
	numConns                 = root + ".numConns"
	consistency              = root + ".consistency"
	socketKeepalive          = root + ".socketKeepAlive"
	maxPreparedStmts         = root + ".maxPreparedStmts"
	maxRoutingKeyInfo        = root + ".maxRoutingKeyInfo"
	pageSize                 = root + ".pageSize"
	defaultTimestamp         = root + ".defaultTimestamp"
	reconnectInterval        = root + ".reconnectInterval"
	maxWaitSchemaAgreement   = root + ".maxWaitSchemaAgreement"
	disableInitialHostLookup = root + ".disableInitialHostLookup"
	writeCoalesceWaitTime    = root + ".writeCoalesceWaitTime"
	PluginsRoot              = root + ".plugins"
)

func init() {

	config.Add(hosts, []string{"127.0.0.1"}, "addresses for the initial connections")
	config.Add(port, 9042, "define port")
	config.Add(dc, "", "define DC")
	config.Add(username, "", "define username")
	config.Add(password, "", "define password", config.WithHide())
	config.Add(cqlVersion, "3.0.0", "define cql version")
	config.Add(protoVersion, 0, "define version of the native protocol to use")
	config.Add(timeout, 600*time.Millisecond, "connection timeout")
	config.Add(connectTimeout, 600*time.Millisecond, "initial connection timeout, used during initial dial to server")
	config.Add(keyspace, "", "initial keyspace (optional)")
	config.Add(numConns, 2, "number of connections per host")
	config.Add(consistency, "QUORUM", "default consistency level (default: Quorum) (values: ANY, ONE, TWO, THREE, QUORUM, ALL, LOCAL_QUORUM, EACH_QUORUM, LOCAL_ONE)")
	config.Add(socketKeepalive, 0*time.Millisecond, "The keepalive period to use, enabled if > 0 (default: 0)")
	config.Add(maxPreparedStmts, 1000, "Sets the maximum cache size for prepared statements globally for gocql")
	config.Add(maxRoutingKeyInfo, 1000, "Sets the maximum cache size for query info about statements for each session")
	config.Add(pageSize, 5000, "Default page size to use for created sessions")
	config.Add(defaultTimestamp, true, "Sends a client side timestamp for all requests which overrides the timestamp at which it arrives at the server. (default: true, only enabled for protocol 3 and above)")
	config.Add(reconnectInterval, 10*time.Millisecond, "If not zero, gocql attempt to reconnect known DOWN nodes in every ReconnectInterval")
	config.Add(maxWaitSchemaAgreement, 60*time.Second, "The maximum amount of time to wait for schema agreement in a cluster after receiving a schema change frame")
	config.Add(disableInitialHostLookup, true, "If true then the driver will not attempt to get host info from the system.peers table")
	config.Add(writeCoalesceWaitTime, 200*time.Microsecond, "The time to wait for frames before flushing the frames connection to Cassandra")
}
