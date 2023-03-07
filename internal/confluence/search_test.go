package confluence

import (
	"github.com/GreatG0ose/release-automator/internal/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetPageLink(t *testing.T) {
	expectedLink := "https://theantlrguy.atlassian.net/wiki/pages/viewpage.action?pageId=2687376"
	cfg := config.Config{
		Confluence: config.Confluence{
			Endpoint:       "https://theantlrguy.atlassian.net/wiki/rest/api",
			ReleasesPageId: "2687234",
		},
	}

	actualLink, err := GetReleasePageLink(cfg, "3.4")

	require.NoError(t, err)
	require.Equal(t, expectedLink, actualLink)
}
