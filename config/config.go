package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoDBURI      string `json:"database_url"`
	ServerPort      string `json:"server_port"`
	ShortCodeLength int    `json:"short_code_length"`
	AllowedChars    string `json:"allowed_chars"`
	AppEnv          string `json:"app_env"`
}

func LoadConfigs() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	shortCodeLength, err := strconv.Atoi(getEnv("SHORT_CODE_LENGTH"))

	if err != nil {
		log.Fatal(err)
	}
	return &Config{
		MongoDBURI:      getEnv("MONGODB_URI"),
		ServerPort:      getEnv("SERVER_PORT"),
		ShortCodeLength: shortCodeLength,
		AllowedChars:    getEnv("ALLOWED_CHARS"),
		AppEnv:          getEnv("APP_ENV"),
	}, nil
}

func getEnv(key string) string {
	value, exist := os.LookupEnv(key)

	if !exist {
		err := fmt.Errorf("failed to load config setting for key: %v", key)
		log.Fatal(err)
	}

	return value
}
