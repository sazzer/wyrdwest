package attributes

import (
	"time"

	"github.com/sazzer/wyrdwest/service/internal/service"
)

// AttributeID represents the ID of an Attribute
type AttributeID string

// Attribute represents a single Attribute in the system
type Attribute struct {
	ID          AttributeID
	Version     string
	Created     time.Time
	Updated     time.Time
	Name        string
	Description string
}

// AttributePage represents a page of Attribute records
type AttributePage struct {
	service.PageInfo
	Data []Attribute
}
