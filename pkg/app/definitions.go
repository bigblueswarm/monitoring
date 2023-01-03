// Package app provide the main application package
package app

import (
	"github.com/bigblueswarm/monitoring/pkg/auth"
	"github.com/bigblueswarm/monitoring/pkg/config"
	"github.com/bigblueswarm/monitoring/pkg/service"
	"github.com/gin-gonic/gin"
)

const authCookieName = "bbs_monitoring_session"

// Server is the monitoring web server struct
type Server struct {
	Router         *gin.Engine
	Config         *config.Config
	AuthProvider   auth.IProvider
	clusterService service.IClusterService
}

type loginForm struct {
	Password string `form:"pwd"`
}
