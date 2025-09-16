package internal_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mouuff/GoSubAI/internal"
)

// A simple struct to test unmarshalling
type TestStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestReadFromJson_Success(t *testing.T) {
	// Create a temporary directory
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.json")

	// Write valid JSON to file
	content := `{"name":"Alice","age":30}`
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	var result TestStruct
	err := internal.ReadFromJson(filePath, &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Name != "Alice" || result.Age != 30 {
		t.Errorf("unexpected result: %+v", result)
	}
}

func TestAddPrefixToFilename(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		prefix   string
		want     string
	}{
		{
			name:     "simple filename with extension",
			filename: "file.txt",
			prefix:   "_new",
			want:     "file_new.txt",
		},
		{
			name:     "simple filename with path and extension",
			filename: "./data/HVOR_BLIR_DET_AV_PENGA.srt",
			prefix:   "_new",
			want:     "./data/HVOR_BLIR_DET_AV_PENGA_new.srt",
		},
		{
			name:     "simple filename with windows path and extension",
			filename: ".\\data\\HVOR_BLIR_DET_AV_PENGA.srt",
			prefix:   "_new",
			want:     ".\\data\\HVOR_BLIR_DET_AV_PENGA_new.srt",
		},
		{
			name:     "filename without extension",
			filename: "file",
			prefix:   "_new",
			want:     "file_new",
		},
		{
			name:     "filename with multiple dots",
			filename: "archive.tar.gz",
			prefix:   "_backup",
			want:     "archive.tar_backup.gz",
		},
		{
			name:     "hidden file without extension",
			filename: ".gitignore",
			prefix:   "_x",
			want:     "_x.gitignore", // first dot treated as extension separator
		},
		{
			name:     "empty filename",
			filename: "",
			prefix:   "_new",
			want:     "_new",
		},
		{
			name:     "empty prefix",
			filename: "file.txt",
			prefix:   "",
			want:     "file.txt",
		},
		{
			name:     "filename ending with dot",
			filename: "file.",
			prefix:   "_v2",
			want:     "file_v2.",
		},
		{
			name:     "prefix with dot",
			filename: "file.txt",
			prefix:   ".v2",
			want:     "file.v2.txt",
		},
		{
			name:     "unicode filename",
			filename: "файл.txt",
			prefix:   "_v2",
			want:     "файл_v2.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := internal.AddPrefixToFilename(tt.filename, tt.prefix)
			if got != tt.want {
				t.Errorf("AddPrefixToFilename(%q, %q) = %q, want %q",
					tt.filename, tt.prefix, got, tt.want)
			}
		})
	}
}
