package types

import (
	"context"
	"time"
)

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
	GenerateString(ctx context.Context, propertyName, prompt string) (string, error)
}
