package main

import (
	"github.com/sazzer/wyrdwest/service/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	server := server.New()
	if err := server.Start(3000); err != nil {
		logrus.WithError(err).Error("Failed to start server")
	}
}
