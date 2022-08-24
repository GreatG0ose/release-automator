package main

import (
	"flag"
	"fmt"
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

func generateSignoffMessage(projectName string, changelog string, version string, mentions []string) (string, error) {
	releaseChangelog, err := ExtractReleaseNotes(changelog, version)
	if err != nil {
		return "", fmt.Errorf("failed to extract release notes: %w", err)
	}

	release := Release{
		ProjectName: projectName,
		Version:     version,
		Changelog:   releaseChangelog,
	}

	message, err := renderSignOffMessage(release, mentions)
	if err != nil {
		return "", fmt.Errorf("failed to render message: %w", err)
	}

	return message, nil
}

type Release struct {
	ProjectName string // Name of the project
	Version     string // Version to release
	Changelog   ReleaseChangelog
}
