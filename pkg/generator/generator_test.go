package generator_test

import (
	"context"
	"testing"
	"time"

	"github.com/mouuff/GoSubAI/pkg/generator"
	"github.com/mouuff/GoSubAI/pkg/types"
)

type MockBrain struct{}

func (mb *MockBrain) GenerateString(ctx context.Context, r *types.PromptRequest) (string, error) {
	// Simple mock implementation that just returns a fixed string for testing
	return "This is a translated text.", nil
}

func TestGenerate(t *testing.T) {

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
				Text:  "Som igjen vil føre til at alt blirbilligere i velferdsstaten Norge.",
			},
		},
	}

	g := generator.SubtitleGenerator{
		Brain:         &MockBrain{},
		Context:       context.Background(),
		SubstitleData: subtitleData,
		Config: &types.GeneratorConfig{
			PropertyName: "translated_text",
			SystemPrompt: "You are a translation assistant. Your only task is to translate any input text into clear and natural English. Do not add explanations, comments, or extra details—only provide the translation.",
			Prompt:       "Translate this to english: '{TEXT}'",
			Template:     "{TEXT}\n----\n{GENERATED_TEXT}",
		},
	}

	result, err := g.Generate()
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	if len(result.Entries) != 2 {
		t.Fatalf("Expected 2 entries, got %d", len(result.Entries))
	}

	expectedFirstEntryText := "Jeg så for øvrig Bjørnar\nhadde vært uheldig.\n----\nThis is a translated text."
	if result.Entries[0].Text != expectedFirstEntryText {
		t.Errorf("Unexpected text for first entry:\nGot: %s\nWant: %s", result.Entries[0].Text, expectedFirstEntryText)
	}

	if result.Entries[0].Index != 1 || result.Entries[0].Start != 1760*time.Millisecond || result.Entries[0].End != 4000*time.Millisecond {
		t.Errorf("Unexpected values for first entry: %+v", result.Entries[0])
	}

	expectedSecondEntryText := "Som igjen vil føre til at alt blirbilligere i velferdsstaten Norge.\n----\nThis is a translated text."
	if result.Entries[1].Text != expectedSecondEntryText {
		t.Errorf("Unexpected text for second entry:\nGot: %s\nWant: %s", result.Entries[1].Text, expectedSecondEntryText)
	}

	if result.Entries[1].Index != 2 || result.Entries[1].Start != 4000*time.Millisecond || result.Entries[1].End != 8360*time.Millisecond {
		t.Errorf("Unexpected values for second entry: %+v", result.Entries[1])
	}
}
