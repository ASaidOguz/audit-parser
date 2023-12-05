package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func collectAuditTags(folderPath string, tag string, outputFilename string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			fileContent := string(content)
			lines := strings.Split(fileContent, "\n")
			for _, line := range lines {
				if strings.Contains(line, tag) {
					ch <- line
				}
			}
		}
		return nil
	})

	if err != nil {
		log.Println(err)
	}
	close(ch) // Close the channel after completing tasks
}

func writeToOutput(ch chan string, outputFilename string, wg *sync.WaitGroup) {
	defer wg.Done()

	outputFile, err := os.Create(outputFilename)
	if err != nil {
		log.Println(err)
		return
	}
	defer outputFile.Close()

	var lineWG sync.WaitGroup

	for line := range ch {
		lineWG.Add(1)
		go func(line string) {
			defer lineWG.Done()
			// Process the line (remove indicators, etc.)
			processedLine := strings.ReplaceAll(line, "//", "")
			processedLine = strings.ReplaceAll(processedLine, "/*", "")
			processedLine = strings.ReplaceAll(processedLine, "*/", "")

			_, err := outputFile.WriteString(processedLine + "\n")
			if err != nil {
				log.Println(err)
			}
		}(line)
	}

	lineWG.Wait() // Wait for all lines to be processed
}

func main() {
	args := os.Args[1:] // Get all command-line arguments except the program name

	// Check if there are enough arguments provided
	if len(args) < 3 {
		log.Fatal("Usage: ./auditparser <folder-path> <output-filename> <tag>")
	}

	folder := args[0]
	outputFilename := args[1]
	tag := args[2]

	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(2) // Two goroutines to wait for

	go collectAuditTags(folder, tag, outputFilename, ch, &wg)
	go writeToOutput(ch, outputFilename, &wg)

	wg.Wait() // Wait for all tasks to complete

	// Print a message indicating completion
	log.Println("Parsing and writing completed.")
}
