package signoff_message

import (
	_ "embed"
	"fmt"
	"github.com/GreatG0ose/release-automator/internal/config"
	"github.com/GreatG0ose/release-automator/internal/release"
	"github.com/atc0005/go-teams-notify/v2/adaptivecard"
)

func createSignOffMessage(cfg config.Config, release release.Release) (*adaptivecard.Message, error) {
	card := adaptivecard.NewCard()

	// Teams message consists of different blocks
	messageElements := []adaptivecard.Element{
		adaptivecard.NewTitleTextBlock(fmt.Sprintf("Request for management sign-off: %s Release Version v%s", cfg.Project.Name, release.Version), true),
		adaptivecard.NewTextBlock(release.Changes.Summary, true),                                               // Summary block
		adaptivecard.NewTextBlock("Release will take place provided the Go/NoGo meeting resulted in Go", true), // Go/NoGo part
		adaptivecard.NewTitleTextBlock("Changelog", true),                                                      // Changes title
	}

	// changes titles and blocks
	for h, b := range release.Changes.Changes {
		messageElements = append(messageElements,
			adaptivecard.NewTitleTextBlock(h, true),
			adaptivecard.NewTextBlock(b, true),
		)
	}

	// checklist block
	messageElements = append(messageElements,
		adaptivecard.NewTextBlock(
			fmt.Sprintf(
				"[Release checklist for %s v%s](%s)",
				cfg.Project.Name,
				release.Version,
				release.ChecklistUrl,
			),
			true,
		),
	)

	err := card.AddElement(false, messageElements...)
	if err != nil {
		return nil, fmt.Errorf("failed to add message blocks: %w", err)
	}

	// Mentions
	var cardMentions []adaptivecard.Mention
	for _, m := range cfg.SignOff.Mentions {
		mention, err := adaptivecard.NewMention(m.Name, m.TeamsID)
		if err != nil {
			return nil, fmt.Errorf("failed to add mention for %s: %w", m.Name, err)
		}
		cardMentions = append(cardMentions, mention)
	}
	err = card.AddMention(false, cardMentions...)
	if err != nil {
		return nil, fmt.Errorf("failed to add mentions: %w", err)
	}

	return adaptivecard.NewMessageFromCard(card)
}
