package database

import "database/sql"

// DB represents a wrapper around a database
type DB struct {
	db *sql.DB
}

// NewFromDB creates a new DB Wrapper from an already opened DB connection
func NewFromDB(db *sql.DB) DB {
	return DB{
		db: db,
	}
}
