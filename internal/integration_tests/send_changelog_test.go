package integration_tests

import (
	"github.com/GreatG0ose/release-automator/internal/changelog"
	"github.com/GreatG0ose/release-automator/internal/config"
	"github.com/GreatG0ose/release-automator/internal/ms_teams"
	"github.com/GreatG0ose/release-automator/internal/release"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSendChangeLogToTeams(t *testing.T) {
	version := "1.2.0.0"

	cfg := config.Config{
		Project: config.Project{
			Name:          "Java SDK",
			ChangelogPath: "testdata/changelog.md",
		},
		SignOff: config.SignOff{
			TeamsWebhook: "https://heidelpay.webhook.office.com/webhookb2/b574ac44-cfa5-4cd2-8b23-7c3d52d0acea@797b2bda-888b-44ab-9967-3c9448c99377/IncomingWebhook/c1d68196df5e42f0a28edbfd85245843/f43437ec-7ce7-4ae5-8082-942c9829d10d",
			Mentions: []config.Mention{
				{
					Name:    "Vlad",
					TeamsID: "vladislav.lipianin@unzer.com",
				},
			},
		},
	}

	releaseChangelog, err := changelog.ExtractReleaseChanges(cfg, version)
	require.NoError(t, err)

	r := release.Release{
		Version: version,
		Changes: releaseChangelog,
	}

	err = ms_teams.SendSignOffMessage(cfg, r)
	require.NoError(t, err)
}
