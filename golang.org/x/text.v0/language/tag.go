package language

import (
	"golang.org/x/text/language"
)

var (
	t language.Tag
)

func DefaultTag() (language.Tag, error) {
	if t.String() == "" {
		tag, err := language.Parse(Default())
		if err != nil {
			return t, err
		}
		t = tag
	}
	return t, nil
}
