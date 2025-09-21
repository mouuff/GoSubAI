package generator

import (
	"context"
	"fmt"
	"strings"

	"github.com/mouuff/GoSubAI/internal/constants"
	"github.com/mouuff/GoSubAI/pkg/types"
)

type GenerationType int32

type SubtitleGenerator struct {
	Context       context.Context
	Brain         types.Brain
	SubstitleData *types.SubtitleData
	Config        *types.GeneratorConfig
}

type ReplacementValues struct {
	Text          string
	PreviousText  string
	GeneratedText string
}

func trimGeneratedText(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "\"")
	s = strings.Trim(s, "'")
	return s
}

func (v *ReplacementValues) ReplaceAll(s string) string {
	s = strings.ReplaceAll(s, constants.PlaceholderText, v.Text)
	s = strings.ReplaceAll(s, constants.PlaceholderPreviousText, v.PreviousText)
	s = strings.ReplaceAll(s, constants.PlaceholderGeneratedText, trimGeneratedText(v.GeneratedText))
	return s
}

func (g *SubtitleGenerator) Generate() (*types.SubtitleData, error) {
	resultSubtitleData := &types.SubtitleData{
		Entries: []types.SubtitleEntry{},
	}

	total := len(g.SubstitleData.Entries)
	var previousText string

	for _, entry := range g.SubstitleData.Entries {
		var err error

		v := &ReplacementValues{
			Text:         entry.Text,
			PreviousText: previousText,
		}

		r := &types.PromptRequest{
			PropertyName: g.Config.PropertyName,
			SystemPrompt: g.Config.SystemPrompt,
			Model:        g.Config.Model,
			Prompt:       v.ReplaceAll(g.Config.Prompt),
		}

		v.GeneratedText, err = g.Brain.GenerateString(g.Context, r)

		if err != nil {
			return nil, err
		}

		subtitleEntry := types.SubtitleEntry{
			Index: entry.Index,
			Start: entry.Start,
			End:   entry.End,
			Text:  v.ReplaceAll(g.Config.Template),
		}

		resultSubtitleData.Entries = append(resultSubtitleData.Entries, subtitleEntry)

		if g.Config.Debug {
			fmt.Printf("Index: %d / %d\n", entry.Index, total)
			fmt.Printf("Prompt:\n%s\n", r.Prompt)
			fmt.Printf("Response:\n%s\n", v.GeneratedText)
			fmt.Printf("ResultText:\n%s\n", subtitleEntry.Text)
			fmt.Printf("************************\n")
		}

		previousText = entry.Text
	}

	return resultSubtitleData, nil
}
