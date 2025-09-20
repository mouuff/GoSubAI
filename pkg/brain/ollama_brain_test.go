package brain_test

import (
	"context"
	"strings"
	"testing"

	"github.com/mouuff/GoSubAI/pkg/brain"
	"github.com/mouuff/GoSubAI/pkg/types"
)

func TestOllamaBrainGenerateString(t *testing.T) {
	ctx := context.Background()
	gen, err := brain.NewOllamaBrain("")

	if err != nil {
		t.Fatal(err)
	}

	r := &types.PromptRequest{
		Model:        "llama3.2",
		SystemPrompt: "You are a subtitle translation assistant. Your only task is to translate subtitles into the target language specified by the user. Subtitles may contain incomplete sentences-when that happens, translate them literally without trying to complete or alter their meaning. Always keep the translation faithful to the original text and do not add explanations or extra words.",
		PropertyName: "translated_text",
		Prompt:       "Translate this to french: 'hello'",
	}

	result, err := gen.GenerateString(ctx, r)
	lowerResult := strings.ToLower(result)

	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(lowerResult, "bonjour") {
		t.Fatal("Did not get expected translation, got: " + lowerResult)
	}
}
