package test_utils

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func FilesEqual(t *testing.T, expected, actual string) {
	actualFile, err := os.ReadFile(actual)
	assert.NoError(t, err)

	expectedFile, err := os.ReadFile(expected)
	assert.NoError(t, err)

	require.Equal(t, string(expectedFile), string(actualFile))
}
