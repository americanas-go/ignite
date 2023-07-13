package language

import (
	"testing"

	"github.com/americanas-go/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConfigSuite struct {
	suite.Suite
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
}

func (suite *ConfigSuite) SetupTest() {
	config.Load()
}

func (suite *ConfigSuite) TestDefault() {
	testCases := []struct {
		Description string
		Expected    string
	}{
		{
			Description: "Get default language when it is set",
			Expected:    "en-US",
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.Description, func() {
			result := Default()
			assert.Equal(suite.T(), tc.Expected, result)
		})
	}
}

func (suite *ConfigSuite) TestUserKey() {
	testCases := []struct {
		Description   string
		ExpectedValue string
	}{
		{
			Description:   "Get user context key when it is set",
			ExpectedValue: "userLang",
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.Description, func() {
			result := UserKey()
			assert.Equal(suite.T(), tc.ExpectedValue, result)
		})
	}
}
