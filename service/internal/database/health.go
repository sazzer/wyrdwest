package database

import "github.com/sirupsen/logrus"

// CheckHealth checks that we can create a connection to the database and execute a trivial query
func (db DB) CheckHealth() error {
	_, err := db.db.Exec("SELECT 1")

	if err != nil {
		logrus.WithError(err).Error("Database Healthcheck Failed")
	} else {
		logrus.Debug("Database Healthcheck Passed")
	}

	return err
}
