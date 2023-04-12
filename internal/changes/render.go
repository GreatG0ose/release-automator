package changes

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

const changesTmpl = `{{ .Changes.Summary }}
{{ range $header, $content := .Changes.Changes}}
## {{ $header }}

{{ $content }}
{{ end -}}`

// GenerateVersionChangesFile generates Markdown formatted file in Config.Output directory
func GenerateVersionChangesFile(l zerolog.Logger, c config.Config, r release.Release) error {
	tmpl := template.Must(template.New("changes").Parse(changesTmpl))
	outputFile := filepath.Join(c.Output, "changes.md")

	l.Info().Msg("rendering changes...")

	var data bytes.Buffer
	err := tmpl.Execute(&data, r)
	if err != nil {
		return fmt.Errorf("cannot render changes: %w", err)
	}

	l.Info().Msgf("generating output file %s", outputFile)
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
		return fmt.Errorf("cannot write changes to file: %w", err)
	}
	return nil
}
