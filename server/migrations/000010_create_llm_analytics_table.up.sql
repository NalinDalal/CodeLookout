-- +goose Up
CREATE TABLE IF NOT EXISTS llm_analytics (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    prompt TEXT NOT NULL,
    response TEXT,
    duration_ms INTEGER NOT NULL,
    error TEXT,
    pr_number INTEGER,
    repo TEXT,
    owner TEXT
);
