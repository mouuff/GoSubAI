package parser

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/mouuff/GoSubAI/pkg/types"
)

type SrtParser struct {
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
	re := regexp.MustCompile(`(?ms)^(\d+)\n(.*?)(?:\n(?=\d+\n)|\z)`)
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

func (p *SrtParser) Parse(input string) (*types.SubtitleData, error) {
	subtitleData := &types.SubtitleData{
		Entries: []types.SubtitleEntry{},
	}

	// Regex to find all the numbers at the start of lines
	reNumber := regexp.MustCompile(`(?m)^\d+$`)
	numbers := reNumber.FindAllString(input, -1)

	// Split the input on lines that are just numbers
	reSplit := regexp.MustCompile(`(?m)^\d+$`)
	parts := reSplit.Split(input, -1)

	// The first split part is empty (before the first number)
	for i, part := range parts[1:] {
		index, err := strconv.Atoi(numbers[i])
		if err != nil {
			log.Println("Skipping block with invalid index:", numbers[i], err)
			continue
		}
		content := strings.TrimSpace(part)
		lines := strings.SplitN(content, "\n", 2)
		if len(lines) < 2 {
			log.Println("Skipping block with insufficient lines:", index)
			continue
		}

		text := strings.TrimSpace(lines[1])
		timeParts := strings.SplitN(lines[0], " --> ", 2)
		if len(timeParts) != 2 {
			log.Println("Skipping block with invalid time format:", index)
			continue
		}
		startTime, err := parseTimecode(strings.TrimSpace(timeParts[0]))
		if err != nil {
			log.Println("Skipping block with invalid start time:", index, err)
			continue
		}
		endTime, err := parseTimecode(strings.TrimSpace(timeParts[1]))
		if err != nil {
			log.Println("Skipping block with invalid end time:", index, err)
			continue
		}

		subtitleData.Entries = append(subtitleData.Entries, types.SubtitleEntry{
			Index: index,
			Start: time.Duration(startTime) * time.Millisecond,
			End:   time.Duration(endTime) * time.Millisecond,
			Text:  text,
		})
	}

	return subtitleData, nil
}
