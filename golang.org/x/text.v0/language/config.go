package language

import (
	"github.com/americanas-go/config"
)

var (
	uKey string
	d    string
)

const (
	root   = "ignite.language"
	def    = root + ".default"
	usrKey = root + ".userKey"
)

func init() {
	config.Add(def, "en-US", "default lang")
	config.Add(usrKey, "userLang", "user context key")
}

func Default() string {
	if d == "" {
		d = config.String(def)
	}
	return d
}

func UserKey() string {
	if uKey == "" {
		uKey = config.String(usrKey)
	}
	return uKey
}
