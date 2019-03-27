package database

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// DB represents a wrapper around a database
type DB struct {
	db *sqlx.DB
}

// NewFromDB creates a new DB Wrapper from an already opened DB connection
func NewFromDB(db *sql.DB) DB {
	return DB{
		db: sqlx.NewDb(db, "postgres"),
	}
}
