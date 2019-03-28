package attributes

//go:generate mockgen -destination=./mocks/service_mock.go -package=mocks github.com/sazzer/wyrdwest/service/internal/characters/attributes Retriever

// Retriever defines a mechanism by which we can retrieve Attribute details
type Retriever interface {
	GetAttributeByID(id AttributeID) (Attribute, error)
}
