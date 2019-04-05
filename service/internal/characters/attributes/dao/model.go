package dao

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sazzer/wyrdwest/service/internal/characters/attributes"
)

type dbAttribute struct {
	ID          string    `db:"attribute_id"`
	Created     time.Time `db:"created"`
	Updated     time.Time `db:"updated"`
	Version     string    `db:"version"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
}

// ToAPI converts a Database Model into an API Model
func (a dbAttribute) ToAPI() attributes.Attribute {
	return attributes.Attribute{
		ID:          attributes.AttributeID(uuid.FromStringOrNil(a.ID)),
		Version:     uuid.FromStringOrNil(a.Version),
		Created:     a.Created,
		Updated:     a.Updated,
		Name:        a.Name,
		Description: a.Description,
	}
}
