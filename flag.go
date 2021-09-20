package main

import "flag"

var recursive bool

var module string

func init() {
	flag.BoolVar(&recursive, "r", false, "Recursively traverse the source")
	flag.StringVar(&module, "module", "", "Module path prefix for the source")
}
