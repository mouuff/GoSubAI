package parser_test

import (
	"path/filepath"
	"testing"

	parser "github.com/mouuff/GoSubAI/pkg/parser"
)

func TestParseSrtFile(t *testing.T) {

	testSrtFile := filepath.Join("testdata", "subtitles_1.srt")

	parser := parser.SrtParser{}

	subtitleData, err := parser.Parse(testSrtFile)

	if err != nil {
		t.Errorf("Failed to parse SRT file: %v", err)
	}

	if len(subtitleData.Entries) != 922 {
		t.Errorf("Expected 922 subtitle entries, got %d", len(subtitleData.Entries))
	}
}
