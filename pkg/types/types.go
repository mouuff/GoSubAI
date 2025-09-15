package types

type SubtitleEntry struct {
	Index     int
	StartTime int
	EndTime   int
	Text      string
}

type SubtitleData struct {
	Entries []SubtitleEntry
}
