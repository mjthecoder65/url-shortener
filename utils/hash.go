package utils

import (
	"math/rand"
	"time"
)

func GenerateShortCode(shortedCodeLength int, allowedChars string) string {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	result := make([]byte, shortedCodeLength)

	for i := 0; i < shortedCodeLength; i++ {
		result[i] = allowedChars[rand.Intn(len(allowedChars))]
	}

	return string(result)
}
