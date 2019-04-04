package integration

import (
	"github.com/stretchr/testify/suite"

	"github.com/sazzer/wyrdwest/service/integration/database"
	"github.com/sazzer/wyrdwest/service/integration/http"
	"gopkg.in/h2non/baloo.v3"
)

// Suite is the base for all Integration Tests
type Suite struct {
	suite.Suite
}

// SetupTest will ensure that the Integration Suite is correctly defined
func (suite *Suite) SetupTest() {
	if !http.Enabled() || !database.Enabled() {
		suite.T().Skip()
	}
}

// StartTest will start the Baloo Test Client for our tests
func (suite *Suite) StartTest() *baloo.Client {
	return http.StartTest()
}

// ParseJSONToMap will convert the given JSON String to a Map for assertions
func (suite *Suite) ParseJSONToMap(input string) map[string]interface{} {
	return http.ParseJSONToMap(input)
}
