package attributes

// AttributeNotFoundError represents when a requested Attribute could not be found
type AttributeNotFoundError struct{}

// Error returns the error message for the AttributeNotFoundError
func (e AttributeNotFoundError) Error() string {
	return "The requested Attribute could not be found"
}
