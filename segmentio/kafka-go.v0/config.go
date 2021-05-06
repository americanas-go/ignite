package kafka

import "github.com/americanas-go/config"

const (
	root      = "ignite.kafka"
	address   = root + ".address"
	topic     = root + ".topic"
	partition = root + ".partition"
	network   = root + ".network"
	connType  = root + ".connType"
)

func init() {
	config.Add(address, "localhost:9092", "defines host address")
	config.Add(topic, "", "defines topic name")
	config.Add(partition, 0, "defines partition number")
	config.Add(network, "tcp", "defines network protocol")
	config.Add(connType, "LEADER", "defines connectio type. LEADER, PARTITION or SERVER")
}
