// Package app provide the main application package
package app

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/b3lb/monitoring/pkg/auth"
	"github.com/b3lb/monitoring/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

//go:generate cp -r ../../dist ./dist
//go:embed dist
var dist embed.FS

func newRedisClient(address string, password string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
}

// NewServer initialize a new web server
func NewServer(config *config.Config) *Server {
	rc := newRedisClient(config.RDB.Address, config.RDB.Password, config.RDB.DB)

	return &Server{
		Router:       gin.Default(),
		Config:       config,
		AuthProvider: auth.NewProvider(rc, *config.Monitoring.Auth),
	}
}

// Run launch the Server web application
func (s *Server) Run() error {
	s.initRoutes()
	err := s.Router.Run(fmt.Sprintf(":%d", s.Config.Monitoring.Port))

	s.Router.Use(gin.Recovery())
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) initRoutes() {
	s.Router.StaticFS("/public", http.FS(dist))
	auth := s.Router.Group("/auth")
	{
		auth.GET("/login", loginPage)
		auth.POST("/login", s.loginHandler)
		auth.GET("/logout", s.logoutHandler)
	}
	s.Router.Use(s.authHandler)
	s.Router.GET("/", getIndex)
	s.Router.GET("/api", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello from api")
	})
}
