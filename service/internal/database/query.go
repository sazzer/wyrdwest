package database

import (
	"github.com/jmoiron/sqlx"
)

// Query will execute a query against the database and return the resultset
func (db DB) Query(sql string, args interface{}) (*sqlx.Rows, error) {
	return db.db.NamedQuery(sql, args)
}
