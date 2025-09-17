package llm

import (
	"fmt"

	"github.com/Mentro-Org/CodeLookout/internal/config"
)

// NewClient returns an AIClient based on configuration (OpenAI, REST LLM, etc.)
func NewClient(cfg *config.Config) (AIClient, error) {
       switch cfg.AIProvider {
       case "openai":
	       return NewOpenAIClient(cfg), nil
       case "restllm":
	       // Example: REST-based LLM (HuggingFace, Ollama, etc.)
	       return &RESTLLMClient{Endpoint: cfg.LLMEndpoint}, nil
       case "sonarqube":
	       // SonarQube integration (see go-static-analyzers-research.md)
	       // Requires SONARQUBE_ENDPOINT and SONARQUBE_TOKEN in config
	       return &SonarQubeClient{Endpoint: cfg.SonarQubeEndpoint, Token: cfg.SonarQubeToken}, nil
       default:
	       return nil, fmt.Errorf("unsupported AI provider: %s", cfg.AIProvider)
       }
}
