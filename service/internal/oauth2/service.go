package oauth2

// Retriever defines a way to load client details
type Retriever interface {
	// GetClientByID allows us to load a Client knowing only it's ID
	GetClientByID(id ClientID) (Client, error)

	// GetClientByID allows us to load a Client knowing it's ID and Secret
	GetClientByIDAndSecret(id ClientID, secret string) (Client, error)
}
