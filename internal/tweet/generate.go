package tweet

import (
	"bytes"
	"fmt"
	"github.com/GreatG0ose/release-automator/internal/config"
	"github.com/GreatG0ose/release-automator/internal/files"
	"github.com/GreatG0ose/release-automator/internal/release"
	"github.com/rs/zerolog"
	"path/filepath"
	"text/template"
)

// Generate generates Tweet message in ${Config.Output}/tweet.txt
func Generate(l zerolog.Logger, c config.Config, r release.Release) error {
	l.Info().Msg("rendering tweet message...")
	msg, err := renderTweet(c, r)
	if err != nil {
		return fmt.Errorf("cannot render tweet: %w", err)
	}

	outputFile := filepath.Join(c.Output, "tweet.txt")
	l.Info().Msgf("generating output file %s", outputFile)
	err = files.WriteToFile(outputFile, msg)
	if err != nil {
		return fmt.Errorf("failed to generate tweet file: %w", err)
	}

	return nil
}

const messageTmpl = `We just released version {{ .Release.Version }} of our {{ .Project.Name }}.
{{ .Release.Changes.Summary }}
Please read the changelog here
{{ .Project.Repository }}/blob/{{ .Release.Version }}/CHANGELOG.md

#ecommerce`

// renderTweet renders tweet message using messageTmpl
func renderTweet(c config.Config, r release.Release) ([]byte, error) {
	type tmplCtx struct {
		Project config.Project
		Release release.Release
	}

	tmpl := template.Must(template.New("tweet").Parse(messageTmpl))
	var data bytes.Buffer
	err := tmpl.Execute(&data, tmplCtx{
		Project: c.Project,
		Release: r,
	})

	return data.Bytes(), err
}
