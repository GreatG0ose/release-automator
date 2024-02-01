package config

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name           string
		path           string
		expectedConfig Config
		expectedErr    error
	}{
		{
			name: "full config",
			path: "testdata/good.yaml",
			expectedConfig: Config{
				Project: Project{
					Name:          "<Project Name>",
					ChangelogPath: "CHANGELOG.md",
					Repository:    "",
				},
				SignOff: SignOff{
					TeamsWebhook: "https://webhook.office.com/webhookb2/<SOME-MAGIC-NUMBERS>",
					Mentions: []Mention{
						{
							Name:    "Firstname Lastname",
							TeamsID: "firstname.lastname@company.com",
						},
						{
							Name:    "Other Person",
							TeamsID: "other.person@company.com",
						},
					},
					Content: Content{
						Prepend: "Hello,\n\n",
						Append:  "\n\nBest regards,\nFirstname Lastname",
					},
				},
				FullReleaseEmail: FullReleaseEmail{
					OutlookWebhook: "https://webhook.office.com/webhookb2/<SOME-MAGIC-NUMBERS>",
					Template:       "Hello everyone,\n\nWe will release {{ .Config.Project.Name }} v{{ .Release.Version }} soon.\n\n{{ .Release.Changes.Summary }}\n\n<blockquote>\n<h1>Changelog</h1>\n\n{{ range $header, $body := .Release.Changes.Changes }}\n<h2>{{ $header }}</h2>\n\n{{ $body }}\n\n{{ end }}\n</blockquote>\n\n<a href=\"{{ .Release.ChecklistUrl }}\">Release checklist</a>\n\n<p>\nBest regards,\n</p>\nDevelopment Team\n",
				},
				Output: "target/",
				Confluence: Confluence{
					Endpoint:       "https://<COMPANY>.atlassian.net/wiki/rest/api",
					ReleasesPageId: "123",
					Credentials: Credentials{
						Username:    "your@company.com",
						AccessToken: "<access-token-obtained-from-confluence>",
					},
				},
			},
		},

		{
			name:        "no required template",
			path:        "testdata/no_template.yaml",
			expectedErr: errors.New("failed to validate config testdata/no_template.yaml: Key: 'Config.FullReleaseEmail.Template' Error:Field validation for 'Template' failed on the 'required' tag"),
		},

		{
			name:        "extra fields",
			path:        "testdata/extra_fields.yaml",
			expectedErr: errors.New("failed to unmarshall config testdata/extra_fields.yaml: yaml: unmarshal errors:\n  line 40: field content not found in type config.FullReleaseEmail"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config, err := LoadConfig(test.path)

			if test.expectedErr == nil {
				require.NoError(t, err)
				require.Equal(t, test.expectedConfig, config)
			} else {
				require.Error(t, err)
				require.EqualError(t, err, test.expectedErr.Error())
			}
		})
	}
}
