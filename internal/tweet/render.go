package tweet

import (
	"bytes"
	"fmt"
	"github.com/GreatG0ose/release-automator/internal/config"
	"github.com/GreatG0ose/release-automator/internal/release"
	"github.com/rs/zerolog"
	"os"
	"path/filepath"
	"text/template"
)

const messageTmpl = `We just released version {{ .Release.Version }} of our {{ .Project.Name }}.
{{ .Release.Changes.Summary }}
Please read the changelog here
{{ .Project.Repository }}/blob/{{ .Release.Version }}/CHANGELOG.md

@java #ecommerce`

func Render(l zerolog.Logger, c config.Config, r release.Release) error {
	type tmplCtx struct {
		Project config.Project
		Release release.Release
	}

	tmpl := template.Must(template.New("tweet").Parse(messageTmpl))

	l.Info().Msg("Rendering tweet message...")
	var data bytes.Buffer
	err := tmpl.Execute(&data, tmplCtx{
		Project: c.Project,
		Release: r,
	})
	if err != nil {
		return fmt.Errorf("cannot render tweet: %w", err)
	}

	outputFile := filepath.Join(c.Output, "tweet.txt")
	l.Info().Msgf("Generating output file %s", outputFile)
	err = os.MkdirAll(filepath.Dir(c.Output), 0744)
	if err != nil {
		return fmt.Errorf("failed to create dirs: %w", err)
	}

	f, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("cannot generate output file: %w", err)
	}
	defer f.Close()

	_, err = f.Write(data.Bytes())
	if err != nil {
		return fmt.Errorf("cannot write tweet to file: %w", err)
	}
	return nil
}
