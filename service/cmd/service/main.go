package main

import (
	"errors"

	"github.com/sazzer/wyrdwest/service/internal/health"
	"github.com/sazzer/wyrdwest/service/internal/server"
	"github.com/sirupsen/logrus"
)

type dummyHealth struct {
}

func (d dummyHealth) CheckHealth() error {
	return errors.New("Oops")
}

func main() {
	healthchecker := health.New()
	healthchecker.AddHealthcheck("dummy", dummyHealth{})

	server := server.New()

	server.Register(health.RegisterHandler(&healthchecker))

	if err := server.Start(3000); err != nil {
		logrus.WithError(err).Error("Failed to start server")
	}
}
