package main

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
)

type ReleaseChangelog struct {
	Summary string // Overall summary of the release
	Changes string // Markdown formatted changes
}

func ExtractReleaseNotes(changelog string) (ReleaseChangelog, error) {
	reader := text.NewReader([]byte(changelog))
	changelogRoot := goldmark.New().Parser().Parse(reader)

	txt := changelogRoot.FirstChild()

	return ReleaseChangelog{}, nil
}
