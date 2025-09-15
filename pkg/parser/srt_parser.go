package parser

import (
	"fmt"
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

var timePattern = regexp.MustCompile(`(\d{2}):(\d{2}):(\d{2}),(\d{3})`)

func parseTimecode(tc string) (int, error) {
	m := timePattern.FindStringSubmatch(tc)
	if m == nil {
		return 0, fmt.Errorf("invalid timecode: %s", tc)
	}
	h, _ := strconv.Atoi(m[1])
	min, _ := strconv.Atoi(m[2])
	sec, _ := strconv.Atoi(m[3])
	ms, _ := strconv.Atoi(m[4])
	return ((h*3600 + min*60 + sec) * 1000) + ms, nil
}

func parseBlocks(input string) ([]Block, error) {
	// Regex: a number at start of line, then everything until next number or end of text
	re := regexp.MustCompile(`(?m)^(\d+)\s*\n([\s\S]*?)(?=^\d+|\z)`)
	matches := re.FindAllStringSubmatch(input, -1)

	var blocks []Block
	for _, m := range matches {
		idx, err := strconv.Atoi(m[1])
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, Block{
			Index: idx,
			Text:  m[2],
		})
	}
	return blocks, nil
}

func (p *SRTParser) Parse(input string) (*types.SubtitleData, error) {
	blocks, err := parseBlocks(input)
	if err != nil {
		return nil, err
	}

	subtitleData := &types.SubtitleData{
		Entries: []types.SubtitleEntry{},
	}

	for _, block := range blocks {
		lines := strings.SplitN(block.Text, "\n", 2)
		if len(lines) < 2 {
			log.Println("Skipping block with insufficient lines:", block.Index)
			continue
		}

		text := strings.TrimSpace(lines[1])
		timeParts := strings.SplitN(lines[0], " --> ", 2)
		if len(timeParts) != 2 {
			log.Println("Skipping block with invalid time format:", block.Index)
			continue
		}
		startTime, err := parseTimecode(strings.TrimSpace(timeParts[0]))
		if err != nil {
			log.Println("Skipping block with invalid start time:", block.Index, err)
			continue
		}
		endTime, err := parseTimecode(strings.TrimSpace(timeParts[1]))
		if err != nil {
			log.Println("Skipping block with invalid end time:", block.Index, err)
			continue
		}

		subtitleData.Entries = append(subtitleData.Entries, types.SubtitleEntry{
			Index:     block.Index,
			StartTime: startTime,
			EndTime:   endTime,
			Text:      text,
		})
	}

	return subtitleData, nil
}
