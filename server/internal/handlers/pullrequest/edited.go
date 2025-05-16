package pullrequest

import (
	"context"

	"github.com/Mentro-Org/CodeLookout/internal/config"

	"github.com/google/go-github/github"
)

type PullRequestEditedHandler struct {
	Cfg *config.Config
}

func (h *PullRequestEditedHandler) Handle(ctx context.Context, event *github.PullRequestEvent) error {
	return HandleReviewForPR(ctx, event, h.Cfg)
}
