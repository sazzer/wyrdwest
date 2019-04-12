package http

import (
	"github.com/sazzer/wyrdwest/service/internal/api/uritemplate"

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
		Self:        uritemplate.BuildURI("/attributes{/id}", map[string]interface{}{"id": attribute.ID}),
		Name:        attribute.Name,
		Description: attribute.Description,
	}
}
