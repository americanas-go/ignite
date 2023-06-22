package language

import (
	"golang.org/x/text/language"
)

var (
	t *language.Tag
)

func DefaultTag() (language.Tag, error) {
	if t == nil {
		tag, err := language.Parse(Default())
		if err != nil {
			return *t, err
		}
		t = &tag
	}
	return *t, nil
}
