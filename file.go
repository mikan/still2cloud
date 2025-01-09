package main

import (
	"fmt"
	"os"
)

func readFile(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("[SRC] failed to read %s: %v", path, err)
	}
	return content, nil
}

func writeFile(path string, content []byte) error {
	if err := os.WriteFile(path, content, 0644); err != nil {
		return fmt.Errorf("[DEST] failed to write %s: %v", path, err)
	}
	return nil
}
