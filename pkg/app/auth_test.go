package app

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/b3lb/monitoring/pkg/auth"
	"github.com/b3lb/test_utils/pkg/request"
	"github.com/b3lb/test_utils/pkg/test"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthHandler(t *testing.T) {
	var w *httptest.ResponseRecorder
	var c *gin.Context
	server := &Server{
		AuthProvider: &auth.ProviderMock{},
	}
	tests := []test.Test{
		{
			Name: "An error returned by cookie should redirect to login page",
			Mock: func() {
				request.SetRequestParams(c, "")
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.Equal(t, http.StatusOK, w.Code) // Tests returns 200 instead of 302
				assert.Equal(t, "/auth/login", w.Header().Get("Location"))
			},
		},
		{
			Name: "An error returned by AuthProvider should redirect to login page",
			Mock: func() {
				request.SetRequestHeader(c, "cookie", fmt.Sprintf("%s=%s", authCookieName, "fake_cookie"))
				auth.RetrieveSessionProviderMockFunction = func(id string) (*auth.Session, error) {
					return nil, errors.New("provider error")
				}
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.Equal(t, http.StatusOK, w.Code) // Tests returns 200 instead of 302
				assert.Equal(t, "/auth/login", w.Header().Get("Location"))
			},
		},
		{
			Name: "An invalid session should return to login page",
			Mock: func() {
				auth.RetrieveSessionProviderMockFunction = func(id string) (*auth.Session, error) {
					return &auth.Session{
						ID:        "fake_cookie",
						Timestamp: time.Now().Unix() - (int64(time.Hour * 4)),
					}, nil
				}
				auth.DropSessionProviderMockFunction = func(id string) error { return nil }
				request.SetRequestHeader(c, "cookie", fmt.Sprintf("%s=%s", authCookieName, "fake_cookie"))
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.Equal(t, http.StatusOK, w.Code) // Tests returns 200 instead of 302
				assert.Equal(t, "/auth/login", w.Header().Get("Location"))
			},
		},
		{
			Name: "A valid authentication should call next handelr",
			Mock: func() {
				auth.RetrieveSessionProviderMockFunction = func(id string) (*auth.Session, error) {
					return &auth.Session{
						ID:        "fake_cookie",
						Timestamp: time.Now().Unix() - (int64(time.Hour * 1)),
					}, nil
				}
				auth.DropSessionProviderMockFunction = func(id string) error { return nil }
				auth.SaveSessionProviderMockFunction = func(s *auth.Session) error { return nil }
				request.SetRequestHeader(c, "cookie", fmt.Sprintf("%s=%s", authCookieName, "fake_cookie"))
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.False(t, c.IsAborted())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			w = httptest.NewRecorder()
			c, _ = gin.CreateTestContext(w)
			test.Mock()
			server.authHandler(c)
			test.Validator(t, nil, nil)
		})
	}
}

func TestLoginHandler(t *testing.T) {
	var w *httptest.ResponseRecorder
	var c *gin.Context
	server := &Server{
		AuthProvider: &auth.ProviderMock{},
	}
	tests := []test.Test{
		{
			Name: "A binding error should return a bad request error",
			Mock: func() {
				c.Request, _ = http.NewRequest("POST",
					"/auth/login", nil)
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.Equal(t, http.StatusBadRequest, w.Code)
			},
		},
		{
			Name: "An invalid password should redirect to login page with an invalid password message",
			Mock: func() {
				body := bytes.NewBufferString("pwd=fake_password")
				c.Request, _ = http.NewRequest("POST",
					"/auth/login", body)
				c.Request.Header.Add("Content-Type", gin.MIMEPOSTForm)
				auth.AuthenticateProviderMockFunction = func(pwd string) bool {
					return false
				}
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.Equal(t, http.StatusOK, w.Code)
				assert.Equal(t, "/auth/login?msg=invalid_pwd", w.Header().Get("Location"))
			},
		},
		{
			Name: "An error during session saving should return an internal server error",
			Mock: func() {
				body := bytes.NewBufferString("pwd=fake_password")
				c.Request, _ = http.NewRequest("POST",
					"/auth/login", body)
				c.Request.Header.Add("Content-Type", gin.MIMEPOSTForm)
				auth.AuthenticateProviderMockFunction = func(pwd string) bool {
					return true
				}
				auth.GenerateSessionProviderMockFunction = func() *auth.Session {
					return &auth.Session{
						ID:        "id",
						Timestamp: time.Now().Unix(),
					}
				}
				auth.SaveSessionProviderMockFunction = func(s *auth.Session) error {
					return errors.New("provider error")
				}
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.Equal(t, http.StatusInternalServerError, w.Code)
			},
		},
		{
			Name: "A valid authentication should redirect to / and set a cookie",
			Mock: func() {
				body := bytes.NewBufferString("pwd=fake_password")
				c.Request, _ = http.NewRequest("POST",
					"/auth/login", body)
				c.Request.Header.Add("Content-Type", gin.MIMEPOSTForm)
				auth.AuthenticateProviderMockFunction = func(pwd string) bool {
					return true
				}
				auth.GenerateSessionProviderMockFunction = func() *auth.Session {
					return &auth.Session{
						ID:        "id",
						Timestamp: time.Now().Unix(),
					}
				}
				auth.SaveSessionProviderMockFunction = func(s *auth.Session) error {
					return nil
				}
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.Equal(t, http.StatusOK, w.Code)
				assert.Equal(t, "/", w.Header().Get("Location"))
				assert.Equal(t, fmt.Sprintf("%s=id; Path=/; Max-Age=259200000000000", authCookieName), w.Header().Get("Set-Cookie"))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			w = httptest.NewRecorder()
			c, _ = gin.CreateTestContext(w)
			test.Mock()
			server.loginHandler(c)
			test.Validator(t, nil, nil)
		})
	}
}

func TestLogoutHandler(t *testing.T) {
	var w *httptest.ResponseRecorder
	var c *gin.Context
	server := &Server{
		AuthProvider: &auth.ProviderMock{},
	}
	tests := []test.Test{
		{
			Name: "No cookie found should redirect to login page",
			Mock: func() {
				request.SetRequestParams(c, "")
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.Equal(t, http.StatusOK, w.Code)
				assert.Equal(t, "/auth/login", w.Header().Get("Location"))
			},
		},
		{
			Name: "Log out should redirect to login page",
			Mock: func() {
				request.SetRequestHeader(c, "cookie", fmt.Sprintf("%s=%s", authCookieName, "fake_cookie"))
				auth.DropSessionProviderMockFunction = func(id string) error { return nil }
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.Equal(t, http.StatusOK, w.Code)
				assert.Equal(t, "/auth/login", w.Header().Get("Location"))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			w = httptest.NewRecorder()
			c, _ = gin.CreateTestContext(w)
			test.Mock()
			server.logoutHandler(c)
			test.Validator(t, nil, nil)
		})
	}
}
