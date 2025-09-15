package fileio

import "os"

// TempDir creates a new temporary directory
func TempDir() (string, error) {
	return os.MkdirTemp("", "gosubai_tmp")
}
