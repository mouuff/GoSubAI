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

	for _, entry := range g.SubstitleData.Entries {
		prompt := strings.Replace(g.Prompt, "{TEXT}", entry.Text, 1)
		response, err := g.Brain.GenerateString(g.Context, g.PropertyName, prompt)

		if err != nil {
			return nil, err
		}

		response = strings.TrimSpace(response)
		response = strings.Trim(response, "\"")
		response = strings.Trim(response, "'")

		resultText := g.Template
		resultText = strings.Replace(resultText, "{TEXT}", entry.Text, 1)
		resultText = strings.Replace(resultText, "{GENERATED_TEXT}", response, 1)

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
	}

	return resultSubtitleData, nil
}
