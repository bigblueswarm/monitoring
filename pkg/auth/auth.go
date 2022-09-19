// Package auth provides the application authentication
package auth

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/SLedunois/b3lb/v2/pkg/utils"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

// NewProvider initialize a new authentication provider
func NewProvider(redis *redis.Client, password string) IProvider {
	return &Provider{
		client:   redis,
		password: password,
	}
}

// Authenticate check if provided password equals provider password
func (p *Provider) Authenticate(pwd string) bool {
	return p.password == pwd
}

// GenerateSession generate a new authentication session
func (p *Provider) GenerateSession() *Session {
	return &Session{
		ID:        uuid.New().String(),
		Timestamp: time.Now().Unix(),
	}
}

// IsValid check if session is valid based on provided timestamp. To be valid,
// a session need to be less than 3 hours from the last activity
func (s *Session) IsValid(t int64) bool {
	diff := t - s.Timestamp
	return ((int64(time.Hour) * 3) - diff) > 0
}

func sessionKey(id string) string {
	return "session:" + id
}

// SaveSession save session in system
func (p *Provider) SaveSession(s *Session) error {
	_, err := p.client.Set(context.Background(), sessionKey(s.ID), s.Timestamp, 0).Result()
	return utils.ComputeErr(err)
}

// RetrieveSession retrieve session in system
func (p *Provider) RetrieveSession(id string) (*Session, error) {
	timestamp, err := p.client.Get(context.Background(), sessionKey(id)).Result()
	if utils.ComputeErr(err) != nil {
		return nil, err
	}

	if err == redis.Nil {
		return nil, nil
	}

	t, err := strconv.ParseInt(timestamp, 10, 64)

	if err != nil {
		return nil, fmt.Errorf("failed to parse timestamp %s: %s", timestamp, err)
	}

	return &Session{
		ID:        id,
		Timestamp: t,
	}, nil
}

// DropSession remove session from system
func (p *Provider) DropSession(id string) error {
	_, err := p.client.Del(context.Background(), sessionKey(id)).Result()
	return utils.ComputeErr(err)
}
