package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"salvadorsru/po-generator/internal/action"
	"salvadorsru/po-generator/internal/bar"
	"salvadorsru/po-generator/internal/pot"
	"salvadorsru/po-generator/internal/utils"
	"sync"
)

func main() {
	inputPath := flag.String("i", ".", "Path to the folder to scan (default is current directory)")
	outputFile := flag.String("o", "to-translate.pot", "Name of the output POT file")
	pattern := flag.String("p", "php", "File extension  (default is php). Use '/' to separate multiple extensions (e.g., 'php/js/html').")

	flag.StringVar(inputPath, "input", ".", "Path to the folder to scan (default is current directory)")
	flag.StringVar(outputFile, "output", "to-translate.pot", "Name of the output POT file")

	help := flag.Bool("h", false, "Display help information")
	flag.Bool("help", false, "Display help information")

	flag.Parse()

	if *help {
		utils.PrintHelp()
		return
	}

	*inputPath = utils.ParsePath(*inputPath)

	if err := utils.CreateFolders(filepath.Dir(*outputFile)); err != nil {
		log.Fatalf("Could not create folders: %v", err)
	}

	potFile, potFileError := os.Create(*outputFile)
	if potFileError != nil {
		log.Panicf("error creating file %s", *outputFile)
	}
	defer potFile.Close()

	writeStringInPotFile := func(text string) error {
		_, writePotFileError := potFile.WriteString(text)
		if writePotFileError != nil {
			return fmt.Errorf("error writing to pot file: %v", writePotFileError)
		}
		return nil
	}

	writeStringInPotFile(pot.Headers())

	progressBar := bar.New()
	progressBar.TotalAsFilesIn(*inputPath)

	var wg sync.WaitGroup
	errorsChannel := make(chan error, 1)
	entriesChannel := make(chan pot.Entry)
	historyEntries := make(map[string]struct{})
	mu := &sync.Mutex{}

	inputPathError := filepath.Walk(*inputPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		wg.Add(1)
		go func(path string) {
			defer wg.Done()

			if utils.ValidatePathExtension(path, *pattern) {
				file, fileError := os.ReadFile(path)
				if fileError != nil {
					errorsChannel <- fmt.Errorf("error reading file %s: %v", path, fileError)
					return
				}

				content := string(file)

				for line, action := range action.AvailableActions {
					pattern := fmt.Sprintf("%s\\(.+\\)", action.Name)

					contentMatches := regexp.MustCompile(pattern)

					matches := contentMatches.FindAllString(content, -1)

					for _, match := range matches {
						sentence := string(match)

						entry := pot.ParseEntryFromString(
							pot.Reference{Base: path, Line: line},
							sentence,
							action.Parameters,
						)

						entriesChannel <- entry
					}
				}

			}
			progressBar.Print(1)
		}(path)

		return nil
	})

	go func() {
		for entry := range entriesChannel {
			entrySentence := entry.AsString()
			mu.Lock()
			if _, exists := historyEntries[entrySentence]; !exists {
				historyEntries[entrySentence] = struct{}{}
				if err := writeStringInPotFile(entry.ReferenceAsString()); err != nil {
					errorsChannel <- err
				}
				if err := writeStringInPotFile(entrySentence); err != nil {
					errorsChannel <- err
				}
			}
			mu.Unlock()
		}
	}()

	wg.Wait()
	close(entriesChannel)
	close(errorsChannel)

	for err := range errorsChannel {
		if err != nil {
			log.Println(err)
		}
	}

	if inputPathError != nil {
		log.Fatal(inputPathError)
	}

	trimError := utils.TrimFile(*outputFile)
	if trimError != nil {
		log.Fatal("Error trimming output file")
	}

	fmt.Println()
}
