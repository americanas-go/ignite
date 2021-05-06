package pubsub

import "github.com/americanas-go/config"

const (
	root     = "ignite.gocloud"
	resource = root + ".resource"
	tp       = root + ".type"
	region   = root + ".region"
)

func init() {
	config.Add(tp, "memory", "define queue type")
	config.Add(resource, "topicA", "define queue resource")
	config.Add(region, "", "define queue region")
}
