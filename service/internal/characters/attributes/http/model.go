package http

import (
	"fmt"

	uuid "github.com/satori/go.uuid"

	"github.com/sazzer/wyrdwest/service/internal/characters/attributes"
)

// Attribute represents a single attribute returned over the API
type Attribute struct {
	Self        string `json:"self"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// Attributes represents a page of attribute data returned over the API
type Attributes struct {
	Self   string      `json:"self"`
	First  string      `json:"first,omitempty"`
	Prev   string      `json:"prev,omitempty"`
	Next   string      `json:"next,omitempty"`
	Offset uint64      `json:"offset"`
	Total  uint64      `json:"total"`
	Data   []Attribute `json:"data"`
}

func buildAttribute(attribute attributes.Attribute) Attribute {
	return Attribute{
		Self:        fmt.Sprintf("/attributes/%s", uuid.UUID(attribute.ID).String()),
		Name:        attribute.Name,
		Description: attribute.Description,
	}
}
