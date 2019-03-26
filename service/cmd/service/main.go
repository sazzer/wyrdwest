package main

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/paked/configure"
	"github.com/sazzer/wyrdwest/service/internal/database"
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

	healthchecker := health.New()

	dbConn, err := sql.Open("postgres", *dbURL)
	if err != nil {
		logrus.WithError(err).Error("Failed to connect to database")
		panic(err)
	}
	db := database.NewFromDB(dbConn)
	logrus.WithField("db", db).Info("Connected to database")
	healthchecker.AddHealthcheck("database", db)

	server := server.New()

	server.Register(health.RegisterHandler(&healthchecker))

	if err := server.Start(*port); err != nil {
		logrus.WithError(err).Error("Failed to start server")
	}
}
