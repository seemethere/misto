package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	files = kingpin.Arg("files", "Files to parse through").Required().ExistingFiles()
)

// Output to stdout
func stdout(msg string) {
	fmt.Fprintln(os.Stdout, msg)
}

// Output to stderr
func stderr(msg string) {
	fmt.Fprintln(os.Stderr, msg)
}

func main() {
	kingpin.Parse()
	fmt.Printf("%v", *files)
}
