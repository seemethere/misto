package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	files                  = kingpin.Arg("files", "Files to parse through").Required().ExistingFiles()
	leadingTabWithSpace, _ = regexp.Compile("^\t+ +")
	leadingSpaceWithTab, _ = regexp.Compile("^ +\t+")
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

// Determines error code for a given line
// Error codes are as follows:
//  * 1 = Leading tab(s) with following space(s)
//  * 2 = Leading space(s) with following tab(s)
func determineErrorCode(line string) int {
	errorCode := 0
	if match := leadingTabWithSpace.MatchString(line); match {
		errorCode = 1
	}
	if match := leadingSpaceWithTab.MatchString(line); match {
		errorCode = 2
	}
	return errorCode
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
		linePrinted := false
		printLine := func(errorCode int) {
			if !linePrinted {
				stdout(fmt.Sprintf("%s:%d:MST%d %s", filename, lineNo, errorCode, line))
				linePrinted = true
			}
		}
		errorCode := determineErrorCode(line)
		if errorCode != 0 {
			printLine(errorCode)
		}
		lineNo++
	}
	check(scanner.Err())
}

func main() {
	kingpin.Parse()
	for _, file := range *files {
		parseFile(file)
	}
}
