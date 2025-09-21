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

func trimGeneratedText(s string) string {
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
