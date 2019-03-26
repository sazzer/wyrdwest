package main

import (
	"database/sql"

	"github.com/gobuffalo/packr"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sazzer/wyrdwest/service/internal/database"
	"github.com/sirupsen/logrus"
)

func buildDatabase(dbURL string) database.DB {
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		logrus.WithError(err).Error("Failed to connect to database")
		panic(err)
	}
	db := database.NewFromDB(dbConn)
	logrus.Debug("Connected to database")

	migrations := &migrate.PackrMigrationSource{
		Box: packr.NewBox("../../migrations"),
	}
	n, err := migrate.Exec(dbConn, "postgres", migrations, migrate.Up)
	if err != nil {
		logrus.WithError(err).Error("Failed to migrate database")
		panic(err)
	}
	logrus.WithField("migrations", n).Info("Migrated database")

	return db
}
