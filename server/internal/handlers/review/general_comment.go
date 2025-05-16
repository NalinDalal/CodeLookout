package review

import (
	"context"
	"log"

	"github.com/Mentro-Org/CodeLookout/internal/config"
	"github.com/google/go-github/github"
)

type GeneralComment struct {
	Message string
}

func (gc *GeneralComment) Execute(ctx context.Context, event *github.PullRequestEvent, cfg *config.Config) error {
	log.Printf("Received a pull request event for #%d\n", event.GetNumber())

	client, err := NewGitHubClient(ctx, cfg, event.GetInstallation().GetID())

	_, _, err = client.Issues.CreateComment(
		ctx,
		event.GetRepo().GetOwner().GetLogin(),
		event.GetRepo().GetName(),
		event.GetNumber(),
		&github.IssueComment{Body: &gc.Message},
	)

	if err != nil {
		log.Printf("Failed to create general comment: %v\n", err)
	}
	return err
}
