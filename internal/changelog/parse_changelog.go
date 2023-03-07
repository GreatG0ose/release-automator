package changelog

import (
	"fmt"
	"github.com/GreatG0ose/release-automator/internal/config"
	"os"
	"strings"
)

// ReleaseChanges contains summary of a release and it's changes
type ReleaseChanges struct {
	Summary string            // Overall summary of the release
	Changes map[string]string // Headers to markdown formatted changes
}

// ExtractReleaseChanges parses changelog and returns ReleaseChanges for version
func ExtractReleaseChanges(cfg config.Config, releaseVersion string) (ReleaseChanges, error) {
	changelogText, err := os.ReadFile(cfg.Project.ChangelogPath)
	if err != nil {
		return ReleaseChanges{}, fmt.Errorf("failed to read changelog %s: %w", cfg.Project.ChangelogPath, err)
	}

	lines := strings.Split(normalizeString(string(changelogText)), "\n")

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
		return ReleaseChanges{}, fmt.Errorf("release version '%s' not found in changelog", releaseVersion)
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
		if strings.HasPrefix(l, "### ") {
			summaryEnd = i
			break
		}
	}
	summary := strings.TrimSpace(
		strings.Join(releaseSegment[1:summaryEnd], "\n"),
	)

	return ReleaseChanges{
		Summary: summary,
		Changes: parseChanges(releaseSegment[summaryEnd:]),
	}, nil
}

// parseChanges extracts headers and related text blocks from version section of changelog
func parseChanges(changes []string) map[string]string {
	mappedBlocks := make(map[string][]string)

	currentHeader := ""
	for _, l := range changes {
		if strings.HasPrefix(l, "### ") {
			currentHeader = strings.TrimSpace(strings.TrimPrefix(l, "###"))
		} else if currentHeader == "" {
			continue
		} else {
			mappedBlocks[currentHeader] = append(mappedBlocks[currentHeader], l)
		}
	}

	result := make(map[string]string)
	for h, b := range mappedBlocks {
		result[h] = strings.TrimSpace(strings.Join(b, "\n"))
	}

	return result
}

func normalizeString(changeLog string) string {
	return strings.ReplaceAll(changeLog, "\r", "")
}
