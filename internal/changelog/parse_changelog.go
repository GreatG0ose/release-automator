package changelog

import (
	"fmt"
	"github.com/GreatG0ose/release-automator/internal/config"
	"os"
	"sort"
	"strings"
)

// ReleaseChanges contains summary of a release and it's changes
type ReleaseChanges struct {
	Summary string        // Overall summary of the release
	Changes []ChangeBlock // Headers to markdown formatted changes
}

// ChangeBlock contains header and related text block
type ChangeBlock struct {
	Header string // Header of the change block
	Body   string // Body of the change block
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
	releaseEnd := len(lines)
	for i := releaseStart + 1; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], "## ") {
			releaseEnd = i
			break
		}
	}

	releaseSegment := lines[releaseStart:releaseEnd]

	// Extract summary
	summary := ""

	summaryEnd := -1
	for i, l := range releaseSegment[1:] {
		if strings.HasPrefix(l, "### ") {
			summaryEnd = i
			break
		}
	}

	if summaryEnd > 0 {
		summary = strings.TrimSpace(
			strings.Join(releaseSegment[1:summaryEnd], "\n"),
		)
	}

	return ReleaseChanges{
		Summary: summary,
		Changes: parseChanges(releaseSegment[summaryEnd:]),
	}, nil
}

// parseChanges extracts headers and related text blocks from version section of changelog
func parseChanges(changes []string) []ChangeBlock {
	mappedBlocks := make(map[string][]string)

	var headers []string
	currentHeader := ""
	for _, l := range changes {
		if strings.HasPrefix(l, "### ") {
			currentHeader = strings.TrimSpace(strings.TrimPrefix(l, "###"))
			headers = append(headers, currentHeader)
		} else if currentHeader == "" {
			continue
		} else {
			mappedBlocks[currentHeader] = append(mappedBlocks[currentHeader], l)
		}
	}

	headerToContentMap := make(map[string]string)
	for h, b := range mappedBlocks {
		headerToContentMap[h] = strings.TrimSpace(strings.Join(b, "\n"))
	}
	sort.Strings(headers)

	var result []ChangeBlock
	for _, h := range headers {
		result = append(result, ChangeBlock{
			Header: h,
			Body:   headerToContentMap[h],
		})
	}

	return result
}

func normalizeString(changeLog string) string {
	return strings.ReplaceAll(changeLog, "\r", "")
}
