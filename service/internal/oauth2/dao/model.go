package dao

import (
	"time"

	"github.com/lib/pq"
)

type dbClient struct {
	ID            string         `db:"client_id"`
	Created       time.Time      `db:"created"`
	Updated       time.Time      `db:"updated"`
	Version       string         `db:"version"`
	Name          string         `db:"name"`
	Owner         string         `db:"owner_id"`
	Secret        string         `db:"client_secret"`
	RedirectURIs  pq.StringArray `db:"redirect_uris"`
	ResponseTypes pq.StringArray `db:"response_types"`
	GrantTypes    pq.StringArray `db:"grant_types"`
}
