package parser_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	parser "github.com/mouuff/GoSubAI/pkg/parser"
)

func TestParseSrtFile(t *testing.T) {

	testSrt := filepath.Join("testdata", "subtitles_1.srt")

	b, err := os.ReadFile(testSrt)
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'

	parser := parser.SrtParser{}

	subtitleData, err := parser.Parse(str)

	if err != nil {
		t.Errorf("Failed to parse SRT file: %v", err)
	}

	if len(subtitleData.Entries) != 922 {
		t.Errorf("Expected 922 subtitle entries, got %d", len(subtitleData.Entries))
	}
}
