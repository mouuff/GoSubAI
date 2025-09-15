package generator

import (
	"context"
	"fmt"
	"strings"

	"github.com/mouuff/GoSubAI/pkg/brain"
	"github.com/mouuff/GoSubAI/pkg/types"
)

type GenerationType int32

type SubtitleGenerator struct {
	context       context.Context
	Brain         *brain.OllamaBrain
	SubstitleData *types.SubtitleData
	PropertyName  string
	Prompt        string
	Template      string
	Debug         bool
}

func (g *SubtitleGenerator) Generate() (*types.SubtitleData, error) {
	resultSubtitleData := &types.SubtitleData{
		Entries: []types.SubtitleEntry{},
	}

	for _, entry := range g.SubstitleData.Entries {
		prompt := strings.Replace(g.Prompt, "{TEXT}", entry.Text, 1)
		response, err := g.Brain.GenerateString(g.context, g.PropertyName, prompt)

		if err != nil {
			return nil, err
		}

		resultText := g.Template
		resultText = strings.Replace(resultText, "{ORIGINAL_TEXT}", entry.Text, 1)
		resultText = strings.Replace(resultText, "{GENERATED_TEXT}", response, 1)
		resultText = strings.Replace(resultText, "{PROMPT}", prompt, 1) // For debugging

		resultSubtitleData.Entries = append(resultSubtitleData.Entries, types.SubtitleEntry{
			Index: entry.Index,
			Start: entry.Start,
			End:   entry.End,
			Text:  resultText,
		})

		if g.Debug {
			fmt.Printf("Index: %d\n", entry.Index)
			fmt.Printf("Prompt: %s\n", prompt)
			fmt.Printf("Response: %s\n", response)
			fmt.Printf("ResultText: %s\n", resultText)
		}
	}

	return g.SubstitleData, nil
}
