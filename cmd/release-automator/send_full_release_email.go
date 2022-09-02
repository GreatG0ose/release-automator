package main

import (
	"github.com/GreatG0ose/release-automator/internal/changelog"
	"github.com/GreatG0ose/release-automator/internal/config"
	"github.com/GreatG0ose/release-automator/internal/fullrelease_message"
	"github.com/GreatG0ose/release-automator/internal/release"
	"github.com/rs/zerolog"
	"os"
)

// sendFullReleaseEmail creates and sends full release email to Outlook group
func sendFullReleaseEmail(
	l zerolog.Logger,
	cfg config.Config,
	releaseVersion *string,
	releaseChecklistUrl *string,
	releaseChanges changelog.ReleaseChanges,
) {
	l.Info().Msg("sending message to teams")
	err := fullrelease_message.Send(
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
