package oauth2

// ClientNotFoundError represents when a requested OAuth2 Client could not be found
type ClientNotFoundError struct{}

// Error returns the error message for the ClientNotFoundError
func (e ClientNotFoundError) Error() string {
	return "The requested Client could not be found"
}
