package types

import "time"

type SubtitleEntry struct {
	Index int
	Start time.Duration
	End   time.Duration
	Text  string
}

type SubtitleData struct {
	Entries []SubtitleEntry
}
