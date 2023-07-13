package kafka

import (
	"github.com/americanas-go/config"
	"github.com/segmentio/kafka-go"
)

const (
	root          = "ignite.kafka"
	address       = ".address"
	topic         = ".topic"
	partition     = ".partition"
	network       = ".network"
	connType      = ".connType"
	queueCapacity = ".queueCapacity"
	minBytes      = ".minBytes"
	maxBytes      = ".maxBytes"
	startOffset   = ".startOffset"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+address, "localhost:9092", "defines host address")
	config.Add(path+topic, "", "defines topic name")
	config.Add(path+partition, 0, "defines partition number")
	config.Add(path+network, "tcp", "defines network protocol")
	config.Add(path+connType, "LEADER", "defines connection type. LEADER, PARTITION or SERVER")
	config.Add(path+queueCapacity, 100, "defines queue capacity")
	config.Add(path+minBytes, 1, "defines batch min bytes")
	config.Add(path+maxBytes, 10485760, "defines batch max bytes")
	config.Add(path+startOffset, kafka.LastOffset, "defines start offset LastOffset=-1, FirstOffset=-2")
}
