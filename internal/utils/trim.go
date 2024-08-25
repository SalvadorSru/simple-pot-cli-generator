package utils

import (
	"bytes"
	"os"
	"strings"
)

func TrimFile(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	lines := bytes.Split(content, []byte("\n"))

	startIndex := 0
	for startIndex < len(lines) && strings.TrimSpace(string(lines[startIndex])) == "" {
		startIndex++
	}

	endIndex := len(lines) - 1
	for endIndex >= 0 && strings.TrimSpace(string(lines[endIndex])) == "" {
		endIndex--
	}

	if startIndex > endIndex {
		lines = [][]byte{}
	} else {
		lines = lines[startIndex : endIndex+1]
	}

	trimmedContent := bytes.Join(lines, []byte("\n"))

	err = os.WriteFile(filePath, trimmedContent, 0644)
	if err != nil {
		return err
	}

	return nil
}
