package dao

import "time"

type dbAttribute struct {
	ID          string    `db:"attribute_id"`
	Created     time.Time `db:"created"`
	Updated     time.Time `db:"updated"`
	Version     string    `db:"version"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
}
