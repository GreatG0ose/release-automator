package fullrelease_message

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/GreatG0ose/release-automator/internal/config"
	"github.com/GreatG0ose/release-automator/internal/release"
	"github.com/atc0005/go-teams-notify/v2/messagecard"
	"text/template"
)

//go:embed message.gotmpl
var tmpl string

// TODO: add docs
func createMessage(cfg config.Config, r release.Release) (*messagecard.MessageCard, error) {
	m := messagecard.NewMessageCard()
	m.Title = fmt.Sprintf("%s %s", cfg.Project.Name, r.Version)

	text, err := renderMessage(cfg, r)
	if err != nil {
		return nil, err
	}

	m.Text = text

	return m, nil
}

// TODO: add docs
func renderMessage(cfg config.Config, r release.Release) (string, error) {
	type tmplCtx struct {
		Config  config.Config
		Release release.Release
	}

	t, err := template.New("message").Parse(tmpl)
	if err != nil {
		// template is totally fine. this branch must not be executed on production
		panic(err)
	}

	var rendered bytes.Buffer
	err = t.Execute(&rendered, tmplCtx{
		Config:  cfg,
		Release: r,
	})
	if err != nil {
		return "", fmt.Errorf("failed to render message: %w", err)
	}

	return rendered.String(), nil
}
