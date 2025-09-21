package generator

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/mouuff/GoSubAI/pkg/types"
)

type GenerationType int32

type SubtitleGenerator struct {
	Context       context.Context
	Brain         types.Brain
	SubstitleData *types.SubtitleData
	Config        *types.GeneratorConfig
}

func (g *SubtitleGenerator) GenerateSingle(r *types.PromptRequest) (string, error) {
	if g.Config.Regex == "" {
		return g.Brain.GenerateString(g.Context, r)
	}

	attempts := 0
	for {
		generatedText, err := g.Brain.GenerateString(g.Context, r)

		if err != nil {
			return "", err
		}

		re := regexp.MustCompile(g.Config.Regex)
		matches := re.FindStringSubmatch(generatedText)

		if len(matches) == 2 {
			return matches[1], nil
		} else {
			log.Printf("GenerateSingle: could not match regex '%s'", g.Config.Regex)
		}

		if attempts >= g.Config.RegexRetryLimit {
			log.Printf("GenerateSingle: reached max attempts %d", g.Config.RegexRetryLimit)
			return generatedText, nil
		}

		attempts += 1
	}
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

		v.GeneratedText, err = g.GenerateSingle(r)
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
