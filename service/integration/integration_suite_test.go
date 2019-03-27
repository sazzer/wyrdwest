package integration_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sazzer/wyrdwest/service/integration/http"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)

	if http.Enabled() {
		RunSpecs(t, "Integration Suite")
	}
}
