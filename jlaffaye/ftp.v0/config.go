package ftp

import "github.com/americanas-go/config"

const (
	root    = "ignite.jlaffaye"
	addr    = ".addr"
	pu      = ".username"
	pp      = ".password"
	timeout = ".timeout"
	retry   = ".retry"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+addr, "", "ftp address")
	config.Add(path+pu, "", "ftp username")
	config.Add(path+pp, "", "ftp password", config.WithHide())
	config.Add(path+timeout, 10, "ftp timeout")
	config.Add(path+retry, 3, "ftp retry")
}
