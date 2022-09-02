package main

import (
	"flag"
	"github.com/GreatG0ose/release-automator/internal/changelog"
	"github.com/GreatG0ose/release-automator/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

const configDefaultPath = "release-automator.yaml"

func main() {
	// Setup logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Setup and parse arguments
	releaseVersion := flag.String(
		"version",
		"",
		"Version that will be released",
	)

	releaseChecklistUrl := flag.String(
		"checklist-url",
		"",
		"Link to Checklist page created for the release",
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

	if *releaseChecklistUrl == "" {
		log.Error().Msg("checklist-url is required")
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
	l := log.With().Str("project", cfg.Project.Name).Str("version", *releaseVersion).Logger()

	// Extract changes
	l.Info().Msg("extracting changes from changelog")
	releaseChanges, err := changelog.ExtractReleaseChanges(cfg, *releaseVersion)
	if err != nil {
		l.Error().Err(err).Msg("unable to extract release changes")
	}

	switch flag.Arg(0) {
	case "signoff":
		sendSignOffMessage(l, cfg, releaseVersion, releaseChecklistUrl, releaseChanges)
	case "mail":
		sendFullReleaseEmail(l, cfg, releaseVersion, releaseChecklistUrl, releaseChanges)
	default:
		l.Error().Str("command", flag.Arg(0)).Msg("unknown command")
		os.Exit(1)
	}

	l.Info().Msg("execution completed")
}
