package dao

import (
	"github.com/sazzer/wyrdwest/service/internal/database"
)

// AttributesDaoImpl is the standard implementation of the Attributes DAO
type AttributesDaoImpl struct {
	db database.DB
}

// New creates a new AttributeDao
func New(db database.DB) AttributesDao {
	return AttributesDaoImpl{db}
}
