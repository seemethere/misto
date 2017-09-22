package main

import (
	"bufio"
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	files = kingpin.Arg("files", "Files to parse through").Required().ExistingFiles()
)

// Output to stdout
func stdout(msg string) {
	fmt.Fprintln(os.Stdout, msg) //nolint
}

// Output to stderr
func stderr(msg string) {
	fmt.Fprintln(os.Stderr, msg) //nolint
}

func check(err error) {
	if err != nil {
		stderr(fmt.Sprintf("ERROR: %v", err))
	}
}

func parseFile(filename string) {
	inFile, err := os.Open(filename)
	check(err)
	defer func() {
		check(inFile.Close())
	}()
	scanner := bufio.NewScanner(inFile)

	lineNo := 0
	for scanner.Scan() {
		line := scanner.Text()
		stdout(fmt.Sprintf("%s:%d %s", filename, lineNo, line))
		lineNo++
	}
	check(scanner.Err())
}

func main() {
	kingpin.Parse()

	for _, filename := range *files {
		parseFile(filename)
	}
}
