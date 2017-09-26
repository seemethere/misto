package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	files                  = kingpin.Arg("files", "Files to parse through").Required().ExistingFiles()
	indentation, _         = regexp.Compile(`^[ \t]`)
	leadingTabWithSpace, _ = regexp.Compile(`^\t+ +`)
	leadingSpaceWithTab, _ = regexp.Compile(`^ +\t+`)
	space                  = " "
	tab                    = "\t"
)

// Output to stdout
func stdout(msg string) {
	_, err := fmt.Fprintln(os.Stdout, msg)
	check(err)
}

// Output to stderr
func stderr(msg string) {
	_, err := fmt.Fprintln(os.Stderr, msg)
	check(err)
}

func check(err error) {
	if err != nil {
		stderr(fmt.Sprintf("ERROR: %v", err))
	}
}

// DetectMixedIndent detects mixed indentation within a line
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
	formattedLine := strings.Replace(line, space, "•", -1)
	formattedLine = strings.Replace(formattedLine, tab, "›   ", -1)
	return formattedLine
}

// FileLine represents a line in a file but with extra metadata
type FileLine struct {
	LineContents string
	LineNumber   int
	ErrorCode    int
	IndentStyle  string
}

// DetectIndents detects indentation within a file
// returns fileLines in the form of a FileLine slice, also returns
// the majorityIndentStyle, either a TAB or SPACE
func DetectIndents(lines []string) (fileLines []FileLine, majorityIndentStyle string) {
	majorityIndentStyle = space
	tabCount := 0
	spaceCount := 0
	for lineNumber, line := range lines {
		// If the indent style doesn't match a SPACE or TAB
		// this will return an empty string
		indentStyle := indentation.FindString(line)
		switch indentStyle {
		case space:
			spaceCount++
		case tab:
			tabCount++
		}
		fileLines = append(
			fileLines,
			FileLine{
				LineContents: line,
				LineNumber:   lineNumber + 1,
				ErrorCode:    DetectMixedIndent(line),
				IndentStyle:  indentStyle,
			},
		)
	}
	if tabCount > spaceCount {
		majorityIndentStyle = tab
	}
	return fileLines, majorityIndentStyle
}

func processFile(filename string) {
	byteContents, err := ioutil.ReadFile(filename)
	check(err)
	fileContents := string(byteContents)
	fileLines, majorityIndentStyle := DetectIndents(strings.Split(fileContents, "\n"))
	for _, line := range fileLines {
		printLine := func(errorCode int) {
			stdout(fmt.Sprintf("%s:%d:MST%d:%s", filename, line.LineNumber, errorCode, formatLine(line.LineContents)))
		}
		if line.ErrorCode != 0 {
			printLine(line.ErrorCode)
		}
		if line.IndentStyle != "" && line.IndentStyle != majorityIndentStyle {
			printLine(3)
		}
	}
}

func main() {
	kingpin.Parse()
	for _, file := range *files {
		processFile(file)
	}
}
