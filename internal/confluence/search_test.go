package confluence

import (
	"github.com/GreatG0ose/release-automator/internal/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetPageLink(t *testing.T) {
	testCases := []struct {
		name        string
		cfg         config.Config
		version     string
		expected    string
		expectedErr string
	}{
		{
			name:    "happy path",
			version: "3.4",
			cfg: config.Config{
				Project: config.Project{
					Name: "ANTLR",
				},
				Confluence: config.Confluence{
					Endpoint:       "https://theantlrguy.atlassian.net/wiki/rest/api",
					ReleasesPageId: "2687234",
				},
			},
			expected: "https://theantlrguy.atlassian.net/wiki/pages/viewpage.action?pageId=2687376",
		},

		{
			name:    "wrong project name",
			version: "3.4",
			cfg: config.Config{
				Project: config.Project{
					Name: "ANTLRv4",
				},
				Confluence: config.Confluence{
					Endpoint:       "https://theantlrguy.atlassian.net/wiki/rest/api",
					ReleasesPageId: "2687234",
				},
			},
			expectedErr: "release page not found for project ANTLRv4 version 3.4",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actual, err := GetReleasePageLink(test.cfg, test.version)
			if test.expectedErr != "" {
				require.EqualError(t, err, test.expectedErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, test.expected, actual)
			}
		})
	}
}
