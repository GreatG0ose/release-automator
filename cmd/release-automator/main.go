package main

import (
	"flag"
	"fmt"
	"github.com/GreatG0ose/release-automator/internal/config"
	"github.com/GreatG0ose/release-automator/internal/fullrelease_message"
	"github.com/GreatG0ose/release-automator/internal/release"
	"github.com/GreatG0ose/release-automator/internal/signoff_message"
	"github.com/GreatG0ose/release-automator/internal/tweet"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

const configDefaultPath = "release-automator.yaml"
const (
	signOffCmd     = "signoff"
	sendMailCmd    = "mail"
	renderTweetCmd = "tweet"
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
		"config-path",
		"",
		"Path to release-automator YAML config",
	)
	flag.Parse()

	// Validate arguments
	if *releaseVersion == "" {
		log.Error().Msg("version is required")
		os.Exit(1)
	}

	if *configPath == "" {
		log.Warn().Str("config-path", configDefaultPath).Msg("config-path is not set. default value is used")
		*configPath = configDefaultPath
	} else {
		log.Info().Str("config-path", *configPath).Msg("custom config is used")
	}

	// Load config
	log.Info().Str("config-path", *configPath).Msg("loading config")
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.With().Str("config-path", configDefaultPath).Err(err).Str("msg", "couldn't load config")
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
		err = tweet.Render(l, cfg, r)
	default:
		err = fmt.Errorf("unknown command %s", cmd)
	}

	if err != nil {
		l.Err(err).Msg("execution failed")
		os.Exit(1)
	}

	l.Info().Msg("finished successfully")
}
