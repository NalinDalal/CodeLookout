package llm

import (
	"fmt"

	"github.com/google/go-github/github"
)

// BuildPRReviewPrompt generates a prompt string using PR metadata and the changed files
func BuildPRReviewPrompt(event *github.PullRequestEvent, files []*github.CommitFile) string {
	title := event.GetPullRequest().GetTitle()
	body := event.GetPullRequest().GetBody()

	prompt := fmt.Sprintf(
		`You're an expert software engineer reviewing a pull request.

Title: %s
Description: %s

Please analyze the following file changes and provide code review feedback:
`, title, body)

	for _, f := range files {
		filename := f.GetFilename()
		patch := f.GetPatch()
		prompt += fmt.Sprintf("\nFile: %s\nDiff:\n%s\n", filename, patch)
	}

	prompt += "\nRespond with helpful review comments, suggested improvements, and bugs if any. Be concise and clear."
	return prompt
}
