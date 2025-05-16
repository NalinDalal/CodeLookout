package pullrequest

import (
	"context"
	"log"
	"net/http"

	"github.com/Mentro-Org/CodeLookout/internal/config"
	ghclient "github.com/Mentro-Org/CodeLookout/internal/github"
	"github.com/Mentro-Org/CodeLookout/internal/handlers/review"
	"github.com/Mentro-Org/CodeLookout/internal/llm"
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

func HandleReviewForPR(ctx context.Context, event *github.PullRequestEvent, cfg *config.Config, ghClientFactory *ghclient.ClientFactory, aIClient llm.AIClient) error {
	log.Printf("Received a pull request event for #%d\n", event.GetNumber())

	// prompt := llm.BuildPRReviewPrompt(event)

	review.CommentSelector("general").Execute(ctx, event, cfg, ghClientFactory)
	review.CommentSelector("inline").Execute(ctx, event, cfg, ghClientFactory)
	review.CommentSelector("review").Execute(ctx, event, cfg, ghClientFactory)

	return nil
}
