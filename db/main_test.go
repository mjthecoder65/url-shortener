package db

import (
	"log"
	"os"
	"testing"

	"github.com/mjthecoder65/url-shortener/config"
)

var testQueries *Queries
var testConfig *config.Config

func TestMain(m *testing.M) {

	configs, err := config.LoadConfigs("../.env")
	if err != nil {
		log.Fatal(err)
	}
	mongoClient := config.GetMongoDBClient(configs)
	testConfig = configs

	testQueries = New(mongoClient)
	os.Exit(m.Run())
}
