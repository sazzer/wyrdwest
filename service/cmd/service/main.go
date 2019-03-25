package main

import (
	"errors"

	"github.com/paked/configure"
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
	conf := configure.New()
	port := conf.Int("port", 3000, "The port to listen on")

	conf.Use(configure.NewFlag())
	conf.Use(configure.NewEnvironment())

	conf.Parse()

	healthchecker := health.New()
	healthchecker.AddHealthcheck("dummy", dummyHealth{})

	server := server.New()

	server.Register(health.RegisterHandler(&healthchecker))

	if err := server.Start(*port); err != nil {
		logrus.WithError(err).Error("Failed to start server")
	}
}
