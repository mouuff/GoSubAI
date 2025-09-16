package generator

import (
	"context"
	"fmt"
	"strings"

	"github.com/mouuff/GoSubAI/pkg/types"
)

type GenerationType int32

type SubtitleGenerator struct {
	Context       context.Context
	Brain         types.Brain
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

	total := len(g.SubstitleData.Entries)
	var previousText string

	for _, entry := range g.SubstitleData.Entries {
		prompt := strings.ReplaceAll(g.Prompt, "{TEXT}", entry.Text)

		if previousText != "" {
			// Can be used to provide context from previous subtitle for better translations
			prompt = strings.ReplaceAll(prompt, "{PREVIOUS_TEXT}", previousText)
		}

		response, err := g.Brain.GenerateString(g.Context, g.PropertyName, prompt)

		if err != nil {
			return nil, err
		}

		response = strings.TrimSpace(response)
		response = strings.Trim(response, "\"")
		response = strings.Trim(response, "'")

		resultText := g.Template
		resultText = strings.ReplaceAll(resultText, "{TEXT}", entry.Text)
		resultText = strings.ReplaceAll(resultText, "{GENERATED_TEXT}", response)

		resultSubtitleData.Entries = append(resultSubtitleData.Entries, types.SubtitleEntry{
			Index: entry.Index,
			Start: entry.Start,
			End:   entry.End,
			Text:  resultText,
		})

		if g.Debug {
			fmt.Printf("Index: %d / %d\n", entry.Index, total)
			fmt.Printf("Prompt:\n%s\n", prompt)
			fmt.Printf("Response:\n%s\n", response)
			fmt.Printf("ResultText:\n%s\n", resultText)
			fmt.Printf("************************\n")
		}

		previousText = entry.Text
	}

	return resultSubtitleData, nil
}
