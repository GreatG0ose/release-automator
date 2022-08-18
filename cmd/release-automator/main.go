package main

import "flag"

func main() {
	// Parse and validate arguments
	releaseVersion := flag.String("release-version", "", "") // TODO: add usage

	// TODO: add confluence release page link
	// TODO: obtain go/nogo page from confluence page

	flag.Parse()

	if *releaseVersion == "" {
		panic("release-version is required")
	}

	mentions := []string{"Oleksii Ishchenko", "Rene Salecker", "Tess Akerlund", "Timo Seifert", "Simon Gabriel"}
	print(mentions)
}

type Release struct {
	ProjectName string // Name of the project
	Version     string // Version to release
	Changelog   ReleaseChangelog
}
