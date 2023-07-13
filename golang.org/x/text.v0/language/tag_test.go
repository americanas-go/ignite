package language

import (
	"testing"

	"github.com/americanas-go/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/text/language"
)

type TagSuite struct {
	suite.Suite
}

func TestTagSuite(t *testing.T) {
	suite.Run(t, new(TagSuite))
}

func (suite *TagSuite) SetupTest() {
	config.Load()
}

func (suite *TagSuite) TestDefaultTag() {
	testCases := []struct {
		Description    string
		ExpectedResult language.Tag
		ExpectedError  bool
	}{
		{
			Description:    "Get default language tag when it is set",
			ExpectedResult: language.MustParse("en-US"),
			ExpectedError:  false,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.Description, func() {

			result, err := DefaultTag()
			if tc.ExpectedError {
				assert.Error(suite.T(), err)
			} else {
				assert.NoError(suite.T(), err)
				assert.Equal(suite.T(), tc.ExpectedResult, result)
			}
		})
	}
}
