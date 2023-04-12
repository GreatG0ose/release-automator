package tweet

import (
	"github.com/GreatG0ose/release-automator/internal/changelog"
	"github.com/GreatG0ose/release-automator/internal/config"
	"github.com/GreatG0ose/release-automator/internal/release"
	"github.com/GreatG0ose/release-automator/internal/test_utils"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

func TestRender(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "")
	require.NoError(t, err)

	c := config.Config{
		Project: config.Project{
			Name:       "Java SDK",
			Repository: "https://github.com/unzerdev/java-sdk",
		},
		SignOff:          config.SignOff{},
		FullReleaseEmail: config.FullReleaseEmail{},
		Output:           tmpDir,
		Confluence:       config.Confluence{},
	}
	r := release.Release{
		Version: "3.0.0",
		Changes: changelog.ReleaseChanges{
			Summary: "This release brings liability and exemption type support to Java SDK",
		},
	}

	err = Generate(zerolog.Nop(), c, r)
	require.NoError(t, err)

	test_utils.FilesEqual(t, filepath.Join("testdata", "tweet.txt"), filepath.Join(tmpDir, "tweet.txt"))
}
