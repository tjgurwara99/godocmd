package main

import "flag"

var (
	recursive   bool
	write       bool
	module      string
	sourcePath  string
	fileToWrite string
)

func init() {
	flag.BoolVar(&recursive, "r", false, "Recursively traverse the source")
	flag.StringVar(&module, "module", "", "Module path prefix for the source")
	flag.BoolVar(&write, "w", false, "Write the output to the file given by -file flag.")
	flag.StringVar(&fileToWrite, "file", "./README.md", "File to write the output to.")
}
