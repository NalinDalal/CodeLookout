package config

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
)

// Removed duplicate struct declaration
type Config struct {
	Port              string `envconfig:"PORT" default:"8080"`
	AppEnv            string `envconfig:"APP_ENV" default:"development"`
	GithubAppID       int64  `envconfig:"GITHUB_APP_ID" required:"true"`
	AIProvider        string `envconfig:"AI_PROVIDER" required:"true"`
	OpenAIKey         string `envconfig:"OPENAI_API_KEY" required:"true"`
	DatabaseURL       string `envconfig:"DATABASE_URL" required:"true"`
	RedisAddress      string `envconfig:"REDIS_ADDRESS" required:"true"`
	WorkerConcurrency int    `envconfig:"WORKER_CONCURRENCY" required:"true"`
	QueueSize         int    `envconfig:"QUEUE_SIZE" default:"100"`

	WebhookSecret           string `envconfig:"WEBHOOK_SECRET" required:"true"`
	GithubAppPrivateKeyPath string `envconfig:"GITHUB_APP_PRIVATE_KEY_PATH" required:"true"`

	// LLM REST endpoint for HuggingFace, Ollama, etc.
	LLMEndpoint string `envconfig:"LLM_ENDPOINT"`

	// SonarQube integration
	SonarQubeEndpoint string `envconfig:"SONARQUBE_ENDPOINT"`
	SonarQubeToken    string `envconfig:"SONARQUBE_TOKEN"`

	GithubAppPrivateKey []byte `ignored:"true"` // not from env
}

func Load() *Config {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Load GitHub App private key
	privateKeyPEM, err := os.ReadFile(cfg.GithubAppPrivateKeyPath)
	if err != nil {
		log.Fatalf("Failed to read private key at %s: %v", cfg.GithubAppPrivateKeyPath, err)
	}
	cfg.GithubAppPrivateKey = privateKeyPEM

	return &cfg
}
