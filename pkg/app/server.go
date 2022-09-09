package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Server represents the monitoring web server struct
type Server struct {
	Router *gin.Engine
}

// NewServer initialize a new web server
func NewServer() *Server {
	return &Server{
		Router: gin.Default(),
	}
}

// Run launch the Server web application
func (s *Server) Run() error {
	err := s.Router.Run(fmt.Sprintf(":8080"))

	if err != nil {
		return err
	}

	return nil
}
