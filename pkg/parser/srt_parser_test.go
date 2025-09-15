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

	firstEntry := subtitleData.Entries[0]
	if firstEntry.Index != 1 {
		t.Errorf("Expected first entry index to be 1, got %d", firstEntry.Index)
	}
	if firstEntry.Text != "Jeg så for øvrig Bjørnar\nhadde vært uheldig." {
		t.Errorf("Unexpected first entry text: %s", firstEntry.Text)
	}
	if firstEntry.Start.String() != "1.76s" {
		t.Errorf("Unexpected first entry start time: %s", firstEntry.Start.String())
	}
	if firstEntry.End.String() != "4s" {
		t.Errorf("Unexpected first entry end time: %s", firstEntry.End.String())
	}

	secondEntry := subtitleData.Entries[1]
	if secondEntry.Index != 2 {
		t.Errorf("Expected second entry index to be 2, got %d", secondEntry.Index)
	}
	if secondEntry.Text != "Ja, han falt på sykkel og brakk armen." {
		t.Errorf("Unexpected second entry text: %s", secondEntry.Text)
	}
	if secondEntry.Start.String() != "4s" {
		t.Errorf("Unexpected second entry start time: %s", secondEntry.Start.String())
	}
	if secondEntry.End.String() != "8.36s" {
		t.Errorf("Unexpected second entry end time: %s", secondEntry.End.String())
	}
}
