package main

import (
	"log"

	"github.com/mjthecoder65/url-shortener/api"
	"github.com/mjthecoder65/url-shortener/config"
)

func main() {
	configs, err := config.LoadConfigs()

	if err != nil {
		panic(err)
	}

	mongoClient := config.GetMongoDBClient(configs)

	server, err := api.NewServer(configs, mongoClient)

	if err != nil {
		log.Fatal(err)
	}

	err = server.Start()

	if err != nil {
		log.Fatal(err)
	}
}
