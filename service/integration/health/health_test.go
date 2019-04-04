package health_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/sazzer/wyrdwest/service/integration"
)

type HealthSuite struct {
	integration.Suite
}

func (suite *HealthSuite) TestHealthchecks() {
	expected := `{
		"database": {
			"status": "ok"
		}
	}`

	suite.StartTest().Get("/health").
		Expect(suite.T()).
		Status(200).
		Type("json").
		JSON(suite.ParseJSONToMap(expected)).
		Done()
}

func TestHealthSuite(t *testing.T) {
	suite.Run(t, new(HealthSuite))
}
