package signoff_message

import (
	"github.com/GreatG0ose/release-automator/internal/changelog"
	"github.com/GreatG0ose/release-automator/internal/config"
	"github.com/GreatG0ose/release-automator/internal/release"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateSignoffMessage(t *testing.T) {
	cfg := config.Config{
		Project: config.Project{
			Name: "Test Project",
		},
		SignOff: config.SignOff{
			Mentions: []config.Mention{
				{
					Name:    "Test User",
					TeamsID: "testuser@email.com",
				},
				{
					Name:    "Test User 2",
					TeamsID: "testuser2@email.com",
				},
			},
			Content: config.Content{
				Prepend: "Release will take place provided the Go/NoGo meeting resulted in Go",
				Append:  "Please review the changelog and checklist and provide your sign-off by replying to this message with either :thumbsup: or :thumbsdown:",
			},
		},
	}
	r := release.Release{
		Version: "1.0.0",
		Changes: changelog.ReleaseChanges{
			Summary: "Test Summary",
			Changes: []changelog.ChangeBlock{
				{"Changed", "Test Body"},
				{"Added", "Test Body2"},
			},
		},
	}

	expectedMessage := `{
	"type": "message",
	"attachments": [
		{
			"contentType": "application/vnd.microsoft.card.adaptive",
			"content": {
				"type": "AdaptiveCard",
				"$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
				"version": "1.5",
				"body": [
					{
						"type": "TextBlock",
						"text": "Request for management sign-off: Test Project Release Version v1.0.0",
						"size": "large",
						"weight": "bolder",
						"style": "heading",
						"wrap": true
					},
					{
						"type": "TextBlock",
						"text": "Release will take place provided the Go/NoGo meeting resulted in Go",
						"wrap": true
					},
					{
						"type": "TextBlock",
						"text": "Changelog",
						"size": "large",
						"weight": "bolder",
						"style": "heading",
						"wrap": true
					},
					{
						"type": "TextBlock",
						"text": "Test Summary",
						"wrap": true
					},
					{
						"type": "TextBlock",
						"text": "Changed",
						"size": "large",
						"weight": "bolder",
						"style": "heading",
						"wrap": true
					},
					{
						"type": "TextBlock",
						"text": "Test Body",
						"wrap": true
					},
					{
						"type": "TextBlock",
						"text": "Added",
						"size": "large",
						"weight": "bolder",
						"style": "heading",
						"wrap": true
					},
					{
						"type": "TextBlock",
						"text": "Test Body2",
						"wrap": true
					},
					{
						"type": "TextBlock",
						"text": "[Release checklist for Test Project v1.0.0]()",
						"wrap": true
					},
					{
						"type": "TextBlock",
						"text": "Please review the changelog and checklist and provide your sign-off by replying to this message with either :thumbsup: or :thumbsdown:",
						"wrap": true
					},
					{
						"type": "TextBlock",
						"text": "\u003cat\u003eTest User\u003c/at\u003e \u003cat\u003eTest User 2\u003c/at\u003e ",
						"wrap": true
					}
				],
				"msteams": {
					"entities": [
						{
							"type": "mention",
							"text": "\u003cat\u003eTest User\u003c/at\u003e",
							"mentioned": {
								"id": "testuser@email.com",
								"name": "Test User"
							}
						},
						{
							"type": "mention",
							"text": "\u003cat\u003eTest User 2\u003c/at\u003e",
							"mentioned": {
								"id": "testuser2@email.com",
								"name": "Test User 2"
							}
						}
					]
				}
			}
		}
	]
}`

	msg, err := createSignOffMessage(cfg, r)

	require.NoError(t, msg.Prepare())
	require.NoError(t, err)
	require.Equal(t, expectedMessage, msg.PrettyPrint())
}
