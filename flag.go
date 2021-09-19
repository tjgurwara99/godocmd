package main

import "flag"

var recursive bool

func init() {
	flag.BoolVar(&recursive, "r", false, "Recursively traverse the source")
}
