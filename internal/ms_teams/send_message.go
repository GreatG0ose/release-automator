package ms_teams

import (
	"fmt"
	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/atc0005/go-teams-notify/v2/adaptivecard"
)

type Message struct {
	Title    string
	Content  string
	Mentions []Mention
}
type Mention struct {
	Name    string
	TeamsID string
}

func SendMessage(webhookUrl string, msg Message) error {
	// init client
	mstClient := goteamsnotify.NewTeamsClient()

	// setup message card
	card, err := adaptivecard.NewSimpleMessage(msg.Content, msg.Title, true)
	if err != nil {
		return fmt.Errorf("failed to create TextBlockCard: %w", err)
	}

	for _, m := range msg.Mentions {
		err = card.Mention(true, m.Name, m.TeamsID, "")
		if err != nil {
			return fmt.Errorf("failed to add mention for %s: %w", m.Name, err)
		}
	}

	// send
	err = mstClient.Send(webhookUrl, card)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}
