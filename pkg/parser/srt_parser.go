package parser

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/mouuff/GoSubAI/pkg/types"
)

type SRTParser struct {
}

// Defines a block in an SRT file
type Block struct {
	Index int
	Text  string
}

func parseBlocks(input string) []Block {
	// Regex: a number at start of line, then everything until next number or end of text
	re := regexp.MustCompile(`(?m)^(\d+)\s*\n([\s\S]*?)(?=^\d+|\z)`)
	matches := re.FindAllStringSubmatch(input, -1)

	var blocks []Block
	for _, m := range matches {
		idx, err := strconv.Atoi(m[1])
		if err != nil {
			log.Println("Skipping block with invalid number:", err)
		}

		blocks = append(blocks, Block{
			Index: idx,
			Text:  m[2],
		})
	}
	return blocks
}

func (p *SRTParser) Parse(input string) (*types.SubtitleData, error) {
	blocks := parseBlocks(input)

	subtitleData := &types.SubtitleData{}

	for _, block := range blocks {
		lines := strings.SplitN(block.Text, "\n", 2)
		if len(lines) < 2 {
			log.Println("Skipping block with insufficient lines:", block.Index)
			continue
		}
	}

	return subtitleData, nil
}
