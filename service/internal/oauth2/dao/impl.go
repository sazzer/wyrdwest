package dao

import (
	"github.com/sazzer/wyrdwest/service/internal/database"
)

// OAuth2ClientsDao is the standard implementation of the OAuth2 Clients DAO
type OAuth2ClientsDao struct {
	db database.DB
}

// New creates a new OAuth2ClientsDao
func New(db database.DB) OAuth2ClientsDao {
	return OAuth2ClientsDao{db}
}
