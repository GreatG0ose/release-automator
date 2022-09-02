package main

import (
	"github.com/GreatG0ose/release-automator/internal/changelog"
	"github.com/GreatG0ose/release-automator/internal/config"
	"github.com/GreatG0ose/release-automator/internal/release"
	"github.com/GreatG0ose/release-automator/internal/signoff_message"
	"github.com/rs/zerolog"
	"os"
)

// sendSignOffMessage creates and sends sing-off message to Teams channel
func sendSignOffMessage(l zerolog.Logger, cfg config.Config, releaseVersion *string, releaseChecklistUrl *string, releaseChanges changelog.ReleaseChanges) {
	// Send SignOff to MS Teams
	l.Info().Msg("sending message to teams")
	err := signoff_message.SendSignOffMessage(
		cfg,
		release.Release{
			Version:      *releaseVersion,
			ChecklistUrl: *releaseChecklistUrl,
			Changes:      releaseChanges,
		},
	)
	if err != nil {
		l.Error().Err(err).Msg("couldn't send sign-off message")
		os.Exit(1)
	}
}
