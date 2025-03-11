package utils

import (
	"log"
	"os"
	"testing"

	"github.com/mjthecoder65/url-shortener/config"
)

var testConfig *config.Config

func TestMain(m *testing.M) {
	configs, err := config.LoadConfigs("../.env")

	if err != nil {
		log.Fatal("failed to load config: ", err)
	}

	testConfig = configs

	os.Exit(m.Run())
}
