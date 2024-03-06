package pubsub

import (
	"github.com/americanas-go/config"
)

const (
	root            = "ignite.pubsub"
	projectID       = ".projectId"
	credentialsRoot = ".credentials"
	credentialsFile = credentialsRoot + ".file"
	credentialsJSON = credentialsRoot + ".json"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+projectID, "default", "defines project ID")
	config.Add(path+credentialsFile, "credentials.json", "sets credentials file")
	config.Add(path+credentialsJSON, "", "sets credentials json")
}
