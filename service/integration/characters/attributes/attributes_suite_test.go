package attributes_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/sazzer/wyrdwest/service/integration"
)

type Suite struct {
	integration.Suite
}

func TestAttributesSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
