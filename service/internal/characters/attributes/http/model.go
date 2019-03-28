package http

// Attribute represents a single attribute returned over the API
type Attribute struct {
	Self        string `json:"self"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}
