package llm

import (
       "context"
       "bytes"
       "encoding/json"
       "io"
       "net/http"
       "fmt"
       "time"
)

// Ensure SonarQubeClient implements AIClient for code review integration
func (c *SonarQubeClient) GenerateReviewForPR(ctx context.Context, prompt string) (string, error) {
       // TODO: Implement SonarQube analysis and return review string
       // This is a stub for integration with SonarQube REST API
       return c.AnalyzeCode(prompt)
}

func (c *SonarQubeClient) GenerateSampleReviewForPR() (string, error) {
       // TODO: Optionally load a sample review from disk or return a static example
       return "[Sample SonarQube review placeholder]", nil
}
// Example usage of LLM integration in Go
//
// func ExampleLLMUsage() {
//     // Configure your LLM endpoint (e.g., HuggingFace Inference API, Ollama, etc.)
//     llmClient := &RESTLLMClient{Endpoint: "http://localhost:11434/api/generate"}
//     code := "func add(a int, b int) int { return a + b }"
//     result, err := llmClient.AnalyzeCode(code)
//     if err != nil {
//         fmt.Println("LLM error:", err)
//         return
//     }
//     fmt.Println("LLM review result:", result)
// }
// Ensure RESTLLMClient implements AIClient for code review integration
func (c *RESTLLMClient) GenerateReviewForPR(ctx context.Context, prompt string) (string, error) {
       // Analytics: log prompt and timing
       start := time.Now()
       result, err := c.AnalyzeCode(prompt)
       duration := time.Since(start)
       if err != nil {
              fmt.Printf("[LLM] ERROR: %v\nPrompt: %s\nDuration: %v\n", err, prompt, duration)
       } else {
              fmt.Printf("[LLM] SUCCESS\nPrompt: %s\nResponse: %s\nDuration: %v\n", prompt, result, duration)
       }
       return result, err
}

func (c *RESTLLMClient) GenerateSampleReviewForPR() (string, error) {
	// TODO: Optionally load a sample review from disk or return a static example
	return "[Sample LLM review placeholder]", nil
}

// LLMClient defines the interface for interacting with an LLM service (e.g., HuggingFace, Ollama, etc.)
type LLMClient interface {
	// AnalyzeCode takes code and returns a review or suggestions.
	AnalyzeCode(code string) (string, error)
}

// RESTLLMClient is a client that calls an external LLM REST API (Python, HuggingFace, Ollama, etc.)
type RESTLLMClient struct {
	Endpoint  string
	AuthToken string // Optional: for HuggingFace, etc.
}

// AnalyzeCode sends code to the LLM REST API and returns the response.
func (c *RESTLLMClient) AnalyzeCode(code string) (string, error) {
       // Real implementation: POST code to LLM endpoint (e.g., HuggingFace, Ollama)
       payload := map[string]string{"inputs": code} // HuggingFace expects "inputs"
       body, err := json.Marshal(payload)
       if err != nil {
               fmt.Printf("[LLM] ERROR: failed to marshal payload: %v\n", err)
               return "", err
       }
       req, err := http.NewRequest("POST", c.Endpoint, bytes.NewBuffer(body))
       if err != nil {
               fmt.Printf("[LLM] ERROR: failed to create request: %v\n", err)
               return "", err
       }
       req.Header.Set("Content-Type", "application/json")
       if c.AuthToken != "" {
               req.Header.Set("Authorization", "Bearer "+c.AuthToken)
       }
       resp, err := http.DefaultClient.Do(req)
       if err != nil {
               fmt.Printf("[LLM] ERROR: request failed: %v\n", err)
               return "", err
       }
       defer resp.Body.Close()
       if resp.StatusCode != http.StatusOK {
               fmt.Printf("[LLM] ERROR: API returned status: %s\n", resp.Status)
               return "", fmt.Errorf("LLM API returned status: %s", resp.Status)
       }
       out, err := io.ReadAll(resp.Body)
       if err != nil {
               fmt.Printf("[LLM] ERROR: failed to read response: %v\n", err)
               return "", err
       }
       // HuggingFace returns JSON array of results, Ollama returns plain text
       fmt.Printf("[LLM] Raw response: %s\n", string(out))
       return string(out), nil
}
// --- SonarQube Integration (Static Analysis) ---
// See: server/go-static-analyzers-research.md for research and recommendations.
//
// SonarQubeClient connects to a real SonarQube server via REST API.
type SonarQubeClient struct {
	Endpoint string
	Token    string
}

// AnalyzeCode triggers SonarQube analysis and fetches results.
func (c *SonarQubeClient) AnalyzeCode(projectKey string) (string, error) {
       // Real implementation: Fetch issues for a project from SonarQube
       // projectKey: the SonarQube project key (must be scanned already)
       url := c.Endpoint + "/api/issues/search?componentKeys=" + projectKey
       req, err := http.NewRequest("GET", url, nil)
       if err != nil {
	       return "", err
       }
       if c.Token != "" {
	       req.Header.Set("Authorization", "Bearer "+c.Token)
       }
       resp, err := http.DefaultClient.Do(req)
       if err != nil {
	       return "", err
       }
       defer resp.Body.Close()
       if resp.StatusCode != http.StatusOK {
	       return "", fmt.Errorf("SonarQube API returned status: %s", resp.Status)
       }
       out, err := io.ReadAll(resp.Body)
       if err != nil {
	       return "", err
       }
       return string(out), nil
}


// --- SonarQube Integration (Static Analysis) ---
// See: server/go-static-analyzers-research.md for research and recommendations.
//
// To integrate SonarQube:
// 1. Deploy SonarQube server (self-hosted or cloud).
// 2. Use SonarQube scanner CLI or REST API to analyze codebase.
// 3. Implement a Go client (SonarQubeClient) that can trigger analysis and fetch results.
// 4. Parse and combine SonarQube results with LLM/AI review for unified feedback.
// 5. Post SonarQube findings as PR comments or status checks.
//
// Example stub for future implementation:
//
// type SonarQubeClient struct {
//     Endpoint string
//     Token    string
// }
//
// func (c *SonarQubeClient) AnalyzeCode(codePath string) (string, error) {
//     // TODO: Call SonarQube scanner or REST API, return results as string/JSON
//     return "[SonarQube analysis result placeholder]", nil
// }

// AIClient is the main interface for AI-powered code review (OpenAI, REST LLM, etc.)
type AIClient interface {
	GenerateReviewForPR(ctx context.Context, prompt string) (string, error)
	GenerateSampleReviewForPR() (string, error)
}
