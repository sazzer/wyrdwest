package attributes

// Retriever defines a mechanism by which we can retrieve Attribute details
type Retriever interface {
	GetAttributeByID(id AttributeID) (Attribute, error)
}
