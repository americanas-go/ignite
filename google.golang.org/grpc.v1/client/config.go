package client

import (
	"time"

	"github.com/americanas-go/config"
)

const (
	root                         = "ignite.grpc.client"
	PluginsRoot                  = root + ".plugins"
	host                         = ".host"
	block                        = ".block"
	initialWindowSize            = ".initialWindowSize"
	initialConnWindowSize        = ".initialConnWindowSize"
	tlsRoot                      = ".tls"
	tlsEnabled                   = tlsRoot + ".enabled"
	certFile                     = tlsRoot + ".certFile"
	keyFile                      = tlsRoot + ".keyFile"
	caFile                       = tlsRoot + ".caFile"
	insecureSkipVerify           = tlsRoot + ".insecureSkipVerify"
	hostOverwrite                = ".hostOverwrite"
	port                         = ".port"
	keepaliveRoot                = ".keepalive"
	keepaliveTime                = keepaliveRoot + ".time"
	keepaliveTimeout             = keepaliveRoot + ".timeout"
	keepalivePermitWithoutStream = keepaliveRoot + ".permitWithoutStream"
	connectParamsRoot            = ".connectParams"
	minConnectTimeout            = connectParamsRoot + ".minConnectTimeout"

	backoffRoot       = connectParamsRoot + ".backoff"
	backoffBaseDelay  = backoffRoot + ".baseDelay"
	backoffMultiplier = backoffRoot + ".multiplier"
	backoffJitter     = backoffRoot + ".jitter"
	backoffMaxDelay   = backoffRoot + ".maxDelay"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+host, "localhost", "defines host")
	config.Add(path+port, 9091, "defines port")
	config.Add(path+block, false, "makes caller of Dial blocks until the underlying connection is up. Without this, Dial returns immediately and connecting the server happens in background")
	config.Add(path+initialWindowSize, 1024*1024*2, "sets the initial window size for a stream")
	config.Add(path+initialConnWindowSize, 1024*1024*2, "sets the initial window size for a connection")
	config.Add(path+tlsEnabled, false, "enable/disable tls")
	config.Add(path+certFile, "", "defines cert file")
	config.Add(path+keyFile, "", "defines key file")
	config.Add(path+caFile, "", "defines ca file")
	config.Add(path+hostOverwrite, "", "defines host overwrite")
	config.Add(path+insecureSkipVerify, true, "enable/disable insecure skip verify ")
	config.Add(path+keepaliveTime, 10*time.Second, "After a duration of this time if the client doesn't see any activity it pings the server to see if the transport is still alive. If set below 10s, a minimum value of 10s will be used instead")
	config.Add(path+keepaliveTimeout, 20*time.Second, "After having pinged for keepalive check, the client waits for a duration of Timeout and if no activity is seen even after that the connection is closed")
	config.Add(path+keepalivePermitWithoutStream, false, "If true, client sends keepalive pings even with no active RPCs. If false, when there are no active RPCs, Time and Timeout will be ignored and no keepalive pings will be sent.")
	config.Add(path+minConnectTimeout, 20*time.Second, "is the minimum amount of time we are willing to give a connection to complete")
	config.Add(path+backoffBaseDelay, 1.0*time.Second, "BaseDelay is the amount of time to backoff after the first failure")
	config.Add(path+backoffMultiplier, 1.6, "is the factor with which to multiply backoffs after a failed retry. Should ideally be greater than 1")
	config.Add(path+backoffJitter, 0.2, "Jitter is the factor with which backoffs are randomized")
	config.Add(path+backoffMaxDelay, 120*time.Second, "MaxDelay is the upper bound of backoff delay")
}
