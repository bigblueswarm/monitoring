// Package app provide the main application package
package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func redirectToLogin(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, "/auth/login")
}

func (s *Server) authHandler(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie(authCookieName)
	if err != nil {
		redirectToLogin(ctx)
		return
	}

	session, err := s.AuthProvider.RetrieveSession(cookie.Value)
	if err != nil {
		redirectToLogin(ctx)
		return
	}

	now := time.Now().Unix()
	if !session.IsValid(now) {
		s.AuthProvider.DropSession(session.ID)
		redirectToLogin(ctx)
		return
	}

	ctx.Next()
	session.Timestamp = now
	s.AuthProvider.SaveSession(session)
}

func loginPage(ctx *gin.Context) {
	renderPage("login", ctx)
}

func getHost(ctx *gin.Context) string {
	if host := ctx.Request.Header.Get("X-Forwarded-Host"); host != "" {
		return host
	}

	return ctx.Request.Host
}

func (s *Server) loginHandler(ctx *gin.Context) {
	form := &loginForm{}
	if err := ctx.Bind(form); err != nil {
		// Note: if there's a bind error, Gin will call
		// ctx.AbortWithError. We just need to return here.
		return
	}

	if !s.AuthProvider.Authenticate(form.Password) {
		ctx.Redirect(http.StatusFound, "/auth/login?msg=invalid_pwd")
		return
	}

	session := s.AuthProvider.GenerateSession()
	if err := s.AuthProvider.SaveSession(session); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to save current session %s: %s", session.ID, err))
		return
	}

	ctx.SetCookie(authCookieName, session.ID, int(time.Hour)*72, "/", getHost(ctx), false, false)

	ctx.Redirect(http.StatusFound, "/")
}

func (s *Server) logoutHandler(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie(authCookieName)
	if err != nil {
		redirectToLogin(ctx)
		return
	}

	ctx.SetCookie(authCookieName, cookie.Value, -1, "/", getHost(ctx), false, false)
	if err := s.AuthProvider.DropSession(cookie.Value); err != nil {
		log.Errorf("failed to drop session %s: %s", cookie.Value, err)
	}

	redirectToLogin(ctx)
}
