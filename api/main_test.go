package api

import (
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mjthecoder65/url-shortener/config"
)

var router *gin.Engine
var testConfig *config.Config

func TestMain(m *testing.M) {
	configs, err := config.LoadConfigs("../.env")

	if err != nil {
		log.Fatal(err)
	}
	testConfig = configs
	mongoClient := config.GetMongoDBClient(configs)

	server, err := NewServer(configs, mongoClient)

	if err != nil {
		log.Fatal(err)
	}

	router = server.router

	os.Exit(m.Run())
}
