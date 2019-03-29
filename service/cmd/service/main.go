package main

import (
	_ "github.com/lib/pq"
	"github.com/paked/configure"
	attributes "github.com/sazzer/wyrdwest/service/internal/characters/attributes/wiring"
	"github.com/sazzer/wyrdwest/service/internal/health"
	"github.com/sazzer/wyrdwest/service/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	conf := configure.New()
	port := conf.Int("port", 3000, "The port to listen on")
	dbURL := conf.String("db_url", "", "The database URL to connect to")

	conf.Use(configure.NewFlag())
	conf.Use(configure.NewEnvironment())

	conf.Parse()

	db := buildDatabase(*dbURL)

	server := server.New()

	_, attributesRegistration := attributes.AttributesWiring(db)
	server.Register(attributesRegistration)

	healthchecker := health.New()
	healthchecker.AddHealthcheck("database", db)
	server.AddRoutes(health.NewRouter(&healthchecker))

	if err := server.Start(*port); err != nil {
		logrus.WithError(err).Error("Failed to start server")
	}
}
