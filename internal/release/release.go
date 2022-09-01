package release

import "github.com/GreatG0ose/release-automator/internal/changelog"

type Release struct {
	ProjectName string // Name of the project
	Version     string // Version to release
	Changelog   changelog.ReleaseChangelog
}
