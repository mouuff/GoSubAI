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
