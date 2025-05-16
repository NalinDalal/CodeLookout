package review

import (
	"context"
	"log"

	"github.com/Mentro-Org/CodeLookout/internal/config"
	ghclient "github.com/Mentro-Org/CodeLookout/internal/github"

	"github.com/google/go-github/github"
)

type ReviewSubmission struct {
	Body     string
	Event    string // "APPROVE", "REQUEST_CHANGES", or "COMMENT"
	Comments []*github.DraftReviewComment
}

func (rs *ReviewSubmission) Execute(ctx context.Context, event *github.PullRequestEvent, cfg *config.Config, ghClientFactory *ghclient.ClientFactory) error {
	client, err := ghClientFactory.GetClient(ctx, event.GetInstallation().GetID())
	if err != nil {
		return err
	}

	review := &github.PullRequestReviewRequest{
		Body:     &rs.Body,
		Event:    &rs.Event,
		Comments: rs.Comments,
	}

	_, _, err = client.PullRequests.CreateReview(ctx,
		event.GetRepo().GetOwner().GetLogin(),
		event.GetRepo().GetName(),
		event.GetNumber(),
		review,
	)

	if err != nil {
		log.Printf("Failed to submit review: %v\n", err)
	}
	return err
}
