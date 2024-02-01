package signoff_message

import (
	"fmt"
	"github.com/GreatG0ose/release-automator/internal/changelog"
	"github.com/GreatG0ose/release-automator/internal/config"
	"github.com/GreatG0ose/release-automator/internal/release"
	"github.com/atc0005/go-teams-notify/v2/adaptivecard"
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
		},
	}
	r := release.Release{
		Version: "1.0.0",
		Changes: changelog.ReleaseChanges{
			Summary: "Test Summary",
			Changes: map[string]string{
				"Changed": "Test Body",
				"Added":   "Test Body2",
			},
		},
	}

	expectedMessage := adaptivecard.Message{
		Type: "message",
		Attachments: []adaptivecard.Attachment{
			{
				ContentType: "application/vnd.microsoft.card.adaptive",
				ContentURL:  "",
				Content: adaptivecard.TopLevelCard{
					Card: adaptivecard.Card{
						Type:         "AdaptiveCard",
						Schema:       "http://adaptivecards.io/schemas/adaptive-card.json",
						Version:      "1.5",
						FallbackText: "",
						Body: []adaptivecard.Element{
							{
								Type:              "TextBlock",
								ID:                "",
								Text:              "Request for management sign-off: Test Project Release Version v1.0.0",
								URL:               "",
								Size:              "large",
								Weight:            "bolder",
								Color:             "",
								Spacing:           "",
								Style:             "heading",
								Items:             []adaptivecard.Element(nil),
								Columns:           []adaptivecard.Column(nil),
								Rows:              []adaptivecard.TableRow(nil),
								GridStyle:         "",
								FirstRowAsHeaders: (*bool)(nil),
								Visible:           (*bool)(nil),
								ShowGridLines:     (*bool)(nil),
								Actions:           []adaptivecard.Action(nil),
								SelectAction:      (*adaptivecard.ISelectAction)(nil),
								Facts:             []adaptivecard.Fact(nil),
								Wrap:              true,
								Separator:         false,
							},
							{
								Type:              "TextBlock",
								ID:                "",
								Text:              "Test Summary",
								URL:               "",
								Size:              "",
								Weight:            "",
								Color:             "",
								Spacing:           "",
								Style:             "",
								Items:             []adaptivecard.Element(nil),
								Columns:           []adaptivecard.Column(nil),
								Rows:              []adaptivecard.TableRow(nil),
								GridStyle:         "",
								FirstRowAsHeaders: (*bool)(nil),
								Visible:           (*bool)(nil),
								ShowGridLines:     (*bool)(nil),
								Actions:           []adaptivecard.Action(nil),
								SelectAction:      (*adaptivecard.ISelectAction)(nil),
								Facts:             []adaptivecard.Fact(nil),
								Wrap:              true,
								Separator:         false,
							},
							{
								Type:              "TextBlock",
								ID:                "",
								Text:              "Release will take place provided the Go/NoGo meeting resulted in Go",
								URL:               "",
								Size:              "",
								Weight:            "",
								Color:             "",
								Spacing:           "",
								Style:             "",
								Items:             []adaptivecard.Element(nil),
								Columns:           []adaptivecard.Column(nil),
								Rows:              []adaptivecard.TableRow(nil),
								GridStyle:         "",
								FirstRowAsHeaders: (*bool)(nil),
								Visible:           (*bool)(nil),
								ShowGridLines:     (*bool)(nil),
								Actions:           []adaptivecard.Action(nil),
								SelectAction:      (*adaptivecard.ISelectAction)(nil),
								Facts:             []adaptivecard.Fact(nil),
								Wrap:              true,
								Separator:         false,
							},
							{
								Type:              "TextBlock",
								ID:                "",
								Text:              "Changelog",
								URL:               "",
								Size:              "large",
								Weight:            "bolder",
								Color:             "",
								Spacing:           "",
								Style:             "heading",
								Items:             []adaptivecard.Element(nil),
								Columns:           []adaptivecard.Column(nil),
								Rows:              []adaptivecard.TableRow(nil),
								GridStyle:         "",
								FirstRowAsHeaders: (*bool)(nil),
								Visible:           (*bool)(nil),
								ShowGridLines:     (*bool)(nil),
								Actions:           []adaptivecard.Action(nil),
								SelectAction:      (*adaptivecard.ISelectAction)(nil),
								Facts:             []adaptivecard.Fact(nil),
								Wrap:              true,
								Separator:         false,
							},
							{
								Type:              "TextBlock",
								ID:                "",
								Text:              "Changed",
								URL:               "",
								Size:              "large",
								Weight:            "bolder",
								Color:             "",
								Spacing:           "",
								Style:             "heading",
								Items:             []adaptivecard.Element(nil),
								Columns:           []adaptivecard.Column(nil),
								Rows:              []adaptivecard.TableRow(nil),
								GridStyle:         "",
								FirstRowAsHeaders: (*bool)(nil),
								Visible:           (*bool)(nil),
								ShowGridLines:     (*bool)(nil),
								Actions:           []adaptivecard.Action(nil),
								SelectAction:      (*adaptivecard.ISelectAction)(nil),
								Facts:             []adaptivecard.Fact(nil),
								Wrap:              true,
								Separator:         false,
							},
							{
								Type:              "TextBlock",
								ID:                "",
								Text:              "Test Body",
								URL:               "",
								Size:              "",
								Weight:            "",
								Color:             "",
								Spacing:           "",
								Style:             "",
								Items:             []adaptivecard.Element(nil),
								Columns:           []adaptivecard.Column(nil),
								Rows:              []adaptivecard.TableRow(nil),
								GridStyle:         "",
								FirstRowAsHeaders: (*bool)(nil),
								Visible:           (*bool)(nil),
								ShowGridLines:     (*bool)(nil),
								Actions:           []adaptivecard.Action(nil),
								SelectAction:      (*adaptivecard.ISelectAction)(nil),
								Facts:             []adaptivecard.Fact(nil),
								Wrap:              true,
								Separator:         false,
							},
							{
								Type:              "TextBlock",
								ID:                "",
								Text:              "Added",
								URL:               "",
								Size:              "large",
								Weight:            "bolder",
								Color:             "",
								Spacing:           "",
								Style:             "heading",
								Items:             []adaptivecard.Element(nil),
								Columns:           []adaptivecard.Column(nil),
								Rows:              []adaptivecard.TableRow(nil),
								GridStyle:         "",
								FirstRowAsHeaders: (*bool)(nil),
								Visible:           (*bool)(nil),
								ShowGridLines:     (*bool)(nil),
								Actions:           []adaptivecard.Action(nil),
								SelectAction:      (*adaptivecard.ISelectAction)(nil),
								Facts:             []adaptivecard.Fact(nil),
								Wrap:              true,
								Separator:         false,
							},
							{
								Type:              "TextBlock",
								ID:                "",
								Text:              "Test Body2",
								URL:               "",
								Size:              "",
								Weight:            "",
								Color:             "",
								Spacing:           "",
								Style:             "",
								Items:             []adaptivecard.Element(nil),
								Columns:           []adaptivecard.Column(nil),
								Rows:              []adaptivecard.TableRow(nil),
								GridStyle:         "",
								FirstRowAsHeaders: (*bool)(nil),
								Visible:           (*bool)(nil),
								ShowGridLines:     (*bool)(nil),
								Actions:           []adaptivecard.Action(nil),
								SelectAction:      (*adaptivecard.ISelectAction)(nil),
								Facts:             []adaptivecard.Fact(nil),
								Wrap:              true,
								Separator:         false,
							},
							{
								Type:              "TextBlock",
								ID:                "",
								Text:              "[Release checklist for Test Project v1.0.0]()",
								URL:               "",
								Size:              "",
								Weight:            "",
								Color:             "",
								Spacing:           "",
								Style:             "",
								Items:             []adaptivecard.Element(nil),
								Columns:           []adaptivecard.Column(nil),
								Rows:              []adaptivecard.TableRow(nil),
								GridStyle:         "",
								FirstRowAsHeaders: (*bool)(nil),
								Visible:           (*bool)(nil),
								ShowGridLines:     (*bool)(nil),
								Actions:           []adaptivecard.Action(nil),
								SelectAction:      (*adaptivecard.ISelectAction)(nil),
								Facts:             []adaptivecard.Fact(nil),
								Wrap:              true,
								Separator:         false,
							},
							{
								Type:              "TextBlock",
								ID:                "",
								Text:              "<at>Test User</at> <at>Test User 2</at> ",
								URL:               "",
								Size:              "",
								Weight:            "",
								Color:             "",
								Spacing:           "",
								Style:             "",
								Items:             []adaptivecard.Element(nil),
								Columns:           []adaptivecard.Column(nil),
								Rows:              []adaptivecard.TableRow(nil),
								GridStyle:         "",
								FirstRowAsHeaders: (*bool)(nil),
								Visible:           (*bool)(nil),
								ShowGridLines:     (*bool)(nil),
								Actions:           []adaptivecard.Action(nil),
								SelectAction:      (*adaptivecard.ISelectAction)(nil),
								Facts:             []adaptivecard.Fact(nil),
								Wrap:              true,
								Separator:         false,
							}},
						Actions: []adaptivecard.Action(nil),
						MSTeams: adaptivecard.MSTeams{
							Entities: []adaptivecard.Mention{
								{
									Type:      "mention",
									Text:      "<at>Test User</at>",
									Mentioned: adaptivecard.Mentioned{ID: "testuser@email.com", Name: "Test User"},
								},
								{
									Type:      "mention",
									Text:      "<at>Test User 2</at>",
									Mentioned: adaptivecard.Mentioned{ID: "testuser2@email.com", Name: "Test User 2"},
								}},
						},
					}},
			}},
		AttachmentLayout: "",
	}

	msg, err := createSignOffMessage(cfg, r)

	fmt.Printf("%#v\n", *msg)
	require.NoError(t, err)
	require.Equal(t, expectedMessage, *msg)
}