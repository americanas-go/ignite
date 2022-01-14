package ignite

import (
	"github.com/americanas-go/config"
)

// ignite options constraint.
type IgniteOptions interface {
	// returns options root path.
	Root() string
	// executes right after options are loaded.
	PostLoad() error
}

// returns ignite options from config file or environment vars.
func Load[O IgniteOptions]() (o O, e error) {
	o = New[O]()
	return LoadWithPath[O](o.Root())
}

// unmarshals ignite options based a given key path.
func LoadWithPath[O IgniteOptions](path string) (o O, e error) {
	o = New[O]()
	// adding default values
	tags := GetTags(o)
	for _, t := range tags {
		if t.Hide {
			config.Add(path+t.Path, t.Default, t.Description, config.WithHide())
		} else {
			config.Add(path+t.Path, t.Default, t.Description)
		}
	}
	// load config
	config.Load()
	if e = config.UnmarshalWithPath(path, o); e != nil {
		return
	}
	e = o.PostLoad()
	return
}
