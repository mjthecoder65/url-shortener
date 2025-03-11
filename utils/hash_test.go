package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateShortCode(t *testing.T) {
	shortCode := GenerateShortCode(testConfig)

	require.NotEmpty(t, shortCode)
	require.Equal(t, testConfig.ShortCodeLength, len(shortCode))
}
