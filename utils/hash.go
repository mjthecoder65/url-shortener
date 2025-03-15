package utils

import (
	"math/rand"
	"time"

	"github.com/mjthecoder65/url-shortener/config"
)

var seedRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateShortCode(config *config.Config) string {

	result := make([]byte, config.ShortCodeLength)

	for i := 0; i < config.ShortCodeLength; i++ {
		result[i] = config.AllowedChars[seedRand.Intn(len(config.AllowedChars))]
	}

	return string(result)
}
