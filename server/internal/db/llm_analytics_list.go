package db

type LLMAnalyticsFilters struct {
	Error    string
	Repo     string
	Owner    string
	PRNumber string
	Since    string // ISO8601 or date string
	Until    string
}

func ListLLMAnalyticsFiltered(ctx context.Context, db *pgxpool.Pool, limit, offset int, f LLMAnalyticsFilters) ([]LLMAnalytics, error) {
       query := `SELECT id, created_at, prompt, response, duration_ms, error, pr_number, repo, owner FROM llm_analytics WHERE 1=1`
       args := []interface{}{}
       i := 1
       if f.Error != "" {
	       query += ` AND error ILIKE '%' || $` + strconv.Itoa(i) + ` || '%'`
	       args = append(args, f.Error)
	       i++
       }
       if f.Repo != "" {
	       query += ` AND repo = $` + strconv.Itoa(i)
	       args = append(args, f.Repo)
	       i++
       }
       if f.Owner != "" {
	       query += ` AND owner = $` + strconv.Itoa(i)
	       args = append(args, f.Owner)
	       i++
       }
       if f.PRNumber != "" {
	       query += ` AND pr_number = $` + strconv.Itoa(i)
	       args = append(args, f.PRNumber)
	       i++
       }
       if f.Since != "" {
	       query += ` AND created_at >= $` + strconv.Itoa(i)
	       args = append(args, f.Since)
	       i++
       }
       if f.Until != "" {
	       query += ` AND created_at <= $` + strconv.Itoa(i)
	       args = append(args, f.Until)
	       i++
       }
       query += ` ORDER BY created_at DESC LIMIT $` + strconv.Itoa(i) + ` OFFSET $` + strconv.Itoa(i+1)
       args = append(args, limit, offset)
       rows, err := db.Query(ctx, query, args...)
       if err != nil {
	       return nil, err
       }
       defer rows.Close()
       var results []LLMAnalytics
       for rows.Next() {
	       var a LLMAnalytics
	       if err := rows.Scan(&a.ID, &a.CreatedAt, &a.Prompt, &a.Response, &a.DurationMs, &a.Error, &a.PRNumber, &a.Repo, &a.Owner); err != nil {
		       return nil, err
	       }
	       results = append(results, a)
       }
       return results, nil
}
package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ListLLMAnalytics(ctx context.Context, db *pgxpool.Pool, limit, offset int) ([]LLMAnalytics, error) {
	rows, err := db.Query(ctx, `
		SELECT id, created_at, prompt, response, duration_ms, error, pr_number, repo, owner
		FROM llm_analytics
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []LLMAnalytics
	for rows.Next() {
		var a LLMAnalytics
		if err := rows.Scan(&a.ID, &a.CreatedAt, &a.Prompt, &a.Response, &a.DurationMs, &a.Error, &a.PRNumber, &a.Repo, &a.Owner); err != nil {
			return nil, err
		}
		results = append(results, a)
	}
	return results, nil
}
