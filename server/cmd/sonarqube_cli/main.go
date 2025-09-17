package main

import (
    "fmt"
    "os"

    "github.com/nalindalal/CodeLookout/server/internal/config"
    "github.com/nalindalal/CodeLookout/server/internal/llm"
)

// This CLI demonstrates SonarQube integration using SonarQubeClient
func main() {
	cfg := config.Load()
	if cfg.SonarQubeEndpoint == "" || cfg.SonarQubeToken == "" {
		fmt.Println("SONARQUBE_ENDPOINT or SONARQUBE_TOKEN not set in config. Exiting.")
		os.Exit(1)
	}
	client := &llm.SonarQubeClient{Endpoint: cfg.SonarQubeEndpoint, Token: cfg.SonarQubeToken}
	codePath := "./" // Example: analyze current directory
	fmt.Println("Sending code path to SonarQube endpoint:", cfg.SonarQubeEndpoint)
	result, err := client.AnalyzeCode(codePath)
	if err != nil {
		fmt.Println("SonarQube error:", err)
		os.Exit(1)
	}
	fmt.Println("SonarQube analysis result:")
	fmt.Println(result)
}
