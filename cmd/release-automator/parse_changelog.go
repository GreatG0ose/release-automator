package main

import (
	"fmt"
	"strings"
)

// ReleaseChangelog contains summary of a release and it's changes
type ReleaseChangelog struct {
	Summary string // Overall summary of the release
	Changes string // Markdown formatted changes
}

// ExtractReleaseNotes parses changelog and returns ReleaseChangelog for version
func ExtractReleaseNotes(changelog string, releaseVersion string) (ReleaseChangelog, error) {
	lines := strings.Split(
		normalizeString(changelog),
		"\n",
	)

	// Find segment start
	releaseStart := -1
	for i, l := range lines {
		if strings.HasPrefix(l, "## ") {
			if strings.Contains(l, releaseVersion) {
				releaseStart = i
				break
			}
		}
	}
	if releaseStart == -1 {
		return ReleaseChangelog{}, fmt.Errorf("release version '%s' not found in changelog", releaseVersion)
	}

	// Find segment end
	releaseEnd := len(lines) - 1
	for i := releaseStart + 1; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], "## ") {
			releaseEnd = i
			break
		}
	}

	releaseSegment := lines[releaseStart:releaseEnd]

	// Extract summary
	summaryEnd := -1
	for i, l := range releaseSegment[1:] {
		if strings.HasPrefix(l, "###") {
			summaryEnd = i
			break
		}
	}
	summary := strings.TrimSpace(
		strings.Join(releaseSegment[1:summaryEnd], "\n"),
	)

	// Extract changes
	changes := strings.TrimSpace(
		strings.Join(releaseSegment[summaryEnd:], "\n"),
	)

	return ReleaseChangelog{
		Summary: summary,
		Changes: changes,
	}, nil
}

func normalizeString(changeLog string) string {
	return strings.ReplaceAll(changeLog, "\r", "")
}
