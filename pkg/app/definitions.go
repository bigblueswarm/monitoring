// Package app provide the main application package
package app

import (
	"github.com/b3lb/monitoring/pkg/auth"
	"github.com/b3lb/monitoring/pkg/config"
	"github.com/gin-gonic/gin"
)

const authCookieName = "b3lb_monitoring_session"

// Server is the monitoring web server struct
type Server struct {
	Router       *gin.Engine
	Config       *config.Config
	AuthProvider auth.IProvider
}

type loginForm struct {
	Password string `form:"pwd"`
}
