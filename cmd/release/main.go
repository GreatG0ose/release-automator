package main

import (
	"flag"
	"fmt"
	"github.com/GreatG0ose/release-automator/internal/config"
	"github.com/GreatG0ose/release-automator/internal/fullrelease_message"
	"github.com/GreatG0ose/release-automator/internal/release"
	"github.com/GreatG0ose/release-automator/internal/release_notes"
	"github.com/GreatG0ose/release-automator/internal/signoff_message"
	"github.com/GreatG0ose/release-automator/internal/tweet"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

const (
	signOffCmd     = "signoff"
	sendMailCmd    = "mail"
	renderTweetCmd = "tweet"
	releaseNotes   = "notes"
)

func main() {
	// Setup logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Setup and parse arguments
	releaseVersion := flag.String(
		"version",
		"",
		"Version that will be released",
	)

	configPath := flag.String(
		"config",
		"release.yaml",
		"Path to release-automator YAML config",
	)
	flag.Parse()

	// Validate arguments
	if *releaseVersion == "" {
		log.Error().Msg("version is required")
		os.Exit(1)
	}

	if *configPath == "" {
		log.Error().Msg("config path cannot be empty")
		os.Exit(1)
	}

	// Load config
	log.Info().Str("config", *configPath).Msg("loading config")
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Err(err).Msg("couldn't load config")
		os.Exit(1)
	}

	// Adjust logger
	cmd := flag.Arg(0)
	l := log.With().
		Str("project", cfg.Project.Name).
		Str("version", *releaseVersion).
		Str("cmd", cmd).
		Logger()

	// Construct release context
	l.Info().Msg("creating release context")
	r, err := release.NewRelease(cfg, *releaseVersion)
	if err != nil {
		l.Error().Err(err).Msg("unable to create release context")
		os.Exit(1)
	}

	switch cmd {
	case signOffCmd:
		err = signoff_message.SendSignOffMessage(cfg, r)
	case sendMailCmd:
		err = fullrelease_message.Send(cfg, r)
	case renderTweetCmd:
		err = tweet.Generate(l, cfg, r)
	case releaseNotes:
		err = release_notes.Generate(l, cfg, r)
	default:
		err = fmt.Errorf("unknown command %s", cmd)
	}

	if err != nil {
		l.Err(err).Msg("execution failed")
		os.Exit(1)
	}

	l.Info().Msg("finished successfully")
}
