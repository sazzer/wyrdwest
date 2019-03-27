package attributes

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// AttributeID represents the ID of an Attribute
type AttributeID uuid.UUID

// Attribute represents a single Attribute in the system
type Attribute struct {
	ID          AttributeID
	Version     uuid.UUID
	Created     time.Time
	Updated     time.Time
	Name        string
	Description string
}
