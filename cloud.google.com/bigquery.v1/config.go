package bigquery

import (
	"github.com/americanas-go/config"
)

const (
	root            = "ignite.gcp.bigquery"
	projectID       = ".projectId"
	credentialsRoot = ".credentials"
	credentialsFile = credentialsRoot + ".file"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+projectID, "default", "defines project ID")
	config.Add(path+credentialsFile, "credentials.json", "sets credentials file")
}