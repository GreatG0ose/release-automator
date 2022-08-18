package main

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed testdata/changelog.md
var testChangelog string

func TestExtractReleaseNotes(t *testing.T) {
	expectedReleaseChangeLog := ReleaseChangelog{
		Summary: "Basket v2 is added and logger-core is removed.",
		Changes: "",
	}

	actualReleaseChangelog, err := ExtractReleaseNotes(testChangelog)

	assert.NoError(t, err)
	assert.Equal(t, expectedReleaseChangeLog, actualReleaseChangelog)
}
