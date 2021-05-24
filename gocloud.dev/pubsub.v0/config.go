package pubsub

import "github.com/americanas-go/config"

const (
	root     = "ignite.gocloud"
	resource = ".resource"
	tp       = ".type"
	region   = ".region"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+tp, "memory", "define queue type")
	config.Add(path+resource, "topicA", "define queue resource")
	config.Add(path+region, "", "define queue region")
}
