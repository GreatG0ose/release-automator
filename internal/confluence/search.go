package confluence

import (
	"fmt"
	"github.com/GreatG0ose/release-automator/internal/config"
	"github.com/virtomize/confluence-go-api"
	"net/url"
	"strings"
)

// GetReleasePageLink searches Confluence Release list page for child page containing target version in title
func GetReleasePageLink(cfg config.Config, version string) (string, error) {
	api, err := goconfluence.NewAPI(
		cfg.Confluence.Endpoint,
		cfg.Confluence.Credentials.Username,
		cfg.Confluence.Credentials.AccessToken,
	)
	if err != nil {
		return "", fmt.Errorf("failed to connect to confluence: %w", err)
	}

	searchResp, err := api.GetChildPages(cfg.Confluence.ReleasesPageId)
	if err != nil {
		return "", fmt.Errorf("failed to find releases page: %w", err)
	}

	for _, p := range searchResp.Results {
		if strings.Contains(p.Title, version) && strings.Contains(p.Title, cfg.Project.Name) {
			return buildPageUrl(cfg.Confluence.Endpoint, p.ID)
		}
	}

	return "", fmt.Errorf("release page not found for project %s version %s", cfg.Project.Name, version)
}

func buildPageUrl(endpoint, pageId string) (string, error) {
	rootUrl, err := url.Parse(endpoint)
	if err != nil {
		panic(err) // if url is broken, client couldn't connect to confluence
	}

	return fmt.Sprintf("https://%s/wiki/pages/viewpage.action?pageId=%s", rootUrl.Host, pageId), nil
}
