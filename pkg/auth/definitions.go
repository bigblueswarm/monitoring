// Package auth provides the application authentication
package auth

import "github.com/go-redis/redis/v8"

// IProvider describe Provider methods
type IProvider interface {
	// Authenticate check if provided password equals provider password
	Authenticate(pwd string) bool
	// GenerateSession generate a new authentication session
	GenerateSession() *Session
	// SaveSession save session in system
	SaveSession(s *Session) error
	// RetrieveSession retrieve session in system
	RetrieveSession(id string) (*Session, error)
	// DropSession remove session from system
	DropSession(id string) error
}

// Provider manage authentication application
type Provider struct {
	password string
	client   *redis.Client
}

// Session represents user session
type Session struct {
	ID        string
	Timestamp int64
}
