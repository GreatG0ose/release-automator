package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"
)

//go:embed teams-message.gotmpl
var signOffTemplate string

type signOffMessageContext struct {
	Release  Release
	Mentions []string
}

func renderSignOffMessage(release Release, mentions []string) (string, error) {
	tmpl, err := template.New("SignOffMessage").Parse(signOffTemplate)

	if err != nil {
		return "", fmt.Errorf("unable to parse template: %w", err)
	}

	var templateOutput bytes.Buffer
	err = tmpl.Execute(&templateOutput, signOffMessageContext{
		Release:  release,
		Mentions: mentions,
	})
	if err != nil {
		return "", fmt.Errorf("unable to render message: %w", err)
	}

	return templateOutput.String(), nil
}
