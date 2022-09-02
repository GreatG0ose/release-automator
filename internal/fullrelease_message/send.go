package fullrelease_message

import (
	"fmt"
	"github.com/GreatG0ose/release-automator/internal/config"
	"github.com/GreatG0ose/release-automator/internal/release"
	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/atc0005/go-teams-notify/v2/messagecard"
)

// TODO: add docs
func Send(cfg config.Config, release release.Release) error {
	msg, err := createMessage(cfg, release)
	if err != nil {
		return fmt.Errorf("failed to create message: %w", err)
	}

	err = sendMessage(cfg.FullReleaseEmail.OutlookWebhook, msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

// sendMessage sends passed message to Outlook Webhook
func sendMessage(webhookUrl string, msg *messagecard.MessageCard) error {
	// init client
	mstClient := goteamsnotify.NewTeamsClient()

	// send message
	err := mstClient.Send(webhookUrl, msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}
