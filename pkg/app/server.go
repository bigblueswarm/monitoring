// Package app provide the main application package
package app

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/b3lb/monitoring/pkg/auth"
	"github.com/b3lb/monitoring/pkg/config"
	"github.com/b3lb/monitoring/pkg/graphql"
	"github.com/b3lb/monitoring/pkg/graphql/generated"
	"github.com/b3lb/monitoring/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
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

func newInfluxDBClient(address string, token string) influxdb2.Client {
	return influxdb2.NewClient(address, token)
}

// Defining the Graphql handler
func (s *Server) graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	resolvers := &graphql.Resolver{
		ClusterService: s.clusterService,
	}
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// NewServer initialize a new web server
func NewServer(config *config.Config) *Server {
	rc := newRedisClient(config.RDB.Address, config.RDB.Password, config.RDB.DB)
	ic := newInfluxDBClient(config.IDB.Address, config.IDB.Token)

	return &Server{
		Router:         gin.Default(),
		Config:         config,
		AuthProvider:   auth.NewProvider(rc, *config.Monitoring.Auth),
		clusterService: service.NewClusterService(ic, config.IDB.Organization, config.IDB.Bucket, *config.Monitoring.AggregationInterval),
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
	s.Router.POST("/query", s.graphqlHandler())
	s.Router.GET("/graphiql", playgroundHandler())
	s.Router.GET("/", getIndex)
	s.Router.GET("/api", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello from api")
	})
}
