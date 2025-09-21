package generator

import (
	"strings"

	"github.com/mouuff/GoSubAI/internal/constants"
)

type ReplacementValues struct {
	Text          string
	PreviousText  string
	GeneratedText string
}

// this allows custom prompts to provide context without including it in the response
func extractAfterDiv(s string) string {
	parts := strings.SplitN(s, "<div>", 2)
	if len(parts) == 2 {
		return parts[1]
	}
	return parts[0]
}

func trimGeneratedText(s string) string {
	s = extractAfterDiv(s)
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "\"")
	s = strings.Trim(s, "'")
	return s
}

func (v *ReplacementValues) ReplaceAll(s string) string {
	s = strings.ReplaceAll(s, constants.PlaceholderText, v.Text)
	s = strings.ReplaceAll(s, constants.PlaceholderPreviousText, v.PreviousText)
	s = strings.ReplaceAll(s, constants.PlaceholderGeneratedText, trimGeneratedText(v.GeneratedText))
	return s
}
