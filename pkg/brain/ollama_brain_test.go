package brain_test

import (
	"context"
	"strings"
	"testing"

	"github.com/mouuff/GoSubAI/pkg/brain"
)

func TestOllamaBrainGenerateString(t *testing.T) {
	ctx := context.Background()
	gen, err := brain.NewOllamaBrain("llama3.2")

	if err != nil {
		t.Fatal(err)
	}

	baseprompt := "Translate this to french: 'Hello'"
	propertyName := "translated_text"

	result, err := gen.GenerateString(ctx, propertyName, baseprompt)
	lowerResult := strings.ToLower(result)

	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(lowerResult, "bonjour") {
		t.Fatal("Did not get expected translation, got: " + lowerResult)
	}
}
