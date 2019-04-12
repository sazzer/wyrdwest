package oauth2

import (
	"time"

	"github.com/sazzer/wyrdwest/service/internal/users"
)

// ResponseType represents the type of response that an OAuth2 Client can support
type ResponseType int

const (
	// ResponseTypeCode represents the OAuth2 Response Type of "code"
	ResponseTypeCode ResponseType = iota
	// ResponseTypeToken represents the OAuth2 Response Type of "token"
	ResponseTypeToken
	// ResponseTypeIDToken represents the OAuth2 Response Type of "id_token"
	ResponseTypeIDToken
)

// GrantType represents the type of grant that an OAuth2 Client can support
type GrantType int

const (
	// GrantTypeAuthorizationCode represents the OAuth2 Grant Type of "authorization_code"
	GrantTypeAuthorizationCode GrantType = iota
	// GrantTypeImplicit represents the OAuth2 Grant Type of "implicit"
	GrantTypeImplicit
	// GrantTypeRefreshToken represents the OAuth2 Grant Type of "refresh_token"
	GrantTypeRefreshToken
	// GrantTypeClientCredentials represents the OAuth2 Grant Type of "client_credentials"
	GrantTypeClientCredentials
)

// ClientID represents the ID of an Client
type ClientID string

// Client represents a single OAuth2 Client in the system
type Client struct {
	ID            ClientID
	Version       string
	Created       time.Time
	Updated       time.Time
	Name          string
	Secret        string
	OwnerID       users.UserID
	RedirectURIs  []string
	ResponseTypes []ResponseType
	GrantTypes    []GrantType
}
