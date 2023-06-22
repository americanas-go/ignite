package language

import (
	"context"
	"testing"

	"github.com/americanas-go/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/text/language"
)

type ContextSuite struct {
	suite.Suite
}

func TestContextSuite(t *testing.T) {
	suite.Run(t, new(ContextSuite))
}

func (suite *ContextSuite) SetupTest() {
	config.Load()
}

func (suite *ContextSuite) TestToContext() {
	testCases := []struct {
		Description   string
		Language      string
		ExpectedLang  string
		ExpectedError bool
	}{
		{
			Description:   "Set user language to context",
			Language:      "en-US",
			ExpectedLang:  "en-US",
			ExpectedError: false,
		},
		{
			Description:   "Set user language to context with invalid language",
			Language:      "invalid-language",
			ExpectedError: true,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.Description, func() {
			ctx := context.Background()
			ctx, err := ToContext(ctx, tc.Language)
			if tc.ExpectedError {
				assert.Error(suite.T(), err)
			} else {
				langTag, ok := ctx.Value(UserKey()).(language.Tag)
				assert.True(suite.T(), ok)
				assert.Equal(suite.T(), tc.ExpectedLang, langTag.String())
			}
		})
	}
}

func (suite *ContextSuite) TestFromContext() {
	testCases := []struct {
		Description    string
		Context        context.Context
		ExpectedResult language.Tag
		ExpectedError  bool
	}{
		{
			Description:    "Get user language from context",
			Context:        context.WithValue(context.Background(), UserKey(), "pt-BR"),
			ExpectedResult: language.MustParse("pt-BR"),
			ExpectedError:  false,
		},
		{
			Description:    "Get default language when user language is not set in context",
			Context:        context.Background(),
			ExpectedResult: language.MustParse("en-US"),
			ExpectedError:  false,
		},
		{
			Description:    "Get default language when user language is set to empty in context",
			Context:        context.WithValue(context.Background(), UserKey(), ""),
			ExpectedResult: language.MustParse("en-US"),
			ExpectedError:  false,
		},
		{
			Description:   "Get default language when user language is not valid",
			Context:       context.WithValue(context.Background(), UserKey(), "invalid-language"),
			ExpectedError: true,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.Description, func() {
			result, err := FromContext(tc.Context)
			if tc.ExpectedError {
				assert.Error(suite.T(), err)
			} else {
				assert.NoError(suite.T(), err)
				assert.Equal(suite.T(), tc.ExpectedResult, result)
			}
		})
	}
}
