package llm

import (
	"context"
	"sync"

	"github.com/Mentro-Org/CodeLookout/internal/config"
	openai "github.com/sashabaranov/go-openai"
)

type AIClient interface {
	ReviewPR(ctx context.Context, prompt string) (string, error)
}

type OpenAIClient struct {
	client *openai.Client
	model  string
}

var (
	instance *OpenAIClient
	once     sync.Once
)

// NewOpenAIClient initializes the singleton instance
func NewOpenAIClient(cfg *config.Config) *OpenAIClient {
	once.Do(func() {
		c := openai.NewClient(cfg.OpenAIKey)
		instance = &OpenAIClient{client: c, model: openai.GPT4}
	})
	return instance
}

func (c *OpenAIClient) ReviewPR(ctx context.Context, prompt string) (string, error) {
	resp, err := c.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: c.model,
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "You are a senior code reviewer. Be concise."},
			{Role: openai.ChatMessageRoleUser, Content: prompt},
		},
	})
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
