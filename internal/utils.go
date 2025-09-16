package internal

import (
	"encoding/json"
	"os"
)

func ReadFromJson(path string, dataOut interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(data), dataOut); err != nil {
		return err
	}

	return nil
}

func AddPrefixToFilename(filename, prefix string) string {
	extIdx := len(filename)
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			extIdx = i
			break
		}
	}
	return filename[:extIdx] + prefix + filename[extIdx:]
}
