package utils

import (
	"crypto/rand"
	"math/big"

	"github.com/mjthecoder65/url-shortener/config"
)

func GenerateShortCode(config *config.Config) string {
	result := make([]byte, config.ShortCodeLength)
	charLength := big.NewInt(int64(len(config.AllowedChars)))

	for i := 0; i < config.ShortCodeLength; i++ {
		randomIndex, err := rand.Int(rand.Reader, charLength)

		if err != nil {
			return "", err
		}

		result[i] = config.AllowedChars[randomIndex.Int64()]
	}

	return string(result), nil
}
