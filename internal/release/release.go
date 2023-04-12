package release

import (
	"fmt"
	"github.com/GreatG0ose/release-automator/internal/changelog"
	"github.com/GreatG0ose/release-automator/internal/config"
	"github.com/GreatG0ose/release-automator/internal/confluence"
)

// Release contains target version, changes and url to release page
type Release struct {
	Version      string                   // Version to release
	Changes      changelog.ReleaseChanges // Changes contains changes related to version release
	ChecklistUrl string                   // ChecklistUrl is a link to release checklist
}

// NewRelease creates Release object
func NewRelease(cfg config.Config, version string) (Release, error) {
	checklistUrl, err := confluence.GetReleasePageLink(cfg, version)
	if err != nil {
		return Release{}, err
	}

	releaseChanges, err := changelog.ExtractReleaseChanges(cfg, version)
	if err != nil {
		return Release{}, fmt.Errorf("unable to extract release changes: %w", err)
	}

	return Release{
		Version:      version,
		Changes:      releaseChanges,
		ChecklistUrl: checklistUrl,
	}, nil
}
