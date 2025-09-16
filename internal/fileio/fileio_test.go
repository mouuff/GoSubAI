package fileio_test

import (
	"os"
	"testing"

	"github.com/mouuff/GoSubAI/internal/fileio"
)

func TestTempDir(t *testing.T) {
	tmpDir, err := fileio.TempDir()
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	defer os.RemoveAll(tmpDir)

	tmpDirInfo, err := os.Stat(tmpDir)
	if err != nil {
		t.Fatalf("Failed to stat temp dir: %v", err)
	}
	if !tmpDirInfo.IsDir() {
		t.Fatalf("Temp path is not a directory")
	}

	t.Logf("Temp dir created successfully: %s", tmpDir)
}
