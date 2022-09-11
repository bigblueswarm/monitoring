package app

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:generate cp -r ../../dist ./dist
//go:embed dist
var dist embed.FS

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
	s.initRoutes()
	err := s.Router.Run(fmt.Sprintf(":8080"))

	if err != nil {
		return err
	}

	return nil
}

func (s *Server) initRoutes() {
	s.Router.GET("/", func(ctx *gin.Context) {
		index, _ := dist.ReadFile("dist/index.html")
		ctx.Writer.Header().Add("Content-Type", "text/html")
		ctx.Writer.Write(index)
	})
	s.Router.StaticFS("/public", http.FS(dist))
	s.Router.GET("/api", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello from api")
	})
}
