package dao

import (
	"github.com/sazzer/wyrdwest/service/internal/database"
)

// AttributesDao is the standard implementation of the Attributes DAO
type AttributesDao struct {
	db database.DB
}

// New creates a new AttributeDao
func New(db database.DB) AttributesDao {
	return AttributesDao{db}
}
