// Package app provide the main application package
package app

import (
	"github.com/b3lb/monitoring/pkg/config"
	"github.com/gin-gonic/gin"
)

// Server is the monitoring web server struct
type Server struct {
	Router *gin.Engine
	Config *config.Config
}
