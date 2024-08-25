package bar

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type bar struct {
	total   int
	current int
}

func New() bar {
	return bar{
		total:   0,
		current: 0,
	}
}

func (bar *bar) Add(number int) {
	bar.current += number
}

func (bar *bar) Print(added int) {
	if added != 0 {
		bar.Add(1)
	}

	percentage := float64(bar.current) / float64(bar.total) * 100
	barLength := 50
	progress := int(float64(barLength) * percentage / 100)

	barToPrint := fmt.Sprintf("[%s%s] %.2f%%",
		strings.Repeat("=", progress),
		strings.Repeat(" ", barLength-progress),
		percentage)

	fmt.Printf("\r%s %d/%d", barToPrint, bar.current, bar.total)
}

func (bar *bar) TotalAsFilesIn(path string) *bar {
	fileCount := 0
	fileCountError := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fileCount++
		}
		return nil
	})

	if fileCountError != nil {
		log.Fatal(fileCountError)
	}

	bar.total = fileCount
	return bar
}
