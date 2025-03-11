package utils

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	ALLOWED_CHARS string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func TestGenerateShortCode(t *testing.T) {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	shortCodeLength := rand.Intn(2) + 6
	shortCode := GenerateShortCode(shortCodeLength, ALLOWED_CHARS)

	require.NotEmpty(t, shortCode)
	require.Equal(t, shortCodeLength, len(shortCode))
}
