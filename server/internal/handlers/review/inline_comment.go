package review

import (
	"context"
	"log"

	"github.com/Mentro-Org/CodeLookout/internal/config"
	"github.com/google/go-github/github"
)

type InlineComment struct {
	Body     string
	Path     string // file path
	Position int    // line number or position in diff
}

func (ic *InlineComment) Execute(ctx context.Context, event *github.PullRequestEvent, cfg *config.Config) error {
	client, err := NewGitHubClient(ctx, cfg, event.GetInstallation().GetID())
	if err != nil {
		return err
	}

	commitSHA := event.GetPullRequest().GetHead().GetSHA()

	comment := &github.PullRequestComment{
		Body:     github.String(ic.Body),
		CommitID: github.String(commitSHA),
		Path:     github.String(ic.Path),
		Position: github.Int(ic.Position),
	}

	_, _, err = client.PullRequests.CreateComment(ctx,
		event.GetRepo().GetOwner().GetLogin(),
		event.GetRepo().GetName(),
		event.GetNumber(),
		comment,
	)

	if err != nil {
		log.Printf("Failed to create inline comment: %v\n", err)
	}
	return err
}
