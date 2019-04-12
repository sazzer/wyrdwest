package users

import (
	"time"
)

// UserID represents the ID of an User
type UserID string

// Authentication represents the authentication details of a single user
type Authentication struct {
	Provider    string
	ProviderID  string
	DisplayName string
}

// User represents a single User in the system
type User struct {
	ID              UserID
	Version         string
	Created         time.Time
	Updated         time.Time
	Name            string
	Email           string
	Authentications []Authentication
}
