package main

import (
       "fmt"
       "os"
       "github.com/Mentro-Org/CodeLookout/internal/config"
       "github.com/Mentro-Org/CodeLookout/internal/llm"
)

// This CLI demonstrates LLM integration using RESTLLMClient
func main() {
       cfg := config.Load()
       if cfg.LLMEndpoint == "" {
              fmt.Println("LLM_ENDPOINT not set in config. Exiting.")
              os.Exit(1)
       }
       // Optionally support LLM_AUTH_TOKEN for HuggingFace
       authToken := os.Getenv("LLM_AUTH_TOKEN")
       client := &llm.RESTLLMClient{Endpoint: cfg.LLMEndpoint, AuthToken: authToken}
       code := "func add(a int, b int) int { return a + b }"
       fmt.Println("Sending code to LLM endpoint:", cfg.LLMEndpoint)
       result, err := client.AnalyzeCode(code)
       if err != nil {
              fmt.Println("LLM error:", err)
              os.Exit(1)
       }
       fmt.Println("LLM review result:")
       fmt.Println(result)
}
