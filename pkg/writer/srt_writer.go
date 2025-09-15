package writer

import (
	"fmt"
	"os"

	"github.com/mouuff/GoSubAI/pkg/types"
	"github.com/plunch/gosrt"
)

type SrtWriter struct {
}

func (p *SrtWriter) Write(outputFile string, subtitleData *types.SubtitleData) error {

	file, err := os.Create(outputFile)

	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()

	for _, entry := range subtitleData.Entries {
		subtitle := gosrt.Subtitle{
			Number: entry.Index,
			Start:  entry.Start,
			End:    entry.End,
			Text:   entry.Text,
		}

		_, err := subtitle.WriteTo(file)
		if err != nil {
			return fmt.Errorf("failed to write subtitle entry: %w", err)
		}
	}

	return nil
}
