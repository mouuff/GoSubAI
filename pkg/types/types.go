package types

import (
	"context"
	"time"
)

type GeneratorConfig struct {
	HostUrl      string
	Model        string
	PropertyName string
	SystemPrompt string
	Prompt       string
	Template     string
	Debug        bool
}

// PromptConfig represents a prompt request
type PromptRequest struct {
	Model        string
	SystemPrompt string
	Prompt       string
	PropertyName string
}

type SubtitleEntry struct {
	Index int
	Start time.Duration
	End   time.Duration
	Text  string
}

type SubtitleData struct {
	Entries []SubtitleEntry
}

type SubtitleParser interface {
	Read(inputFile string) (*SubtitleData, error)
}

type SubtitleWriter interface {
	Write(outputFile string, subtitleData *SubtitleData) error
}

type Brain interface {
	GenerateString(ctx context.Context, r *PromptRequest) (string, error)
}
