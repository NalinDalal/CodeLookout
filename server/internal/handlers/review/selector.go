package review

import (
	"github.com/Mentro-Org/CodeLookout/internal/core"
	"github.com/google/go-github/github"
)

// Currently comments are static, for testing, will be dynamic based on review
func CommentSelector(commentTyep string) core.ReviewAction {
	switch commentTyep {
	case "general":
		return &GeneralComment{Message: "Hello"}
	case "inline":
		return &InlineComment{
			Body:     "Consider renaming this variable for clarity.",
			Path:     "main.go",
			Position: 18,
		}
	case "review":
		return &ReviewSubmission{
			Body:  "Looks Good to me!!",
			Event: "COMMENT",
			Comments: []*github.DraftReviewComment{
				{
					Path:     github.String("main.go"),
					Position: github.Int(5), // position of line in diff (not line number)
					Body:     github.String("Inline comment text"),
				},
			}}
	default:
		return nil
	}
}
