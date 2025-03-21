package api

import (
	"log"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/mjthecoder65/url-shortener/config"
	"github.com/mjthecoder65/url-shortener/db"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	config  *config.Config
	queries *db.Queries
	router  *gin.Engine
	logger  *logrus.Logger
}

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetLevel(logrus.DebugLevel)
	return logger
}

func NewServer(config *config.Config, mongoClient *mongo.Client) (*Server, error) {
	if config.AppEnv == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	server := &Server{
		config:  config,
		queries: db.New(mongoClient),
		logger:  NewLogger(),
	}

	router := SetupRouter(server)

	server.router = router

	return server, nil
}

func SetupRouter(server *Server) *gin.Engine {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("httpurl", urlValidator)
	}

	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", server.HealthCheck)
		v1.POST("/shorten", server.CreateShortURL)
		v1.GET("/shorten/:shortCode", server.GetOriginalURL)
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

func urlValidator(fl validator.FieldLevel) bool {
	u, err := url.ParseRequestURI(fl.Field().String())
	return err == nil && (u.Scheme == "http" || u.Scheme == "https")
}
