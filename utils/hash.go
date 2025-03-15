package utils

import (
	"crypto/rand"
	"math/big"
	mathRand "math/rand"
	"time"

	"github.com/mjthecoder65/url-shortener/config"
)

var seedRand = mathRand.New(mathRand.NewSource(time.Now().UnixNano()))

func GenerateShortCodeFallBack(config *config.Config) string {

	result := make([]byte, config.ShortCodeLength)

	for i := 0; i < config.ShortCodeLength; i++ {
		result[i] = config.AllowedChars[seedRand.Intn(len(config.AllowedChars))]
	}

	return string(result)
}

func GenerateShortCode(config *config.Config) (string, error) {
	result := make([]byte, config.ShortCodeLength)
	charLength := big.NewInt(int64(len(config.AllowedChars)))

	for i := 0; i < config.ShortCodeLength; i++ {
		randomIndex, err := rand.Int(rand.Reader, charLength)
		if err != nil {
			return GenerateShortCodeFallBack(config), nil
		}
		result[i] = config.AllowedChars[randomIndex.Int64()]
	}

	if string(result) == "" {
		return GenerateShortCodeFallBack(config), nil
	}

	return string(result), nil
}
