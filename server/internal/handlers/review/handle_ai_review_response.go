package review

import (
    "context"
    "log"

    "github.com/nalindalal/CodeLookout/server/internal/core"
    "github.com/nalindalal/CodeLookout/server/internal/llm"
    "github.com/nalindalal/CodeLookout/server/internal/queue"
)

func HandleReviewResponseFromAI(ctx context.Context, payload queue.PRReviewTaskPayload, appDeps *core.AppDeps, aiJsonResponse string) error {
	reviewResp, err := llm.ParseReviewResponse(aiJsonResponse)
	if err != nil {
		log.Printf("[AI Review] Failed to parse LLM response: %v\nRaw response: %s", err, aiJsonResponse)
		// Post fallback comment to PR
		reviewCtx := core.ReviewContext{
			Ctx:     ctx,
			Payload: payload,
			AppDeps: appDeps,
		}
		fallback := &InlineComment{
			Body: "[CodeLookout] LLM review failed: unable to parse AI response. Please try again or contact support.",
			Path: "",
			StartLine: 0,
			Line: 0,
		}
		_ = fallback.Execute(&reviewCtx)
		return err
	}

	// Validate required fields
	if reviewResp.Action == "" || len(reviewResp.Comments) == 0 {
		log.Printf("[AI Review] LLM response missing required fields: %+v", reviewResp)
		reviewCtx := core.ReviewContext{
			Ctx:     ctx,
			Payload: payload,
			AppDeps: appDeps,
		}
		fallback := &InlineComment{
			Body: "[CodeLookout] LLM review failed: incomplete AI response. Please try again or contact support.",
			Path: "",
			StartLine: 0,
			Line: 0,
		}
		_ = fallback.Execute(&reviewCtx)
		return nil
	}

	reviewCtx := core.ReviewContext{
		Ctx:     ctx,
		Payload: payload,
		AppDeps: appDeps,
	}

	// Inline comments (multi-line support)
	for _, fileGroup := range reviewResp.Comments {
		for _, c := range fileGroup.Comments {
			// Format the comment body to include the line range if > 1 line
			commentBody := c.Body
			inline := &InlineComment{
				Body:      commentBody,
				Path:      fileGroup.Path,
				StartLine: c.Line.S,
				Line:      c.Line.E,
			}
			if err := inline.Execute(&reviewCtx); err != nil {
				log.Printf("Error posting inline comment on %s:%d: %v", fileGroup.Path, c.Line.S, err)
			}
		}
	}

	// Submit overall review event (APPROVE, REQUEST_CHANGES, COMMENT)
	reviewSubmission := &ReviewSubmission{
		Body:  reviewResp.Summary,
		Event: reviewResp.Action,
	}

	if err := reviewSubmission.Execute(&reviewCtx); err != nil {
		log.Printf("Error submitting review: %v", err)
		return err
	}

	return nil
}
