package kafka

import "github.com/americanas-go/config"

const (
	root      = "ignite.kafka"
	address   = ".address"
	topic     = ".topic"
	partition = ".partition"
	network   = ".network"
	connType  = ".connType"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+address, "localhost:9092", "defines host address")
	config.Add(path+topic, "", "defines topic name")
	config.Add(path+partition, 0, "defines partition number")
	config.Add(path+network, "tcp", "defines network protocol")
	config.Add(path+connType, "LEADER", "defines connectio type. LEADER, PARTITION or SERVER")
}
