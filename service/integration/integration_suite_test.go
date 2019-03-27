package integration_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)

	_, hasURL := os.LookupEnv("SERVICE_URL")
	if hasURL {
		RunSpecs(t, "Integration Suite")
	}
}
