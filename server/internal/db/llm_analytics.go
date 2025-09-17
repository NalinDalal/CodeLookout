package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LLMAnalytics struct {
	ID        int64
	CreatedAt time.Time
	Prompt    string
	Response  string
	DurationMs int
	Error     string
	PRNumber  int
	Repo      string
	Owner     string
}

func InsertLLMAnalytics(ctx context.Context, db *pgxpool.Pool, a *LLMAnalytics) error {
	_, err := db.Exec(ctx, `
		INSERT INTO llm_analytics (prompt, response, duration_ms, error, pr_number, repo, owner)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, a.Prompt, a.Response, a.DurationMs, a.Error, a.PRNumber, a.Repo, a.Owner)
	return err
}
