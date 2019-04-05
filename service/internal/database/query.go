package database

import (
	"github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"
)

// Query will execute a query against the database and return the resultset
func (db DB) Query(sql string, args interface{}) (*sqlx.Rows, error) {
	logrus.WithField("sql", sql).WithField("binds", args).Info("Executing query")
	return db.db.NamedQuery(sql, args)
}

// QueryPositional will execute a query produced by the Squirrel SQL Builder against the database and return the resultset
func (db DB) QueryPositional(sql string, args []interface{}) (*sqlx.Rows, error) {
	logrus.WithField("sql", sql).WithField("binds", args).Info("Executing query")
	return db.db.Queryx(sql, args...)
}
