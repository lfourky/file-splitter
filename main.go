package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"flag"
	"path/filepath"
)

func main() {

	//Flag setup
	var (
		bufferSize = flag.Uint("b", 10000000, "Customize buffer size.")
		limit      = flag.Uint("l", 10, "Line limit per part-file.")

		partsDir       = flag.String("d", "parts", "Provide a directory path where partial files will be stored")
		partsFileName  = flag.String("p", "part_", "Provide a prefix for each part.")
		partFileSuffix = flag.String("s", ".sql", "Provide a suffix for each part.")

		fileToSplit = flag.String("f", "", "Provide a relative path to file you wish to split.")
	)

	flag.Parse()

	//Try to read the file
	file, err := os.Open(*fileToSplit)
	if err != nil {
		log.Fatal("Error trying to open the file specified. Err: ", err)
	}
	defer file.Close()

	//Get a scanner to read the file with
	scanner := bufio.NewScanner(file)

	var c uint = 0

	fCounter := 1

	var pFile *os.File

	scanner.Buffer(make([]byte, *bufferSize), int(*bufferSize))

	//Start scanning
	for scanner.Scan() {
		if c == *limit {
			//Close the file, reset counter
			if err := pFile.Close(); err != nil {
				log.Fatal("Error closing file. Err: ", err)
			}
			c = 0
		}

		if c == 0 {
			//Open the file, start writing to it
			fName := fmt.Sprintf("%s%d%s", *partsFileName, fCounter, *partFileSuffix)
			fName = filepath.Join(*partsDir, fName)
			if pFile, err = os.Create(fName); err != nil {
				log.Fatal("Error creating a file. Err:", err)
			}
			fCounter++
		}

		if c <= *limit {
			//Keep writing to the same file, increment counter
			if _, err := fmt.Fprintln(pFile, scanner.Text()); err != nil {
				log.Fatal("Error writing to file. Err: ", err)
			}
			c++
		}
	}
}
