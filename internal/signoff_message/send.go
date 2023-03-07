package signoff_message

import (
	"fmt"
	"github.com/GreatG0ose/release-automator/internal/config"
	"github.com/GreatG0ose/release-automator/internal/release"
	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/atc0005/go-teams-notify/v2/adaptivecard"
)

// SendSignOffMessage sends release notes to MS Teams
func SendSignOffMessage(cfg config.Config, release release.Release) error {
	msg, err := createSignOffMessage(cfg, release)
	if err != nil {
		return fmt.Errorf("failed to create signoff message: %w", err)
	}

	err = sendMessage(cfg.SignOff.TeamsWebhook, msg)
	if err != nil {
		return fmt.Errorf("failed to send signoff message: %w", err)
	}

	return nil
}

// sendMessage sends passed message to MS Teams Webhook
func sendMessage(webhookUrl string, msg *adaptivecard.Message) error {
	// init client
	mstClient := goteamsnotify.NewTeamsClient()

	// send message
	err := mstClient.Send(webhookUrl, msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}
