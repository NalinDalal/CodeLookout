package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Mentro-Org/CodeLookout/internal/core"
	"github.com/Mentro-Org/CodeLookout/internal/handlers/review"
	"github.com/Mentro-Org/CodeLookout/internal/llm"
	"github.com/Mentro-Org/CodeLookout/internal/queue"
	db "github.com/Mentro-Org/CodeLookout/internal/db"
	"github.com/hibiken/asynq"
)

func HandleReviewForPR(ctx context.Context, t *asynq.Task, appDeps *core.AppDeps) error {
	fmt.Println("received job from Queue")

	var payload queue.PRReviewTaskPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}
	log.Printf("Received a pull request event for #%d\n", payload.PRNumber)
	// Get GitHub client
	client, err := appDeps.GHClientFactory.GetClient(ctx, payload.InstallationID)
	if err != nil {
		return err
	}

	// Get changed files
	files, _, err := client.PullRequests.ListFiles(ctx,
		payload.Owner,
		payload.Repo,
		payload.PRNumber,
		nil,
	)
	if err != nil {
		return err
	}

	// Build prompt using file diffs
	promptText := llm.BuildPRReviewPrompt(&payload, files)

	var response string
	var durationMs int
	var llmErr error
       if appDeps.Config.AppEnv == "development" {
	       response, llmErr = appDeps.AIClient.GenerateSampleReviewForPR()
	       if llmErr != nil {
		       return llmErr
	       }
       } else {
	       start := time.Now()
	       response, llmErr = appDeps.AIClient.GenerateReviewForPR(ctx, promptText)
	       durationMs = int(time.Since(start).Milliseconds())
       }

	// Persist LLM analytics (always, including errors)
	_ = db.InsertLLMAnalytics(ctx, appDeps.DBPool, &db.LLMAnalytics{
		Prompt:    promptText,
		Response:  response,
		DurationMs: durationMs,
		Error:     errToString(llmErr),
		PRNumber:  payload.PRNumber,
		Repo:      payload.Repo,
		Owner:     payload.Owner,
	})

       if llmErr != nil {
	       return llmErr
       }

       err = review.HandleReviewResponseFromAI(ctx, payload, appDeps, response)
       if err != nil {
	       return err
       }

       return nil
}

// errToString returns the error message or empty string
func errToString(err error) string {
       if err == nil {
	       return ""
       }
       return err.Error()
}
