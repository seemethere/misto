package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	files                  = kingpin.Arg("files", "Files to parse through").Required().ExistingFiles()
	indentation, _         = regexp.Compile("^\\s")
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
func DetectMixedIndent(line string) int {
	errorCode := 0
	if match := leadingTabWithSpace.MatchString(line); match {
		errorCode = 1
	}
	if match := leadingSpaceWithTab.MatchString(line); match {
		errorCode = 2
	}
	return errorCode
}

func formatLine(line string) string {
	formattedLine := strings.Replace(line, " ", "•", -1)
	formattedLine = strings.Replace(formattedLine, "\t", "›   ", -1)
	return formattedLine
}

type FileLine struct {
	lineNo int
	line   string
}

func parseFile(filename string) {
	inFile, err := os.Open(filename)
	check(err)
	defer func() {
		check(inFile.Close())
	}()
	scanner := bufio.NewScanner(inFile)

	lineNo := 0
	indents := make(map[string][]FileLine)
	for scanner.Scan() {
		line := scanner.Text()
		// Skip lines that do not leading whitespace
		indent := indentation.FindString(line)
		if indent == "" {
			continue
		}
		linePrinted := false
		printLine := func(errorCode int) {
			if !linePrinted {
				formattedLine := formatLine(line)
				// Make tabs / spaces easier to see
				stdout(fmt.Sprintf("%s:%d:MST%d %s", filename, lineNo, errorCode, formattedLine))
				linePrinted = true
			}
		}
		errorCode := DetectMixedIndent(line)
		if errorCode == 0 {
			indents[indent] = append(indents[indent], FileLine{lineNo, line})
		} else {
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
