package language

import (
	"context"
	"golang.org/x/text/language"
)

func ToContext(ctx context.Context, lang string) (context.Context, error) {
	userLanguage, err := language.Parse(lang)
	if err != nil {
		return nil, err
	}
	return context.WithValue(ctx, UserKey(), userLanguage), nil
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
