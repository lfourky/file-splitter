package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	var (
		bufferSize = flag.Uint("b", 10000000, "Customize buffer size.")
		lineLimit  = flag.Uint("l", 10, "Line limit per part-file.")

		partsDir       = flag.String("d", "parts", "Provide a directory path where partial files will be stored")
		partsFileName  = flag.String("p", "part_", "Provide a prefix for each part.")
		partFileSuffix = flag.String("s", "", "Provide a suffix for each part.")

		fileToSplit = flag.String("f", "", "Provide a relative path to file you wish to split.")
	)

	flag.Parse()

	// Try to read the file
	file, err := os.Open(*fileToSplit)
	if err != nil {
		log.Fatal("error trying to open the file specified:", err)
	}
	defer file.Close()

	// Check if the provided directory path for partials exits; create if it doesn't
	if _, err := os.Stat(*partsDir); os.IsNotExist(err) {
		if err = os.MkdirAll(*partsDir, os.ModePerm); err != nil {
			log.Fatal("error creating directory:", err)
		}
	}

	// Decide on the final partial file suffix. Choose flag suffix over the provided file's suffix
	var partialFileSuffix string
	if *partFileSuffix != "" {
		partialFileSuffix = *partFileSuffix
	} else {
		suffixStart := strings.LastIndex(*fileToSplit, ".")
		if suffixStart == -1 {
			log.Fatal("Please provide a suffix for the partial files using -s. For example -s .sql")
		}
		partialFileSuffix = (*fileToSplit)[suffixStart:]
	}

	// Get a scanner to read the file with
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, *bufferSize), int(*bufferSize))

	// Counts created files
	fileCounter := 1

	var linesWritten uint
	var currentFile *os.File
	var fileName string

	// Start scanning
	for scanner.Scan() {
		if linesWritten == *lineLimit {
			// Close the file, reset counter
			if err := currentFile.Close(); err != nil {
				log.Fatal("error closing file:", err)
			}
			linesWritten = 0
		}

		if linesWritten == 0 {
			// Open the file, start writing to it
			fileName = fmt.Sprintf("%s%d%s", *partsFileName, fileCounter, partialFileSuffix)
			fileName = filepath.Join(*partsDir, fileName)
			if currentFile, err = os.Create(fileName); err != nil {
				log.Fatal("error creating a file:", err)
			}
			fileCounter++
		}

		if linesWritten < *lineLimit {
			// Keep writing to the same file, increment counter
			if _, err := fmt.Fprintln(currentFile, scanner.Text()); err != nil {
				log.Fatal("error writing to file:", err)
			}
			linesWritten++
		}
	}
}
