package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"
)

// Enabled checks if HTTP testing is enabled
func Enabled() bool {
	_, hasURL := os.LookupEnv("DB_URL")
	return hasURL
}

func connect() *sql.DB {
	dbConn, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		logrus.WithError(err).Error("Failed to connect to database")
		panic(err)
	}
	return dbConn
}
