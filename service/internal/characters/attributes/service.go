package attributes

import (
	"github.com/sazzer/wyrdwest/service/internal/service"
)

//go:generate mockgen -destination=./mocks/service_mock.go -package=mocks github.com/sazzer/wyrdwest/service/internal/characters/attributes Retriever

// AttributeMatchCriteria defines the ways that we can match attributes when listing them
type AttributeMatchCriteria struct {
	Name string
}

// Retriever defines a mechanism by which we can retrieve Attribute details
type Retriever interface {
	// GetAttributeByID allows us to get a single attribute by it's unique ID
	GetAttributeByID(id AttributeID) (Attribute, error)

	// ListAttributes allows us to get a list of attributes that match certain criteria
	ListAttributes(criteria AttributeMatchCriteria, sorts []service.SortField, offset uint64, count uint64) (AttributePage, error)
}
