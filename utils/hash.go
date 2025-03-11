package utils

import (
	"math/rand"
	"time"

	"github.com/mjthecoder65/url-shortener/config"
)

func GenerateShortCode(config *config.Config) string {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	result := make([]byte, config.ShortCodeLength)

	for i := 0; i < config.ShortCodeLength; i++ {
		result[i] = config.AllowedChars[rand.Intn(len(config.AllowedChars))]
	}

	return string(result)
}
