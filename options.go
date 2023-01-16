package ignite

import "github.com/americanas-go/config"

// NewOptionsWithPath unmarshal options based a given multi key paths.
func NewOptionsWithPath[O any](paths ...string) (opts *O, err error) {
	opts = new(O)
	return unmarshall[O](opts, paths...)
}

func MergeOptionsWithPath[O any](opts *O, paths ...string) (*O, error) {
	return unmarshall[O](opts, paths...)
}

func unmarshall[O any](opts *O, paths ...string) (*O, error) {
	for _, path := range paths {
		err := config.UnmarshalWithPath(path, opts)
		if err != nil {
			return nil, err
		}
	}
	return opts, nil
}
