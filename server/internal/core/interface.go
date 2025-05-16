package core

import (
	"context"

	"github.com/Mentro-Org/CodeLookout/internal/config"
	"github.com/google/go-github/github"
)

type PullRequestHandler interface {
	Handle(ctx context.Context, event *github.PullRequestEvent) error
}

type ReviewAction interface {
	Execute(ctx context.Context, event *github.PullRequestEvent, cfg *config.Config) error
}
