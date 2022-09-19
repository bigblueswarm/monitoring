// Package auth provides the application authentication
package auth

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/b3lb/test_utils/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestNewProfider(t *testing.T) {
	t.Run("Generating an authentication provider should not throw an error", func(t *testing.T) {
		assert.NotNil(t, NewProvider(nil, "password"))
	})
}

func TestAuthenticate(t *testing.T) {
	provider := NewProvider(nil, "password")
	t.Run("Trying to authenticate with a bad password should return false", func(t *testing.T) {
		assert.False(t, provider.Authenticate("wrong_password"))
	})

	t.Run("Authenticate with the right password should return true", func(t *testing.T) {
		assert.True(t, provider.Authenticate("password"))
	})
}

func TestGenerateSession(t *testing.T) {
	provider := NewProvider(nil, "password")
	t.Run("Generating a new session should return return an error", func(t *testing.T) {
		assert.NotNil(t, provider.GenerateSession())
	})

	t.Run("Generating a new session should return a valid object", func(t *testing.T) {
		session := provider.GenerateSession()
		assert.NotNil(t, session.Timestamp)
		assert.NotNil(t, session.ID)
	})
}

func TestIsValidSession(t *testing.T) {
	provider := NewProvider(nil, "password")
	session := provider.GenerateSession()

	t.Run("A valid session should have a timestamp activity minus than 3 hours ago", func(t *testing.T) {
		timestamp := (session.Timestamp + (int64(time.Hour) * 2))
		assert.True(t, session.IsValid(timestamp))
	})

	t.Run("A new activity 3 hours later and more should invalid the session", func(t *testing.T) {
		timestamp := (session.Timestamp + (int64(time.Hour) * 4))
		assert.False(t, session.IsValid(timestamp))
	})
}

func TestSaveSession(t *testing.T) {
	provider := NewProvider(testRedisClient, "password")

	t.Run("Saving a session should not return an error", func(t *testing.T) {
		session := provider.GenerateSession()
		key := fmt.Sprintf("session:%s", session.ID)
		testRedisMock.ExpectSet(key, session.Timestamp, 0).SetVal(key)
		err := provider.SaveSession(session)
		assert.Nil(t, err)
	})

	t.Run("An error returned by redis while saving session should be returned", func(t *testing.T) {
		session := provider.GenerateSession()
		m := testRedisMock.ExpectSet(fmt.Sprintf("session:%s", session.ID), session.Timestamp, 0)
		m.SetVal("")
		m.SetErr(fmt.Errorf("redis err"))
		err := provider.SaveSession(session)
		assert.NotNil(t, err)
	})
}

func TestRetrieveSessio(t *testing.T) {
	provider := NewProvider(testRedisClient, "password")
	var id string

	tests := []test.Test{
		{
			Name: "An error returned by redis should be returned",
			Mock: func() {
				id = "session_id"
				m := testRedisMock.ExpectGet(fmt.Sprintf("session:%s", id))
				m.SetVal("")
				m.SetErr(errors.New("redis error"))
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.Nil(t, value)
				assert.NotNil(t, err)
			},
		},
		{
			Name: "A session not found in redis should return nil",
			Mock: func() {
				id = "fake_session"
				testRedisMock.ExpectGet(fmt.Sprintf("session:%s", id)).RedisNil()
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.Nil(t, value)
				assert.Nil(t, err)
			},
		},
		{
			Name: "An invalid timestamp should return an error",
			Mock: func() {
				id = "invalid_timestamp"
				testRedisMock.ExpectGet(fmt.Sprintf("session:%s", id)).SetVal("invalid")
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.Nil(t, value)
				assert.NotNil(t, err)
			},
		},
		{
			Name: "A valid id should return a valid session",
			Mock: func() {
				id = "valid_id"
				testRedisMock.ExpectGet(fmt.Sprintf("session:%s", id)).SetVal("123456")
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, value)
				session := value.(*Session)
				assert.Equal(t, session.ID, id)
				assert.Equal(t, session.Timestamp, int64(123456))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			test.Mock()
			s, err := provider.RetrieveSession(id)
			test.Validator(t, s, err)
		})
	}
}

func TestDropSessio(t *testing.T) {
	provider := NewProvider(testRedisClient, "password")
	var id string

	tests := []test.Test{
		{
			Name: "An error returned by redis should be returned",
			Mock: func() {
				id = "fake_id"
				mock := testRedisMock.ExpectDel(sessionKey(id))
				mock.SetErr(errors.New("redis error"))
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			Name: "No error should be returned when redis does not found session",
			Mock: func() {
				id = "unknown_id"
				testRedisMock.ExpectDel(sessionKey(id)).RedisNil()
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			test.Mock()
			test.Validator(t, nil, provider.DropSession(id))
		})
	}
}
