package release_notes

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

const changesTmpl = `{{ .Changes.Summary }}
{{ range .Changes.Changes}}
## {{ .Header }}

{{ .Body }}
{{ end -}}`

// Generate generates Markdown formatted file in Config.Output directory
func Generate(l zerolog.Logger, c config.Config, r release.Release) error {
	tmpl := template.Must(template.New("changes").Parse(changesTmpl))
	outputFile := filepath.Join(c.Output, "ReleaseNotes.md")

	l.Info().Msg("rendering release notes...")
	var data bytes.Buffer
	err := tmpl.Execute(&data, r)
	if err != nil {
		return fmt.Errorf("cannot render changes: %w", err)
	}

	l.Info().Msgf("generating output file %s", outputFile)
	err = files.WriteToFile(outputFile, data.Bytes())
	if err != nil {
		return fmt.Errorf("failed to generate file: %w", err)
	}

	return nil
}
