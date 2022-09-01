package release

import "github.com/GreatG0ose/release-automator/internal/changelog"

type Release struct {
	Version      string                   // Version to release
	Changes      changelog.ReleaseChanges // Changes contains changes related to version release
	ChecklistUrl string                   // ChecklistUrl is a link to release checklist
}
