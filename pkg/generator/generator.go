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
	Config        *types.GeneratorConfig
}

func (g *SubtitleGenerator) Generate() (*types.SubtitleData, error) {
	resultSubtitleData := &types.SubtitleData{
		Entries: []types.SubtitleEntry{},
	}

	total := len(g.SubstitleData.Entries)
	var previousText string

	for _, entry := range g.SubstitleData.Entries {
		prompt := strings.ReplaceAll(g.Config.Prompt, "{TEXT}", entry.Text)

		if previousText != "" {
			// Can be used to provide context from previous subtitle for better translations
			prompt = strings.ReplaceAll(prompt, "{PREVIOUS_TEXT}", previousText)
		}

		r := &types.PromptRequest{
			PropertyName: g.Config.PropertyName,
			SystemPrompt: g.Config.SystemPrompt,
			Model:        g.Config.Model,
			Prompt:       prompt,
		}

		response, err := g.Brain.GenerateString(g.Context, r)

		if err != nil {
			return nil, err
		}

		response = strings.TrimSpace(response)
		response = strings.Trim(response, "\"")
		response = strings.Trim(response, "'")

		resultText := g.Config.Template
		resultText = strings.ReplaceAll(resultText, "{TEXT}", entry.Text)
		resultText = strings.ReplaceAll(resultText, "{GENERATED_TEXT}", response)

		resultSubtitleData.Entries = append(resultSubtitleData.Entries, types.SubtitleEntry{
			Index: entry.Index,
			Start: entry.Start,
			End:   entry.End,
			Text:  resultText,
		})

		if g.Config.Debug {
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
