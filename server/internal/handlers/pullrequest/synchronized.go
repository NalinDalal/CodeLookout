package pullrequest

import (
	"context"

	"github.com/Mentro-Org/CodeLookout/internal/config"
	"github.com/google/go-github/github"
)

type PullRequestSynchronizedHandler struct {
	Cfg *config.Config
}

func (h *PullRequestSynchronizedHandler) Handle(ctx context.Context, event *github.PullRequestEvent) error {
	return HandleReviewForPR(ctx, event, h.Cfg)
}
