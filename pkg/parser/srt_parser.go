package parser

import (
	"fmt"
	"os"
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

func (p *SrtParser) Parse(input string) (*types.SubtitleData, error) {

	rawContent, err := os.ReadFile(input)
	if err != nil {
		fmt.Print(err)
	}

	stringContent := string(rawContent) // convert content to a 'string'
	for _, line := range strings.Split(strings.TrimSuffix(stringContent, "\n"), "\n") {
		fmt.Println(line)
	}

	// The first split part is empty (before the first number)
	for i, part := range parts[1:] {

		subtitleData.Entries = append(subtitleData.Entries, types.SubtitleEntry{
			Index: index,
			Start: time.Duration(startTime) * time.Millisecond,
			End:   time.Duration(endTime) * time.Millisecond,
			Text:  text,
		})
	}

	return subtitleData, nil
}
