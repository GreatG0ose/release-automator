package main

import (
	"flag"
	"fmt"
	"github.com/GreatG0ose/release-automator/internal/changelog"
	"github.com/GreatG0ose/release-automator/internal/ms_teams"
	"github.com/GreatG0ose/release-automator/internal/release"
	"os"
)

func main() {
	// Parse and validate arguments
	releaseVersion := flag.String("release-version", "", "") // TODO: add usage
	projectName := flag.String("project-name", "", "")       // TODO: add usage
	changelogFile := flag.String("changelog-path", "", "")   // TODO: add usage

	// TODO: add confluence release page link
	// TODO: obtain go/nogo page from confluence page

	flag.Parse()

	if *releaseVersion == "" {
		panic("release-version is required")
	}

	if *projectName == "" {
		panic("project-name is required")
	}

	if *changelogFile == "" {
		panic("changelog-path is required")
	}

	// TODO: make mentions configurable
	mentions := []string{"Oleksii Ishchenko", "Rene Salecker", "Tess Akerlund", "Timo Seifert", "Simon Gabriel"}

	changelog, err := os.ReadFile(*changelogFile)
	if err != nil {
		panic(fmt.Errorf("failed to read changelog file '%s': %w", *changelogFile, err))
	}

	msg, err := generateSignoffMessage(*projectName, string(changelog), *releaseVersion, mentions)
	println(msg)
}

func generateSignoffMessage(projectName string, changelogText string, version string, mentions []string) (string, error) {
	releaseChangelog, err := changelog.ExtractReleaseNotes(changelogText, version)
	if err != nil {
		return "", fmt.Errorf("failed to extract release notes: %w", err)
	}

	message, err := ms_teams.RenderSignOffMessage(release.Release{
		ProjectName: projectName,
		Version:     version,
		Changelog:   releaseChangelog,
	}, mentions)
	if err != nil {
		return "", fmt.Errorf("failed to render message: %w", err)
	}

	return message, nil
}
