package parser

import (
	"fmt"
	"os"

	"github.com/mouuff/GoSubAI/pkg/types"
	"github.com/plunch/gosrt"
)

type SrtParser struct {
}

func (p *SrtParser) Parse(inputFile string) (*types.SubtitleData, error) {

	file, err := os.Open(inputFile)

	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()

	scanner := gosrt.NewScanner(file)

	subtitleData := &types.SubtitleData{
		Entries: []types.SubtitleEntry{},
	}

	for scanner.Scan() {
		sub := scanner.Subtitle()

		subtitleData.Entries = append(subtitleData.Entries, types.SubtitleEntry{
			Index: sub.Number,
			Start: sub.Start,
			End:   sub.End,
			Text:  sub.Text,
		})
	}

	return subtitleData, nil
}
