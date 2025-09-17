package core

import (
    "context"

    "github.com/nalindalal/CodeLookout/server/internal/config"
    ghclient "github.com/nalindalal/CodeLookout/server/internal/github"
    "github.com/nalindalal/CodeLookout/server/internal/llm"
    "github.com/nalindalal/CodeLookout/server/internal/queue"
    "github.com/google/go-github/v72/github"
    "github.com/jackc/pgx/v5/pgxpool"
)

type AppDeps struct {
	Config          *config.Config
	GHClientFactory *ghclient.ClientFactory
	AIClient        llm.AIClient
	DBPool          *pgxpool.Pool
	TaskClient      *queue.TaskClient
}

type PullRequestHandler interface {
	Handle(ctx context.Context, event *github.PullRequestEvent) error
}

type ReviewAction interface {
	Execute(reviewCtx *ReviewContext) error
}

type ReviewContext struct {
	Ctx     context.Context
	Payload queue.PRReviewTaskPayload
	AppDeps *AppDeps
}
