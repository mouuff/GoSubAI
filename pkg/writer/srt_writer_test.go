package writer_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/mouuff/GoSubAI/internal/fileio"
	"github.com/mouuff/GoSubAI/pkg/types"
	"github.com/mouuff/GoSubAI/pkg/writer"
)

func TestParseSrtFile(t *testing.T) {
	tmpDir, err := fileio.TempDir()
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	defer os.RemoveAll(tmpDir)

	tmpDest := filepath.Join(tmpDir, "tmpDest")

	w := writer.SrtWriter{}

	subtitleData := &types.SubtitleData{
		Entries: []types.SubtitleEntry{
			{
				Index: 1,
				Start: 1760 * time.Millisecond,
				End:   4000 * time.Millisecond,
				Text:  "Jeg så for øvrig Bjørnar\nhadde vært uheldig.",
			},
			{
				Index: 2,
				Start: 4000 * time.Millisecond,
				End:   8360 * time.Millisecond,
				Text:  "Ja, han falt på sykkel og brakk armen.",
			},
		},
	}

	err = w.Write(tmpDest, subtitleData)
	if err != nil {
		t.Errorf("Failed to write SRT file: %v", err)
	}

	os.ReadFile(tmpDest)

	contentRaw, err := os.ReadFile(tmpDest)

	if err != nil {
		t.Errorf("Failed to read written SRT file: %v", err)
	}

	content := string(contentRaw)

	expectedContent := "1\n00:00:01,760 --> 00:00:04,000\nJeg så for øvrig Bjørnar\nhadde vært uheldig.\n\n2\n00:00:04,000 --> 00:00:08,360\nJa, han falt på sykkel og brakk armen.\n\n"
	if content != expectedContent {
		t.Errorf("Unexpected SRT file content:\n%s", string(content))
	}
}
