package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateShortCode(t *testing.T) {
	shortCode, err := GenerateShortCode(testConfig)
	require.NoError(t, err)
	require.NotEmpty(t, shortCode)
	require.Equal(t, testConfig.ShortCodeLength, len(shortCode))
}
