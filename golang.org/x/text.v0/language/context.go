package language

import (
	"context"
	"golang.org/x/text/language"
)

func ToContext(ctx context.Context, lang string) context.Context {
	userLanguage := language.Make(lang)
	return context.WithValue(ctx, UserKey(), userLanguage)
}

func FromContext(ctx context.Context) (language.Tag, error) {
	var l string
	if ctxDefaultLang, ok := ctx.Value(UserKey()).(string); ok {
		l = ctxDefaultLang
	}
	if l != "" {
		return language.Parse(l)
	}
	return DefaultTag()
}
