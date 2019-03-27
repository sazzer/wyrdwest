package dao

import (
	"github.com/sazzer/wyrdwest/service/internal/characters/attributes"
)

// AttributesDao represents the DAO layer for accessing Attributes
type AttributesDao interface {
	GetAttributeByID(id attributes.AttributeID) (attributes.Attribute, error)
}
