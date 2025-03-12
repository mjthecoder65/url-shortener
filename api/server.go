package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mjthecoder65/url-shortener/config"
	"github.com/mjthecoder65/url-shortener/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	config  *config.Config
	queries *db.Queries
	router  *gin.Engine
}

func NewServer(config *config.Config, mongoClient *mongo.Client) (*Server, error) {
	if config.AppEnv == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	server := &Server{
		config:  config,
		queries: db.New(mongoClient),
	}

	router := SetupRouter(server)

	server.router = router

	return server, nil
}

func SetupRouter(server *Server) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/shorten", server.CreateShortURL)
		v1.GET("/shorten/:shortCode", server.GetOrigionalURL)
		v1.PUT("/shorten/:shortCode", server.UpdateShortURL)
		v1.DELETE("/shorten/:shortCode", server.DeleteShortURL)
		v1.GET("/shorten/:shortCode/stats", server.GetURLStats)
	}

	return router
}

func (server *Server) Start() error {
	err := server.router.Run(server.config.ServerPort)
	if err == nil {
		log.Printf("Server listening on port %v\n", server.config.ServerPort)
	}
	return err
}
