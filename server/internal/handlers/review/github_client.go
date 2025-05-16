package review

import (
	"context"
	"log"
	"net/http"

	"github.com/Mentro-Org/CodeLookout/internal/config"
	"github.com/bradleyfalzon/ghinstallation"

	"github.com/google/go-github/github"
)

func NewGitHubClient(ctx context.Context, cfg *config.Config, installationID int64) (*github.Client, error) {
	tr, err := ghinstallation.New(http.DefaultTransport, cfg.GithubAppID, installationID, []byte(cfg.GithubAppPrivateKey))
	if err != nil {
		log.Printf("Error creating GitHub App transport: %v\n", err)
		return nil, err
	}
	return github.NewClient(&http.Client{Transport: tr}), nil
}
