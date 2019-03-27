package http

import (
	"os"

	"gopkg.in/h2non/baloo.v3"
)

// Enabled checks if HTTP testing is enabled
func Enabled() bool {
	_, hasURL := os.LookupEnv("SERVICE_URL")
	return hasURL
}

// StartTest returns the tester to use to execute tests against the HTTP service
func StartTest() *baloo.Client {
	url := os.Getenv("SERVICE_URL")
	test := baloo.New(url)
	return test
}
